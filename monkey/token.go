package monkey

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子 + リテラル
	IDENT = "IDENT" // add, foobar, x, y ...
	INT   = "INT"   // 1341412

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	L_PAREN = "("
	R_PAREN = ")"
	L_BRACE = "{"
	R_BRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
