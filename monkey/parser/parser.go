package parser

// Pratt構文解析器の考え方で重要なのは、
// * トークンタイプごとに構文解析関数（Prattは「semantic code」と呼ぶ）を関連付けることだ。
// * あるトークンタイプに遭遇するたびに、対応する構文解析関数が呼ばれる。
// * この関数は適切な式を構文解析し、その式を表現するASTノードを返す。
// * トークンタイプごとに、最大2つの構文解析関数が関連付けられる。
// * これらの関数は、トークンが前置で出現したか中置か出現したかによって使い分けられる。
import (
	"fmt"
	"strconv"

	"github.com/tMinamiii/various-parser/monkey/ast"
	"github.com/tMinamiii/various-parser/monkey/lexer"
	"github.com/tMinamiii/various-parser/monkey/mtoken"
)

// LOWESTが優先度MIN CALLが優先度MAX
const (
	_ int = iota // 0をとばす
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X または !X
	CALL        // myFunction(X)
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// 5 + 5 * 10のように、「+」の後に別の演算子式が続く可能性があ
// るからだ。これには後ほど取り組み、式の構文解析について詳しく見ていくことにする。これがこの構
// 文解析器の中でおそらく最も複雑で、最も美しい部分だ
type Parser struct {
	l *lexer.Lexer

	errors    []string
	curToken  mtoken.Token
	peekToken mtoken.Token

	prefixParseFns map[mtoken.TokenType]prefixParseFn // トークンタイプが前置で出現した場合
	infixParseFns  map[mtoken.TokenType]infixParseFn  // トークンタイプが中置で出現した場合
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// マップの初期化し構文解析器を登録する
	p.prefixParseFns = make(map[mtoken.TokenType]prefixParseFn)
	p.registerPrefix(mtoken.IDENT, p.parseIdentifier)
	p.registerPrefix(mtoken.INT, p.parseIntegerLiteral)
	p.registerPrefix(mtoken.BANG, p.parsePrefixExpression)
	p.registerPrefix(mtoken.MINUS, p.parsePrefixExpression)

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
		program.Statements = append(program.Statements, stmt)
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
		return p.parseExpressionStatement()
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

// registerPrefix 前置トークンに対応するパーサーをprefixParseFns格納していく
func (p *Parser) registerPrefix(tokenType mtoken.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType mtoken.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(mtoken.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
func (p *Parser) noPrefixParseFnError(t mtoken.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

// p.curToken.Typeの前置に関連付けられた構文解析関数があるかを確認している
// もし存在していれば、その構文解析関数を呼び出し、その結果を返す。
// そうでなければnilを返す
func (p *Parser) parseExpression(precedence int) ast.Expression {
	if prefix, ok := p.prefixParseFns[p.curToken.Type]; ok {
		leftExp := prefix()
		return leftExp
	}
	p.noPrefixParseFnError(p.curToken.Type)
	return nil
}

func (p *Parser) parseIdentifier() ast.Expression {
	// 単に*ast.Identifierを返すだけ
	// ただし、現在のトークンをTokenフィールドに、トークンのリテラル値をValueフィールド格納する。
	// トークンはすすめない(nextTokenは呼びださない)
	// 構文解析関数に関連付けられたトークンがcurTokenにセットされている状態で動作を開始する。
	// そして、この関数の処理対象である式の一番最後のトークンがcurTokenにセットされた状態になるまで進んで終了する。
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		// int64に変換できない場合
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}
