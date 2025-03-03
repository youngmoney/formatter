package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Match(filename string, matchers *[]Matcher) *Matcher {
	absname, err := filepath.Abs(filename)
	if err != nil {
		fmt.Println("error finding file")
		os.Exit(1)
	}
	shebang := GetShebang(filename)
	for _, m := range *matchers {
		if m.PathRegex != nil && !m.PathRegex.MatchString(absname) {
			continue
		}
		if m.ShebangRegex != nil && !m.ShebangRegex.MatchString(shebang) {
			continue
		}
		return &m
	}
	return nil
}

func GetShebang(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, e := r.ReadString('\n')
	if e == nil && strings.HasPrefix(s, "#!") {
		return s
	}
	return ""
}
