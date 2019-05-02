package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Support
	ILLEGAL = "illegal"
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
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	RETURN   = "return"
	LET      = "let"
	FUNCTION = "fn"
	IF       = "if"
	ELSE     = "else"
	BREAK    = "break"
	FOR      = "for"
	WHILE    = "while"
	SWITCH   = "switch"
	CASE     = "case"
	STRUCT   = "struct"
	TRUE     = "true"
	FALSE    = "false"

	// Identifiers and Literals
	IDENT  = "ident" // identifier
	INT    = "int"
	FLOAT  = "float"
	STRING = "string"
	BOOL   = "bool"
)

var keywords = map[string]TokenType{
	"return": RETURN,
	"let":    LET,
	"fn":     FUNCTION,
	"if":     IF,
	"else":   ELSE,
	"break":  BREAK,
	"for":    FOR,
	"while":  WHILE,
	"switch": SWITCH,
	"case":   CASE,
	"struct": STRUCT,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
