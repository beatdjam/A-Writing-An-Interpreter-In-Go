package lexer

import "monkey/token"

// Lexer : 字句解析器の構造体
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New : Lexerのコンストラクタ
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// 終端まで文字を読む
	// 0 = nullのような「読まなかった」を表す文字
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken : inputから次のTokenを取り出す
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	// symbol
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// 英字から識別子、キーワードを読み出す
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			return tok
		} else {
			// いずれにも該当しない場合は Illegalとして返す
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// Illegal以外のTokenを作る
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 連続する英字を走査し、識別子を取り出す
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// プログラム内に含みうる英字文字列を定義する
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
