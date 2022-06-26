# bftr

Small utility written in Go to translate ascii text into many Brainfuck print commands

## Installation

Using git and make

```
$ git clone https://github.com/mar1332244/bftr.git
$ cd bftr
$ make
```

This will add bftr to your path. If you wish to uninstall bftr run ```$ make clean``` while inside of the cloned directory to remove the binary from your path.

## Usage

```
$ bftr [flags] [src] [dst]
$ cat hello.txt
Hello, World!
$ bftr hello.txt hello.bf
$ cat hello.bf
(72)  | >+++++++[<++++++++++>-]<++.
(101) | >++[<++++++++++>-]<+++++++++.
(108) | +++++++.
(108) | .
(111) | +++.
(44)  | >++++++[<---------->-]<-------.
(32)  | >+[<---------->-]<--.
(87)  | >+++++[<++++++++++>-]<+++++.
(111) | >++[<++++++++++>-]<++++.
(114) | +++.
(108) | ------.
(100) | --------.
(33)  | >++++++[<---------->-]<-------.
(10)  | >++[<---------->-]<---.
```

## License

bftr is under the MIT license
