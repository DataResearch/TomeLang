package lexer

import (
	"programminglanguage/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	let five = 5; 
	let ten = 10; 

	let add = fn(x, y) { x + y; }; 

	let result = add(five, ten); 

	==`

	lex := New(input)
	target := make([]token.Token, 1)

	for target[len(target)-1].Type != token.EOF {
		target = append(target, lex.NextToken())
	}

	//tests := []struct {
	//	expectedType token.TokenType
	//	expectedLiteral string
	//}{
	//	{token.LET, "let"},
	//	{token.IDENT, "five"},
	//	{token.ASSIGN, "="},
	//	{token.INT, "5"},
	//	{token.SEMICOLON, ";"},
	//	{token.LET, "let"},
	//	{token.IDENT, "ten"},
	//	{token.ASSIGN, "="},
	//	{token.INT, "10"},
	//	{token.SEMICOLON, ";"},
	//	{token.LET, "let"},
	//	{token.IDENT, "add"},
	//	{token.ASSIGN, "="},
	//	{token.FUNCTION, "fn"},
	//	{token.LPAREN, "("},
	//	{token.IDENT, "x"},
	//	{token.COMMA, ","},
	//	{token.IDENT, "y"},
	//	{token.RPAREN, ")"},
	//	{token.LBRACE, "{"},
	//	{token.IDENT, "x"},
	//	{token.PLUS, "+"},
	//	{token.IDENT, "y"},
	//	{token.SEMICOLON, ";"},
	//	{token.RBRACE, "}"},
	//	{token.SEMICOLON, ";"},
	//	{token.LET, "let"},
	//	{token.IDENT, "result"},
	//	{token.ASSIGN, "="},
	//	{token.IDENT, "add"},
	//	{token.LPAREN, "("},
	//	{token.IDENT, "five"},
	//	{token.COMMA, ","},
	//	{token.IDENT, "ten"},
	//	{token.RPAREN, ")"},
	//	{token.SEMICOLON, ";"},
	//	{token.EQUAL, "=="},
	//	{token.EOF, ""},
	//}
	//
	//for i, tt := range tests {
	//
	//	tok := lex.NextToken()
	//
	//	if tt.expectedType != tok.Type {
	//		t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
	//			i, tt.expectedType, tok.Type)
	//	}
	//
	//	if tt.expectedLiteral != tok.Literal {
	//		t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
	//			i, tt.expectedLiteral, tok.Literal)
	//	}
	//}
}
