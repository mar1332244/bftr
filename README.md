# bftr

Small utility written in Go to translate ascii text into valid brainfuck code

## Installation

To install bftr run the following command

```sh
go install github.com/mar1332244/bftr@latest
```

This will place the binary executable into your $GOPATH/bin directory. From there you can move the file anywhere
like ```/usr/local/bin``` for example. To uninstall the executable simply delete the file.

## Usage

### Basics

Running the command

```sh
bftr "Hello, World!"
```

will output the following

```
>+++++++[-<++++++++++>]<++.>++[-<++++++++++>]<+++++++++.+++++++..+++.>++++++[-<---------->]<-------.>+[-<---------->]<--.>+++++[-<++++++++++>]<+++++.>++[-<++++++++++>]<++++.+++.------.--------.>++++++[-<---------->]<-------.
```

Make sure when you provide text with spaces it is wrapped in double quotes otherwise only the first word will get translated.

### Flags

| Flag          | Description                                                      |
| ------------- |:----------------------------------------------------------------:|
| -f [filename] | reads the file provided and uses its constents as the input      |
| -o [filename] | creates the file provided and prints the brainfuck code to it    |
| -w [integer]  | splits the output into lines of no wider than the number provided|
