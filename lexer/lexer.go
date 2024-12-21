package lexer

import (
	"interpreter-go/token"
	"log"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current character (ch variable))
	readPosition int  // current reading position in input (after current character)
	ch           byte // current character under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// gives us the next character and advance our position in the input string
func (l *Lexer) readChar() {
	/*
		Checks if we've reached the input of the input. If so, assign 0 to ch which in
		ASCII is the "NUL" character. The "NUL" / 0 will signal we finished reading and lexing
		the input (end of file) or we didn't start yet.
	*/
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else { // if not at the end of our input, set the ch value to the next character
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	log.Printf("l.ch: %v", l.ch)
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LEFT_PAREN, l.ch)
	case ')':
		tok = newToken(token.RIGHT_PAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LEFT_BRACE, l.ch)
	case '}':
		tok = newToken(token.RIGHT_BRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
