package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func GetCommand(name string, commands *[]Command) *Command {
	for _, c := range *commands {
		if c.Name == name {
			return &c
		}
	}
	return nil
}

func ExecuteCommand(filename string, command *Command) {
	cmd := exec.Command("bash", "-s", filename)
	cmd.Env = append(os.Environ(), "FILENAME="+filename)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, command.Command)
	}()

	out, err := cmd.CombinedOutput()
	fmt.Printf("%s", out)
	if err != nil {
		if e, ok := err.(interface{ ExitCode() int }); ok {
			os.Exit(e.ExitCode())
		}
	}
}
