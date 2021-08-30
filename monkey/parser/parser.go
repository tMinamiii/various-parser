package monkey

type Parser struct {
	l *Lexer

	curToken  Token
	peekToken Token
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{l: l}
	// 2つトークンを読み込み。curTokenとpeekTokenの両方がセット
	p.curToken = p.l.NextToken()
	p.peekToken = p.l.NextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *Program {
	return nil
}
