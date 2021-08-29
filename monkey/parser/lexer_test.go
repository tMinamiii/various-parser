package monkey

import (
	"reflect"
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	type fields struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		want   []Token
	}{
		{
			name: "case1",
			fields: fields{
				input: `=+(){},;`,
			},
			want: []Token{
				{Type: ASSIGN, Literal: "="},
				{Type: PLUS, Literal: "+"},
				{Type: L_PAREN, Literal: "("},
				{Type: R_PAREN, Literal: ")"},
				{Type: L_BRACE, Literal: "{"},
				{Type: R_BRACE, Literal: "}"},
				{Type: COMMA, Literal: ","},
				{Type: SEMICOLON, Literal: ";"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			name: "case2",
			fields: fields{
				input: `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 >5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10
10 != 9
`,
			},
			want: []Token{
				{Type: LET, Literal: "let"},
				{Type: IDENT, Literal: "five"},
				{Type: ASSIGN, Literal: "="},
				{Type: INT, Literal: "5"},
				{Type: SEMICOLON, Literal: ";"},
				{Type: LET, Literal: "let"},
				{Type: IDENT, Literal: "ten"},
				{Type: ASSIGN, Literal: "="},
				{Type: INT, Literal: "10"},
				{Type: SEMICOLON, Literal: ";"},
				{Type: LET, Literal: "let"},
				{Type: IDENT, Literal: "add"},
				{Type: ASSIGN, Literal: "="},
				{Type: FUNCTION, Literal: "fn"},
				{Type: L_PAREN, Literal: "("},
				{Type: IDENT, Literal: "x"},
				{Type: COMMA, Literal: ","},
				{Type: IDENT, Literal: "y"},
				{Type: R_PAREN, Literal: ")"},
				{Type: L_BRACE, Literal: "{"},
				{Type: IDENT, Literal: "x"},
				{Type: PLUS, Literal: "+"},
				{Type: IDENT, Literal: "y"},
				{Type: SEMICOLON, Literal: ";"},
				{Type: R_BRACE, Literal: "}"},
				{Type: SEMICOLON, Literal: ";"},
				{Type: LET, Literal: "let"},
				{Type: IDENT, Literal: "result"},
				{Type: ASSIGN, Literal: "="},
				{Type: IDENT, Literal: "add"},
				{Type: L_PAREN, Literal: "("},
				{Type: IDENT, Literal: "five"},
				{Type: COMMA, Literal: ","},
				{Type: IDENT, Literal: "ten"},
				{Type: R_PAREN, Literal: ")"},
				{Type: SEMICOLON, Literal: ";"},
				{Type: BANG, Literal: "!"},
				{Type: MINUS, Literal: "-"},
				{Type: SLASH, Literal: "/"},
				{Type: ASTERISK, Literal: "*"},
				{Type: INT, Literal: "5"},
				{Type: SEMICOLON, Literal: ";"},
				{Type: INT, Literal: "5"},
				{Type: LT, Literal: "<"},
				{Type: INT, Literal: "10"},
				{Type: GT, Literal: ">"},
				{Type: INT, Literal: "5"},
				{Type: SEMICOLON, Literal: ";"},
				{Type: IF, Literal: "if"},
				{Type: L_PAREN, Literal: "("},
				{Type: INT, Literal: "5"},
				{Type: LT, Literal: "<"},
				{Type: INT, Literal: "10"},
				{Type: R_PAREN, Literal: ")"},
				{Type: L_BRACE, Literal: "{"},
				{Type: RETURN, Literal: "return"},
				{Type: TRUE, Literal: "true"},
				{Type: SEMICOLON, Literal: ";"},
				{Type: R_BRACE, Literal: "}"},
				{Type: ELSE, Literal: "else"},
				{Type: L_BRACE, Literal: "{"},
				{Type: RETURN, Literal: "return"},
				{Type: FALSE, Literal: "false"},
				{Type: SEMICOLON, Literal: ";"},
				{Type: R_BRACE, Literal: "}"},
				{Type: INT, Literal: "10"},
				{Type: EQ, Literal: "=="},
				{Type: INT, Literal: "10"},
				{Type: INT, Literal: "10"},
				{Type: NOT_EQ, Literal: "!="},
				{Type: INT, Literal: "9"},
				{Type: EOF, Literal: ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.fields.input)
			for i, w := range tt.want {
				if got := l.NextToken(); !reflect.DeepEqual(got, w) {
					t.Errorf("tests[%d] - Lexer.NextToken() = %v, want %v", i, got, w)
				}
			}
		})
	}
}