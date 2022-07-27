package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

const (
	add10 string = "[<++++++++++>-]"
	sub10 string = "[<---------->-]"
)

func TextToBrainfuck(src, dst string) error {
	inFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer inFile.Close()
	var builder strings.Builder
	var line string
	reader := bufio.NewReader(inFile)
	for {
		line, err = reader.ReadString('\n')
		builder.WriteString(line)
		if err != nil {
			break
		}
	}
	outFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer outFile.Close()
	var last byte = 0
	var change, loops, extra int
	text := builder.String()
	for i := 0; i < len(text); i++ {
		builder.Reset()
		builder.WriteString(fmt.Sprintf("(%d) ", text[i]))
		for builder.Len() < 6 {
			builder.WriteByte(' ')
		}
		builder.WriteString("| ")
		if text[i] == last {
			builder.WriteString(".\n")
			if _, err = outFile.WriteString(builder.String()); err != nil {
				return err
			}
			continue
		}
		change = int(text[i]) - int(last)
		if change > 0 {
			extra = change % 10
			loops = (change - extra) / 10
		} else {
			extra = -change % 10
			loops = (change + extra) / -10
		}
		if loops > 0 {
			builder.WriteByte('>')
			builder.WriteString(strings.Repeat("+", loops))
			if change > 0 {
				builder.WriteString(add10)
			} else {
				builder.WriteString(sub10)
			}
			builder.WriteByte('<')
		}
		if extra > 0 && change > 0 {
			builder.WriteString(strings.Repeat("+", extra))
		} else if extra > 0 {
			builder.WriteString(strings.Repeat("-", extra))
		}
		builder.WriteString(".\n")
		if _, err = outFile.WriteString(builder.String()); err != nil {
			return err
		}
		last = text[i]
	}
	return nil
}

func main() {
    if len(os.Args) < 3 {
        fmt.Fprintln(os.Stdout, "not enough args")
		os.Exit(-1)
    }
    if err := TextToBrainfuck(os.Args[1], os.Args[2]); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}
