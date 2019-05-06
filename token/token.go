package token

import "bytes"

type TokenType = string

type Token struct {
	Type    TokenType
	Literal string
}

func (tok *Token) String() string {
	var out bytes.Buffer

	out.WriteString("{")
	out.WriteString("Type: ")
	out.WriteString(tok.Type)
	out.WriteString(", ")
	out.WriteString("Literal: ")
	out.WriteString(tok.Literal)
	out.WriteString("}")

	return out.String()
}

const (
	// Support
	ILLEGAL TokenType = "illegal"
	EOF     = "eof"

	// Operators
	ASSIGN         = "="
	PLUS           = "+"
	MINUS          = "-"
	MULTIPLICATION = "*"
	DIVISION       = "/"
	LSHIFT         = "<<"
	RSHIFT         = ">>"
	AND            = "&"
	OR             = "|"
	NOT            = "~"
	XOR            = "^"

	EQUAL        = "=="
	NOTEQUAL     = "!="
	GREATER      = ">"
	LESSER       = "<"
	GREATEREQUAL = ">="
	LESSEREQUAL  = "<="

	// Delimeters
	DOT       = "."
	COMMA     = ","
	SEMICOLON = ";"
	QUOTATION = "\""

	LBRACK = "["
	RBRACK = "]"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	LET   = "let"
	CONST = "const"
	TYPE  = "type"

	STRUCT   = "struct"
	FUNCTION = "fn"
	RETURN   = "return"

	FOR      = "for"
	WHILE    = "while"
	SWITCH   = "switch"
	CASE     = "case"
	BREAK    = "break"
	CONTINUE = "continue"

	IF   = "if"
	ELSE = "else"

	TRUE  = "true"
	FALSE = "false"

	// Basic Types (keywords)
	INT    = "int"
	UINT   = "uint"
	INT8   = "int8"
	UINT8  = "uint8"
	INT16  = "int16"
	UINT16 = "uint16"
	INT32  = "int32"
	UINT32 = "uint32"
	INT64  = "int64"
	UINT64 = "uint64"
	FLOAT  = "float"
	DOUBLE = "double"
	BOOL   = "bool"

	// Identifiers
	IDENT = "ident" // identifier

)

var keywords = map[string]TokenType{
	"let":      LET,
	"const":    CONST,
	"type":     TYPE,
	"struct":   STRUCT,
	"fn":       FUNCTION,
	"return":   RETURN,
	"for":      FOR,
	"while":    WHILE,
	"switch":   SWITCH,
	"case":     CASE,
	"break":    BREAK,
	"continue": CONTINUE,
	"if":       IF,
	"else":     ELSE,
	"true":     TRUE,
	"false":    FALSE,
	"int":      INT,
	"uint":     UINT,
	"int8":     INT8,
	"uint8":    UINT8,
	"int16":    INT16,
	"uint16":   UINT16,
	"int32":    INT32,
	"uint32":   UINT32,
	"int64":    INT64,
	"uint64":   UINT64,
	"float":    FLOAT,
	"double":   DOUBLE,
	"bool":     BOOL,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
