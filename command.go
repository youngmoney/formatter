package main

import (
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

func ExitIfNonZero(err interface{}) {
	if err != nil {
		if e, ok := err.(interface{ ExitCode() int }); ok {
			os.Exit(e.ExitCode())
		}
	}
}

func ExecuteCommandInteractive(filename string, command string) error {
	bashArgs := []string{"-c", command, "command"}
	cmd := exec.Command("bash", append(bashArgs, filename)...)

	cmd.Env = append(os.Environ(), "FILENAME="+filename)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
