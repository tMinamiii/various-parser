package lexer

import (
	"fmt"
	"strconv"
	"strings"
)

type Lexer struct {
	scanner *Scanner
}

func NewLexer(scanner *Scanner) *Lexer {
	return &Lexer{scanner: scanner}
}

func NewLexerWithString(s string) *Lexer {
	return &Lexer{scanner: NewScannerString(s)}
}

func (l *Lexer) GetNextToken() (*Token, error) {
	r, err := l.scanner.consume()
	if err != nil {
		return nil, err
	}

	if r == ' ' || r == '\r' || r == '\t' || r == '\n' {
		return l.GetNextToken()
	}

	switch r {
	case '[':
		return NewToken(TokenLeftBracket), nil
	case ']':
		return NewToken(TokenRightBracket), nil
	case '{':
		return NewToken(TokenLeftBrace), nil
	case '}':
		return NewToken(TokenRightBrace), nil
	case ':':
		return NewToken(TokenColon), nil
	case ',':
		return NewToken(TokenComma), nil
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return l.lexNumber(r)
	case '"':
		return l.lexString()
	case 't':
		return l.lextToken("true", TokenTrue)
	case 'f':
		return l.lextToken("false", TokenTrue)
	case 'n':
		return l.lextToken("null", TokenTrue)
	}

	return nil, fmt.Errorf("Unexpected character:%c", r)

}

func (l *Lexer) lexNumber(first rune) (*Token, error) {
	rs := []rune{first}
	numerics := "0123456789-+e."

	for {
		r, err := l.scanner.peek()
		if err != nil {
			break
		}
		if !strings.Contains(numerics, string(r)) {
			break
		}

		_, _ = l.scanner.consume()
		rs = append(rs, r)
	}

	f, err := strconv.ParseFloat(string(rs), 64)
	if err != nil {
		return nil, err
	}
	return NewTokenNumber(f), nil
}

func (l *Lexer) lexString() (*Token, error) {
	rs := []rune{}
	backslash := false

	for {

	}
}
