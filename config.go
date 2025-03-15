package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"regexp"
)

type Command struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
}

type Matcher struct {
	PathRegex    *regexp.Regexp `yaml:"path_regex"`
	ShebangRegex *regexp.Regexp `yaml:"shebang_regex"`
	LinterName   string         `yaml:"linter_name"`
	FixerName    string         `yaml:"fixer_name"`
}

type Formatter struct {
	Linters  []Command `yaml:"linters"`
	Fixers   []Command `yaml:"fixers"`
	Matchers []Matcher `yaml:"matchers"`
}

type Config struct {
	Formatter Formatter `yaml:"formatter"`
}

func ReadConfig(filename string) Config {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	config := Config{}
	if err := yaml.Unmarshal(raw, &config); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return config
}
