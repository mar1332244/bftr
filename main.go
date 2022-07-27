package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	add10 = "[->++++++++++<]>"
	sub10 = "[->----------<]>"
)

func toBrainfuck(input string) string {
	var builder strings.Builder
	return builder.String()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "no input provided")
		os.Exit(1)
	}
	input := os.Args[1]
	output := toBrainfuck(input)
	fmt.Println(output)
}
