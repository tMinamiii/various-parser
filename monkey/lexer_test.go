package monkey

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{ASSIGN, "="},
		{PLUS, "+"},
		{L_PAREN, "("},
		{R_PAREN, ")"},
		{L_BRACE, "{"},
		{R_BRACE, "}"},
		{COMMA, ","},
		{SEMICOLON, ";"},
		{EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
