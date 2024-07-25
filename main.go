package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	repl()
}

func exit() {
	defer os.Exit(0)
	fmt.Println("Exitting")
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
