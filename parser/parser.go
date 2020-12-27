package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

// Parser : Lexerで字句解析した結果から構文解析を行う
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token // 現在のTokenを指す
	peekToken token.Token // 次のTokenを指す
}

// New : 構文解析器のインスタンスを作成する
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// curTokenとpeekTokenに値をセットする
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram : Lexerのインスタンスが持っているToken列を解析する
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
