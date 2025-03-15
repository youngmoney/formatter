package main

import (
	"flag"
	"fmt"
	"os"
)

func commandLint(config Config, filename string) {
	if len(filename) == 0 {
		fmt.Fprintln(os.Stderr, "provide a file to lint")
		os.Exit(1)
	}
	match := Match(filename, &config.Formatter.Matchers)
	if match == nil || match.LinterName == "" {
		fmt.Fprintln(os.Stderr, "Ignoring:", filename)
		return
	}

	command := GetCommand(match.LinterName, &config.Formatter.Linters)

	if command == nil {
		fmt.Fprintln(os.Stderr, "Ignoring:", filename, " no command")
		return
	}

	ExitIfNonZero(ExecuteCommandInteractive(filename, command.Command))
}

func commandFix(config Config, filename string) {
	if len(filename) == 0 {
		fmt.Fprintln(os.Stderr, "provide a file to fix")
		os.Exit(1)
	}

	match := Match(filename, &config.Formatter.Matchers)
	if match == nil || match.FixerName == "" {
		fmt.Fprintln(os.Stderr, "Ignoring:", filename)
		return
	}

	command := GetCommand(match.FixerName, &config.Formatter.Fixers)

	if command == nil {
		fmt.Fprintln(os.Stderr, "Ignoring:", filename, " no command")
		return
	}

	ExitIfNonZero(ExecuteCommandInteractive(filename, command.Command))
}

func main() {
	configFilename := flag.String("config", os.Getenv("FORMATTER_CONFIG"), "config file (yaml)")
	flag.Parse()
	config := ReadConfig(*configFilename)
	switch flag.Arg(0) {
	case "lint":
		commandLint(config, flag.Arg(1))
	case "fix":
		commandFix(config, flag.Arg(1))
	default:
		fmt.Fprintln(os.Stderr, "formatter lint <filename>")
		fmt.Fprintln(os.Stderr, "formatter fix <filename>")
		os.Exit(1)
	}
}
