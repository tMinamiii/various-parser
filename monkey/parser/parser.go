package monkey

import "fmt"

// 5 + 5 * 10のように、「+」の後に別の演算子式が続く可能性があ
// るからだ。これには後ほど取り組み、式の構文解析について詳しく見ていくことにする。これがこの構
// 文解析器の中でおそらく最も複雑で、最も美しい部分だ
type Parser struct {
	l *Lexer

	errors    []string
	curToken  Token
	peekToken Token
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	// 2つトークンを読み込み。curTokenとpeekTokenの両方がセット
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

// 次のトークンが期待しているものでなければp.errorsにメッセージを詰める
func (p *Parser) peekError(t TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	for p.curToken.Type != EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case LET:
		return p.parseLetStatement()
	default:
		return nil
	}

}

func (p *Parser) parseLetStatement() *LetStatement {
	stmt := &LetStatement{Token: p.curToken}
	if !p.expectPeek(IDENT) {
		return nil
	}
	stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(ASSIGN) {
		return nil
	}

	//TODO: セミコロンに遭遇するまで式を読み飛ばしてしまっている
	for !p.curTokenIs(SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t TokenType) bool {
	return p.peekToken.Type == t
}

// 後続するトークンにアサーションを設けつつトークンを進める
func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	// 次のトークンが期待に合わない場合に自動的にエラーを追加するようにできる
	p.peekError(t)
	return false
}
