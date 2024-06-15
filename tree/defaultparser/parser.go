package defaultparser

import (
	"github.com/antoninferrand/pergolator/tree"
	"github.com/antoninferrand/pergolator/tree/defaultparser/lexer"
	. "github.com/antoninferrand/pergolator/tree/defaultparser/token"
)

type parser struct {
	lexer     *lexer.Lexer
	peekToken Token
}

func newParser(l *lexer.Lexer) *parser {
	p := &parser{lexer: l}

	p.nextToken()

	return p
}

func (p *parser) nextToken() {
	p.peekToken = p.lexer.NextToken()
}

// Parse parses tokens and creates an AST. It returns the root node of the AST.
func Parse(query string) (tree.Expr, error) {
	return newParser(lexer.New(query)).parse()
}
