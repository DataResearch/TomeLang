package main

import (
	"os"
	"programminglanguage/build"
	"programminglanguage/repl"
	"strings"
)

/*
	Command Line Options:

	-token | Repl Prints Tokens
	-ast | Repl Prints Ast

	If a *.tm file is provided no repl will start and we attempt to compile
	all *.tm files.
*/

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
	commandlineargs := os.Args[1:]

	mode := build.EXECUTE
	files := []string{}

	for _, param := range commandlineargs {
		if param == "--token" {
			mode = build.TOKENS
		}
		if param == "--ast" {
			mode = build.AST
		}
		if strings.HasSuffix(param, ".tm") {
			files = append(files, param)
		}
	}

	if len(files) == 0 {
		repl.Start(os.Stdin, os.Stdout, mode)
	} else {
		// TODO: here ya go - put your LLVM integration here
	}
}
