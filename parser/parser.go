package parser

import (
	"monkey-lang-interpreted/ast"
	"monkey-lang-interpreted/lexer"
	"monkey-lang-interpreted/token"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	/*
	 read two tokens, so currentToken and peekToken are both set.
	 the first nextToken call will set currentToken as nil and peekToken
	 as the first token from the lexer.
	 the second nextToken call will set currentToken as the peekToken (first token)
	 and the peekToken will be the second available token
	*/
	p.nextToken()
	p.nextToken()

	return p
}

/*
increments currentToken and peekToken from lexer. Need the peekToken
when the currentToken does not give enough info regarding on how to process it.
ex) 5;
We need see INT 5 as the currentToken but we don't know if it is the end of the line
or the start of an expression if we couldn't "peek" into the next token and see that
it is a semi-colon, indicating that it's the end of a line.
*/
func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
