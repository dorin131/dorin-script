package parser

import (
	"github/com/dorin131/dorin-script/lexer"
	"github/com/dorin131/dorin-script/token"
	"github/com/dorin131/dorin-script/ast"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l :l}

	// read two tokens so that curToken and peekToken are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}