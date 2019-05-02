package parser

import (
	"programminglanguage/ast"
	"programminglanguage/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// Operator Precendence Levels
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

var precendences = map[token.TokenType]int{
	token.EQUAL:          EQUALS,
	token.NOTEQUAL:       EQUALS,
	token.LESSER:         LESSGREATER,
	token.GREATER:        LESSGREATER,
	token.PLUS:           SUM,
	token.MINUS:          SUM,
	token.DIVISION:       PRODUCT,
	token.MULTIPLICATION: PRODUCT,
	token.LPAREN:         CALL,
}
