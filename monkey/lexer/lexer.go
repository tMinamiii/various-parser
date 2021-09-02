package lexer

import "github.com/tMinamiii/various-parser/monkey/mtoken"

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置
	readPosition int  // 現在の文字の次
	ch           byte // 現在操作中の文字
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:        input,
		position:     0,
		readPosition: 1,
		ch:           input[0],
	}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 終端
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) newToken(tokenType mtoken.TokenType, ch byte) mtoken.Token {
	return mtoken.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() mtoken.Token {
	var tok mtoken.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = mtoken.Token{Type: mtoken.EQ, Literal: literal}
		} else {
			tok = l.newToken(mtoken.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = mtoken.Token{Type: mtoken.NOT_EQ, Literal: literal}
		} else {
			tok = l.newToken(mtoken.BANG, l.ch)
		}
	case '+':
		tok = l.newToken(mtoken.PLUS, l.ch)
	case '-':
		tok = l.newToken(mtoken.MINUS, l.ch)
	case '/':
		tok = l.newToken(mtoken.SLASH, l.ch)
	case '*':
		tok = l.newToken(mtoken.ASTERISK, l.ch)
	case '<':
		tok = l.newToken(mtoken.LT, l.ch)
	case '>':
		tok = l.newToken(mtoken.GT, l.ch)
	case ';':
		tok = l.newToken(mtoken.SEMICOLON, l.ch)
	case '(':
		tok = l.newToken(mtoken.L_PAREN, l.ch)
	case ')':
		tok = l.newToken(mtoken.R_PAREN, l.ch)
	case '{':
		tok = l.newToken(mtoken.L_BRACE, l.ch)
	case '}':
		tok = l.newToken(mtoken.R_BRACE, l.ch)
	case ',':
		tok = l.newToken(mtoken.COMMA, l.ch)
	case 0:
		tok = mtoken.Token{Type: mtoken.EOF, Literal: ""}
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = mtoken.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = mtoken.INT
			return tok
		} else {
			tok = l.newToken(mtoken.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readNumber() string {
	var number []rune
	for isDigit(l.ch) {
		number = append(number, rune(l.ch))
		l.readChar()
	}
	return string(number)
}

func (l *Lexer) readIdentifier() string {
	var ident []rune
	// 文字(a-z, A-Z, _)が続く限り読み進める
	for isLetter(l.ch) {
		ident = append(ident, rune(l.ch))
		l.readChar()
	}
	return string(ident)
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' ||
		'A' <= ch && ch <= 'Z' ||
		ch == '_'
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
