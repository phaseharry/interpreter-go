package lexer

import (
	"interpreter-go/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LEFT_PAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RIGHT_PAREN, ")"},
		{token.LEFT_BRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RIGHT_BRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LEFT_PAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RIGHT_PAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LESS_THAN, "<"},
		{token.INT, "10"},
		{token.GREATER_THAN, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
	}
	lexer := New(input)

	for i, testToken := range tests {
		resultToken := lexer.NextToken()
		// log.Printf("resultToken: %v, testToken: %v", resultToken, testToken)
		if resultToken.Type != testToken.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, testToken.expectedType, resultToken.Type)
		}

		if resultToken.Literal != testToken.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, testToken.expectedLiteral, resultToken.Literal)
		}

	}
}
