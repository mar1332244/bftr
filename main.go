package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	add10 = ">%s[-<++++++++++>]<"
	sub10 = ">%s[-<---------->]<"
)

type Flags struct {
	width  int
	output string
	from   string
}

func parseFlags() Flags {
	var f Flags
	flag.IntVar(&f.width, "w", -1, "split output into lines of specified `width`")
	flag.StringVar(&f.output, "o", "stdout", "write the `output` to a specified file")
	flag.StringVar(&f.from, "f", "", "get the input `from` a specified file")
	flag.Parse()
	return f
}

func getOutputFile(f Flags) (*os.File, error) {
	if f.output == "stdout" {
		return os.Stdout, nil
	}
	outFile, err := os.Create(f.output)
	if err, ok := err.(*os.PathError); ok {
		return nil, fmt.Errorf("couldn't create `%s` (%s)", err.Path, err.Err)
	}
	return outFile, nil
}

func printOutput(output string, f Flags) error {
	outFile, err := getOutputFile(f)
	if err != nil {
		return err
	}
	if f.width <= 0 {
		fmt.Fprintln(outFile, output)
		return nil
	}
	buffer := bytes.NewBufferString(output)
	for 0 < buffer.Len() {
		line := buffer.Next(f.width)
		fmt.Fprintf(outFile, "%s\n", line)
	}
	if outFile != os.Stdin {
		outFile.Close()
	}
	return nil
}

// TODO: implement a better conversion algorithm
func toBrainfuck(input string, f Flags) string {
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

func readFile(fname string) (string, error) {
	inFile, err := os.Open(fname)
	if err, ok := err.(*os.PathError); ok {
		return "", fmt.Errorf("couldn't open `%s` (%s)", err.Path, err.Err)
	}
	defer inFile.Close()
	var builder strings.Builder
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteByte('\n')
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
	return builder.String(), nil
}

func readPipedInput() (string, error) {
	var builder strings.Builder
	buffer := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buffer)
		if n == 0 && err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		builder.Write(buffer[:n])
	}
	return builder.String(), nil
}

func getInput(f Flags) (string, error) {
	if f.from != "" {
		return readFile(f.from)
	}
	stat, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return readPipedInput()
	}
	if flag.NArg() == 0 {
		return "", fmt.Errorf("no input provided")
	}
	return flag.Arg(0), nil
}

func main() {
	f := parseFlags()
	input, err := getInput(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		flag.Usage()
		os.Exit(1)
	}
	output := toBrainfuck(input, f)
	err = printOutput(output, f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
