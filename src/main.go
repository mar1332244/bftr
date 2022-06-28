package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        return
    }
    if err := TextToBrainfuck(os.Args[1], os.Args[2]); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}
