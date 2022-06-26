package main

import (
    "strconv"
)

var flagNames map[string]interface{}

func StrFlag(name string, defultValue string) *string {
    ptr := new(string)
    *ptr = defaultValue
    flagNames[name] = ptr
    return ptr
}

func BoolFlag(name string, defaultValue bool) *bool {
    ptr := new(bool)
    *ptr = defaultValue
    flagNames[name] = ptr
    return ptr
}

func IntFlag(name string, defaultValue int) *int {
    ptr := new(int)
    *ptr = defaultValue
    flagNames[name] = ptr
    return ptr
}

/*
func StrArg(index int, defVal string) *string

func BoolArg(index int, defVal bool) *bool

func IntArg(index int, defVal int) *int
*/

func Parse(argv []string) {
    for i := 0; i < len(argv); i++ {
        ptr, ok := flagNames[argv[i]]
        if !ok {
            continue
        }
        if i == len(argv) - 1 {
            break
        }
        i++
        switch ptr.(type) {
        case *string:
            *ptr = argv[i]
        case *bool:
            b, err := strconv.ParseBool(argv[i])
            if err != nil {
                continue
            }
            *ptr = b
        case *int:
            x, err := strconv.Atoi(argv[i])
            if err != nil {
                continue
            }
            *ptr = x
        }
    }
}
