package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	add10 = ">%s[-<++++++++++>]<"
	sub10 = ">%s[-<---------->]<"
)

func toBrainfuck(input string) string {
	var builder strings.Builder
	var last byte
	for _, char := range []byte(input) {
		delta := int(char) - int(last)
		loops := delta / 10
		extra := delta % 10
		symbol := "+"
		loop := add10
		if delta < 0 {
			loops = -delta / 10
			extra = -delta % 10
			symbol = "-"
			loop = sub10
		}
		if 0 < loops {
			fmt.Fprintf(&builder, loop, strings.Repeat("+", loops))
		}
		builder.WriteString(strings.Repeat(symbol, extra))
		builder.WriteByte('.')
		last = char
	}
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
