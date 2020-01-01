/*
Package ast creates an AST out of statements and expressions
*/
package ast

// NOTE: All tokenLiteral methods are used only for testing/debugging

import "github.com/dorin131/dorin-script/token"

// Program : holds all the statements in the program
type Program struct {
	Statements []Statement
}

// Statement : something that doesnt return a value
// usually on the left side of the assignment token
type Statement interface {
	Node
	statementNode() // dummy method
}

// Expression : something that returns a value
type Expression interface {
	Node
	expressionNode() // dummy method
}

// Node : the Program, Statement and Expression are all nodes
type Node interface {
	TokenLiteral() string
}

// TokenLiteral : returns the literal value of the first statement token
// this is basically how we het the literal value of the top token
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement : contains the token type (LET), the name of the variable
// and its value
type LetStatement struct {
	Token token.Token // token.LET token
	Name *Identifier
	Value Expression
}

// StatementNode : dummy method
func (ls *LetStatement) StatementNode() {}

// TokenLiteral : get the literal value for the LET token
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier : stores an IDENT token and its value
type Identifier struct {
	Token token.Token // token.IDENT token
	Value string
}

// expressionNode : dummy method
// it is an expression because on its own it returns a value
func (i *Identifier) expressionNode() {}

// TokenLiteral : get the token literal value of an Identifier
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}