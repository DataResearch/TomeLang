package main

import (
	"os"
	"programminglanguage/repl"
)

// Command-Line Options:

// -REPL=Lexer Prints the Tokens
// -REPL=Parser Prints the AST
// -REPL=Exec | -REPL Starts an Interactive Terminal that executes commands

// -F=<file> execute file

/*
>> let x = fn(n) {return if (n < 1) {1} else {n * x(n - 1)};};
	no prefix parse function for { found
	no prefix parse function for else found
	no prefix parse function for { found
	no prefix parse function for } found

TODO: investigate / errorhandling

>> let x = fn(n) {return if (n < 1) {1} else {n * x(n - 1)};};
let x = fn(n)return if (n < 1) 1else (n * x((n - 1)));;
 TODO: formatting is horrible

>> 0123456789
	Could not parse %q as integer0123456789
TODO: starting numbers with a "0" will make to treat them as (probably) an octal value
which is undesired.
*/
func main() {
	repl.Start(os.Stdin, os.Stdout)
}
