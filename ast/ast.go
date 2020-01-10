/*
Package ast creates an AST out of statements and expressions
*/
package ast

// NOTE: All tokenLiteral methods are used only for testing/debugging

import (
	"bytes"

	"github.com/dorin131/dorin-script/token"
)

// Program : holds all the statements in the program
type Program struct {
	Statements []Statement
}

// String : returns all statements as a string
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// TokenLiteral : returns the literal value of the first statement token
// this is basically how we het the literal value of the top token
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// Node : the Program, Statement and Expression are all nodes
type Node interface {
	TokenLiteral() string
	String() string
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

/*-----------------
IDENTIFIER EXPRESSION
-----------------*/

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

func (i *Identifier) String() string {
	return i.Value
}

/*------------------
LET STATEMENT
------------------*/

// LetStatement : contains the token type (LET), the name of the variable
// and its value
type LetStatement struct {
	Token token.Token // token.LET token
	Name  *Identifier
	Value Expression
}

// StatementNode : dummy method
func (ls *LetStatement) statementNode() {}

// TokenLiteral : get the literal value for the LET token
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

/*------------------
RETURN STATEMENT
------------------*/

// ReturnStatement : e.g. "return 1"
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// statementNode : dummy method
func (rs *ReturnStatement) statementNode() {}

// TokenLiteral : as expected
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

/*------------------
EXPRESSION STATEMENT
------------------*/

// ExpressionStatement : is a wrapper for an expression, e.g. "1 + 2;"
type ExpressionStatement struct {
	Token      token.Token // first token of the expression
	Expression Expression
}

// statementNode : dummy method
func (es *ExpressionStatement) statementNode() {}

// TokenLiteral : as expected
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
