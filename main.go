package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	add10 = ">%s[-<++++++++++>]<"
	sub10 = ">%s[-<---------->]<"
)

type Flags struct {
	maxWidth int
	outFile  string
}

func parseFlags() Flags {
	var f Flags
	flag.IntVar(&f.maxWidth, "w", -1, "sets the max width of a given line of output")
	flag.StringVar(&f.outFile, "o", "stdout", "filepath to write to the output to")
	flag.Parse()
	return f
}

func printOutput(output string, f Flags) {
	if f.maxWidth <= 0 {
		fmt.Println(output)
		return
	}
	buffer := bytes.NewBufferString(output)
	for 0 < buffer.Len() {
		line := buffer.Next(f.maxWidth)
		fmt.Printf("%s\n", line)
	}
}

func toBrainfuck(input string) string {
	var builder strings.Builder
	var last byte
	for i := 0; i < len(input); i++ {
		char := input[i]
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
			loop = fmt.Sprintf(loop, strings.Repeat("+", loops))
			builder.WriteString(loop)
		}
		builder.WriteString(strings.Repeat(symbol, extra))
		builder.WriteByte('.')
		last = char
	}
	return builder.String()
}

func main() {
	f := parseFlags()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "no input provided")
		os.Exit(1)
	}
	input := flag.Arg(0)
	output := toBrainfuck(input)
	printOutput(output, f)
}
