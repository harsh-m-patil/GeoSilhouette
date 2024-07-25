package main

import (
	"os"
	"strings"
)

func main() {
	repl()
}

func exit() {
	os.Exit(0)
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
