package lexer

import (
	"interpreter-go/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LEFT_PAREN, "("},
		{token.RIGHT_PAREN, ")"},
		{token.LEFT_BRACE, "{"},
		{token.RIGHT_BRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, testToken := range tests {
		resultToken := lexer.NextToken()
		if resultToken.Type != testToken.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, testToken.expectedType, resultToken.Type)
		}

		if resultToken.Literal != testToken.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, testToken.expectedLiteral, resultToken.Literal)
		}

	}
}
