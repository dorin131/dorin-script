/*
Package parser is using the Lexer to tokenize the input
*/
package parser

import (
	"github/com/dorin131/dorin-script/lexer"
	"github/com/dorin131/dorin-script/token"
	"github/com/dorin131/dorin-script/ast"
)

// Parser : initialises a Lexer instance and reads the tokens
type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}

// New : creates a new parser and stores the first two tokens
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l :l}

	// read two tokens so that curToken and peekToken are set
	p.nextToken()
	p.nextToken()

	return p
}

// nextToken : gets the next token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram : parses the program by getting all tokens and creating an AST
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}