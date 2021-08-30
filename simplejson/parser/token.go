package simplejson

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	STRING  = "STRING"

	// 識別子 + リテラル
	INT = "INT" // 1341412

	// デリミタ
	COMMA = ","
	COLON = ":"

	L_PAREN = "("
	R_PAREN = ")"
	L_BRACE = "{"
	R_BRACE = "}"

	// キーワード
	TRUE  = "TRUE"
	FALSE = "FALSE"
)

var keywords = map[string]TokenType{
	"true":  TRUE,
	"false": FALSE,
}

func Lookup(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return STRING
}
