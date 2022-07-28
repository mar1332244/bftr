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

	createErr = "couldn't create `%s` (%s)"

	wUse = "sets the max `width` of a given line of output"
	oUse = "creates a file and prints the `output` to it"
)

type Flags struct {
	maxWidth int
	outFname string
}

func parseFlags() Flags {
	var f Flags
	flag.IntVar(&f.maxWidth, "w", -1, wUse)
	flag.StringVar(&f.outFname, "o", "stdout", oUse)
	flag.Parse()
	return f
}

func printOutput(output string, f Flags) error {
	outFile := os.Stdout
	if f.outFname != "stdout" {
		var err error
		outFile, err = os.Create(f.outFname)
		if err, ok := err.(*os.PathError); ok {
			return fmt.Errorf(createErr, err.Path, err.Err)
		}
	}
	if f.maxWidth <= 0 {
		fmt.Fprintln(outFile, output)
		return nil
	}
	buffer := bytes.NewBufferString(output)
	for 0 < buffer.Len() {
		line := buffer.Next(f.maxWidth)
		fmt.Fprintf(outFile, "%s\n", line)
	}
	return nil
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
		flag.Usage()
		os.Exit(1)
	}
	input := flag.Arg(0)
	output := toBrainfuck(input)
	err := printOutput(output, f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
