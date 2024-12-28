package ast

import "monkey-lang-interpreted/token"

// abstract syntax tree

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

/*
Program node is the room of every AST parse produces. Every valid
Monkey program is a series of statements. Statements will be contained
in Program.Statements which is a slice of AST nodes that implement the
Statement interface.
*/
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifer  // name of the identifier (variable name)
	Value Expression  // the expression that computes the value
}

// methods used for debugging purposes
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifer struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifer) expressionNode() {}
func (i *Identifer) TokenLiteral() string {
	return i.Token.Literal
}
