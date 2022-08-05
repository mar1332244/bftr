package main

import (
    "flag"
    "fmt"
    "io"
    "math"
    "os"
    "strings"
)

const (
    add10  = ">%s[-<++++++++++>]<"
    sub10  = ">%s[-<---------->]<"
    errMsg = "couldn't %s `%s` (%s)"
)

type Flags struct {
    width  uint
    output string
    from   string
    doWrap bool
}

func parseFlags() Flags {
    var f Flags
    flag.UintVar(
        &f.width, "w", 0, "split output into lines of specified width",
    )
    flag.StringVar(
        &f.output, "o", "", "write the output to a specified file",
    )
    flag.StringVar(
        &f.from, "f", "", "get the input from a specified file",
    )
    flag.BoolVar(
        &f.doWrap, "do-wrap", false, "consider cell wrapping while creating output",
    )
    flag.Parse()
    return f
}

func formatError(err error, verb string) error {
    if err, ok := err.(*os.PathError); ok {
        return fmt.Errorf(errMsg, verb, err.Path, err.Err)
    }
    return fmt.Errorf("unknown error encountered")
}

func getOutputFile(f Flags) (*os.File, error) {
    if f.output == "" {
        return os.Stdout, nil
    }
    outFile, err := os.Create(f.output)
    if err != nil {
        return nil, formatError(err, "create")
    }
    return outFile, nil
}

func printOutput(output string, f Flags) error {
    outFile, err := getOutputFile(f)
    if err != nil {
        return err
    }
    if f.width == 0 {
        f.width = uint(len(output))
    }
    for 0 < len(output) {
        a := float64(len(output))
        b := float64(f.width)
        width := int(math.Min(a, b))
        fmt.Fprintln(outFile, output[:width])
        output = output[width:]
    }
    if outFile != os.Stdout {
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

func readFile(fp io.Reader) string {
    var builder strings.Builder
    buffer := make([]byte, 1024)
    for {
        n, err := fp.Read(buffer)
        if n == 0 && err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
        }
        builder.Write(buffer[:n])
    }
    return builder.String()
}

func getInput(f Flags) (string, error) {
    stat, err := os.Stdin.Stat()
    if err != nil {
        panic(err)
    }
    if (stat.Mode() & os.ModeCharDevice) == 0 {
        return readFile(os.Stdin), nil
    }
    if f.from != "" {
        inFile, err := os.Open(f.from)
        if err != nil {
            return "", formatError(err, "open")
        }
        defer inFile.Close()
        input := readFile(inFile)
        return input, nil
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
