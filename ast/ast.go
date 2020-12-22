package ast

import "monkey/token"

// Node : の基底interface
type Node interface {
	TokenLiteral() string
}

// Statement : 文を表現するinterface
type Statement interface {
	Node
	statementNode()
}

// Expression : 式を表現するinterface
type Expression interface {
	Node
	expressionNode()
}

// Program : プログラム全体を表現する構造体
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

// LetStatement : let文を表現する構造体
type LetStatement struct {
	Token token.Token
	Name  *Identifier
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
