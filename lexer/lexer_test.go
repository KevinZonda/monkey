package lexer

import (
	"fmt"
	"github.com/KevinZonda/monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
   x + y;
};
let result = add(five, ten);
 `
	tests := []struct {
		expectedType    token.TokenKind
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTITY, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTITY, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTITY, "add"},
		{token.ASSIGN, "="},
		{token.FUNC, "fn"},
		{token.L_PAREN, "("},
		{token.IDENTITY, "x"},
		{token.COMMA, ","},
		{token.IDENTITY, "y"},
		{token.R_PAREN, ")"},
		{token.L_BRACE, "{"},
		{token.IDENTITY, "x"},
		{token.PLUS, "+"},
		{token.IDENTITY, "y"},
		{token.SEMICOLON, ";"},
		{token.R_BRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTITY, "result"},
		{token.ASSIGN, "="},
		{token.IDENTITY, "add"},
		{token.L_PAREN, "("},
		{token.IDENTITY, "five"},
		{token.COMMA, ","},
		{token.IDENTITY, "ten"},
		{token.R_PAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Println(">>", "Kind:", tok.Kind, "Liter:", tok.Literal)
		if tok.Kind != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Kind)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
