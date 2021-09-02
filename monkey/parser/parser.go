package parser

import (
	"fmt"

	"github.com/tMinamiii/various-parser/monkey/ast"
	"github.com/tMinamiii/various-parser/monkey/lexer"
	"github.com/tMinamiii/various-parser/monkey/mtoken"
)

// 5 + 5 * 10のように、「+」の後に別の演算子式が続く可能性があ
// るからだ。これには後ほど取り組み、式の構文解析について詳しく見ていくことにする。これがこの構
// 文解析器の中でおそらく最も複雑で、最も美しい部分だ
type Parser struct {
	l *lexer.Lexer

	errors    []string
	curToken  mtoken.Token
	peekToken mtoken.Token
}

func NewParser(l *lexer.Lexer) *Parser {
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
func (p *Parser) peekError(t mtoken.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != mtoken.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case mtoken.LET:
		return p.parseLetStatement()
	case mtoken.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}

}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	//TODO: セミコロンに遭遇するまで式を読み飛ばしてしまっている
	for !p.curTokenIs(mtoken.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(mtoken.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(mtoken.ASSIGN) {
		return nil
	}

	//TODO: セミコロンに遭遇するまで式を読み飛ばしてしまっている
	for !p.curTokenIs(mtoken.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t mtoken.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t mtoken.TokenType) bool {
	return p.peekToken.Type == t
}

// 後続するトークンにアサーションを設けつつトークンを進める
func (p *Parser) expectPeek(t mtoken.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	// 次のトークンが期待に合わない場合に自動的にエラーを追加するようにできる
	p.peekError(t)
	return false
}
