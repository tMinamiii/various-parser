package monkey

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置
	readPosition int  // 現在の文字の次
	ch           byte // 現在操作中の文字
}

func New(input string) *Lexer {
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

func (l *Lexer) NextToken() Token {
	var tok Token
	newToken := func(tokenType TokenType, ch byte) Token {
		return Token{Type: tokenType, Literal: string(ch)}
	}
	switch l.ch {
	case '=':
		tok = newToken(ASSIGN, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case '(':
		tok = newToken(L_PAREN, l.ch)
	case ')':
		tok = newToken(R_PAREN, l.ch)
	case '{':
		tok = newToken(L_BRACE, l.ch)
	case '}':
		tok = newToken(R_BRACE, l.ch)
	case '0':
		tok.Literal = ""
		tok.Type = EOF
	}
	l.readChar()
	return tok
}
