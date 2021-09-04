package ast

import (
	"testing"

	"github.com/tMinamiii/various-parser/monkey/mtoken"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: mtoken.Token{Type: mtoken.LET, Literal: "let"},
				Name: &Identifier{
					Token: mtoken.Token{Type: mtoken.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: mtoken.Token{Type: mtoken.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
