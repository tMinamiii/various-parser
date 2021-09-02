package lexer

import (
	"reflect"
	"testing"

	"github.com/tMinamiii/various-parser/monkey/mtoken"
)

func TestLexer_NextToken(t *testing.T) {
	type fields struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		want   []mtoken.Token
	}{
		{
			name: "case1",
			fields: fields{
				input: `=+(){},;`,
			},
			want: []mtoken.Token{
				{Type: mtoken.ASSIGN, Literal: "="},
				{Type: mtoken.PLUS, Literal: "+"},
				{Type: mtoken.L_PAREN, Literal: "("},
				{Type: mtoken.R_PAREN, Literal: ")"},
				{Type: mtoken.L_BRACE, Literal: "{"},
				{Type: mtoken.R_BRACE, Literal: "}"},
				{Type: mtoken.COMMA, Literal: ","},
				{Type: mtoken.SEMICOLON, Literal: ";"},
				{Type: mtoken.EOF, Literal: ""},
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
			want: []mtoken.Token{
				{Type: mtoken.LET, Literal: "let"},
				{Type: mtoken.IDENT, Literal: "five"},
				{Type: mtoken.ASSIGN, Literal: "="},
				{Type: mtoken.INT, Literal: "5"},
				{Type: mtoken.SEMICOLON, Literal: ";"},
				{Type: mtoken.LET, Literal: "let"},
				{Type: mtoken.IDENT, Literal: "ten"},
				{Type: mtoken.ASSIGN, Literal: "="},
				{Type: mtoken.INT, Literal: "10"},
				{Type: mtoken.SEMICOLON, Literal: ";"},
				{Type: mtoken.LET, Literal: "let"},
				{Type: mtoken.IDENT, Literal: "add"},
				{Type: mtoken.ASSIGN, Literal: "="},
				{Type: mtoken.FUNCTION, Literal: "fn"},
				{Type: mtoken.L_PAREN, Literal: "("},
				{Type: mtoken.IDENT, Literal: "x"},
				{Type: mtoken.COMMA, Literal: ","},
				{Type: mtoken.IDENT, Literal: "y"},
				{Type: mtoken.R_PAREN, Literal: ")"},
				{Type: mtoken.L_BRACE, Literal: "{"},
				{Type: mtoken.IDENT, Literal: "x"},
				{Type: mtoken.PLUS, Literal: "+"},
				{Type: mtoken.IDENT, Literal: "y"},
				{Type: mtoken.SEMICOLON, Literal: ";"},
				{Type: mtoken.R_BRACE, Literal: "}"},
				{Type: mtoken.SEMICOLON, Literal: ";"},
				{Type: mtoken.LET, Literal: "let"},
				{Type: mtoken.IDENT, Literal: "result"},
				{Type: mtoken.ASSIGN, Literal: "="},
				{Type: mtoken.IDENT, Literal: "add"},
				{Type: mtoken.L_PAREN, Literal: "("},
				{Type: mtoken.IDENT, Literal: "five"},
				{Type: mtoken.COMMA, Literal: ","},
				{Type: mtoken.IDENT, Literal: "ten"},
				{Type: mtoken.R_PAREN, Literal: ")"},
				{Type: mtoken.SEMICOLON, Literal: ";"},
				{Type: mtoken.BANG, Literal: "!"},
				{Type: mtoken.MINUS, Literal: "-"},
				{Type: mtoken.SLASH, Literal: "/"},
				{Type: mtoken.ASTERISK, Literal: "*"},
				{Type: mtoken.INT, Literal: "5"},
				{Type: mtoken.SEMICOLON, Literal: ";"},
				{Type: mtoken.INT, Literal: "5"},
				{Type: mtoken.LT, Literal: "<"},
				{Type: mtoken.INT, Literal: "10"},
				{Type: mtoken.GT, Literal: ">"},
				{Type: mtoken.INT, Literal: "5"},
				{Type: mtoken.SEMICOLON, Literal: ";"},
				{Type: mtoken.IF, Literal: "if"},
				{Type: mtoken.L_PAREN, Literal: "("},
				{Type: mtoken.INT, Literal: "5"},
				{Type: mtoken.LT, Literal: "<"},
				{Type: mtoken.INT, Literal: "10"},
				{Type: mtoken.R_PAREN, Literal: ")"},
				{Type: mtoken.L_BRACE, Literal: "{"},
				{Type: mtoken.RETURN, Literal: "return"},
				{Type: mtoken.TRUE, Literal: "true"},
				{Type: mtoken.SEMICOLON, Literal: ";"},
				{Type: mtoken.R_BRACE, Literal: "}"},
				{Type: mtoken.ELSE, Literal: "else"},
				{Type: mtoken.L_BRACE, Literal: "{"},
				{Type: mtoken.RETURN, Literal: "return"},
				{Type: mtoken.FALSE, Literal: "false"},
				{Type: mtoken.SEMICOLON, Literal: ";"},
				{Type: mtoken.R_BRACE, Literal: "}"},
				{Type: mtoken.INT, Literal: "10"},
				{Type: mtoken.EQ, Literal: "=="},
				{Type: mtoken.INT, Literal: "10"},
				{Type: mtoken.INT, Literal: "10"},
				{Type: mtoken.NOT_EQ, Literal: "!="},
				{Type: mtoken.INT, Literal: "9"},
				{Type: mtoken.EOF, Literal: ""},
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
