package build

type Mode = int

const (
	_ Mode = iota
	EXECUTE
	AST
	TOKENS
)
