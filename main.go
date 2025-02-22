package main

import (
	"flag"
	"fmt"
	"os"
)

func commandLint(config Config, filename string) {
	if len(filename) == 0 {
		fmt.Println("provide a file to lint")
		os.Exit(1)
	}
	match := Match(filename, &config.Formatter.Matchers)
	if match == nil || match.LinterName == "" {
		fmt.Println("Ignoring: %s", filename)
	}

	command := GetCommand(match.LinterName, &config.Formatter.Linters)

	if command == nil {
		fmt.Println("Ignoring: %s", filename)
	}

	ExecuteCommand(filename, command)
}

func commandFix(config Config, filename string) {
	if len(filename) == 0 {
		fmt.Println("provide a file to fix")
		os.Exit(1)
	}

	match := Match(filename, &config.Formatter.Matchers)
	if match == nil || match.FixerName == "" {
		fmt.Println("Ignoring: %s", filename)
	}

	command := GetCommand(match.FixerName, &config.Formatter.Fixers)

	if command == nil {
		fmt.Println("Ignoring: %s", filename)
	}

	ExecuteCommand(filename, command)
}

func main() {
	configFilename := flag.String("config", os.Getenv("FORMATTER_CONFIG"), "config file (yaml)")
	flag.Parse()
	config := ReadConfig(*configFilename)
	if flag.NArg() != 2 {
		fmt.Println("command and filename required")
		os.Exit(1)
	}
	switch flag.Arg(0) {
	case "lint":
		commandLint(config, flag.Arg(1))
	case "fix":
		commandFix(config, flag.Arg(1))
	default:
		fmt.Printf("error: unknown command - %q\n", flag.Arg(0))
		// In a real program (not an example) print to os.Stderr and exit the program with non-zero value.
	}
}
