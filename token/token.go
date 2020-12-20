package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// 言語内のキーワードを定義するmap
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent : 入力に該当するキーワードがあれば返却する。なければ識別子として扱う。
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子 + リテラル
	IDENT = "IDENT"
	INT   = "INT"

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
