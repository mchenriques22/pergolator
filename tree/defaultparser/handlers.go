package defaultparser

import (
	"errors"
	"fmt"

	"github.com/antoninferrand/pergolator/tree"
	. "github.com/antoninferrand/pergolator/tree/defaultparser/token"
)

/*
<CompleteQuery>  ::= <Expression>
<Expression>   	 ::= <Term>{ "OR" <Term>}
<Term>      	 ::= <Factor> "AND" <Term> | <Factor>
<Factor>         ::= <String> | "(" <Expression> ")" | "-"<Factor>
<String>         ::= <Key>:<Value>
*/

var (
	ErrUnexpectedToken    = errors.New("unexpected token")
	ErrShouldBeEOF        = errors.New("should be EOF")
	ErrClosingParenthesis = errors.New("expected closing parenthesis")
	ErrExpectedSign       = errors.New("expected sign")
	ErrExpectedValue      = errors.New("expected value")
	ErrExpectedQuery      = errors.New("expected query")
	ErrUnknownSign        = errors.New("unknown sign")
)

func (p *parser) parse() (tree.Expr, error) {
	result, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if p.peekToken.Type != EOF {
		return nil, wrapError(ErrShouldBeEOF, p.peekToken)
	}

	return result, nil
}

func (p *parser) parseExpression() (tree.Expr, error) {
	expr, err := p.parseTerm()
	if err != nil {
		return nil, err
	}

	for p.peekToken.Type == OR {
		p.nextToken()
		right, err := p.parseTerm()
		if err != nil {
			return nil, err
		}

		orLeftExpr, ok := expr.(*tree.Or)
		if ok {
			orLeftExpr.Children = append(orLeftExpr.Children, right)
			return orLeftExpr, nil
		}

		orRightExpr, ok := right.(*tree.Or)
		if ok {
			orRightExpr.Children = append([]tree.Expr{expr}, orRightExpr.Children...)
			return orRightExpr, nil
		}

		expr = &tree.Or{
			Children: []tree.Expr{
				expr,
				right,
			},
		}
	}

	return expr, nil
}

func (p *parser) parseTerm() (tree.Expr, error) {
	expr, err := p.parseFactor()
	if err != nil {
		return nil, err
	}

	if p.peekToken.Type == AND {
		p.nextToken()
		right, err := p.parseTerm()
		if err != nil {
			return nil, err
		}

		andLeftExpr, ok := expr.(*tree.And)
		if ok {
			andLeftExpr.Children = append(andLeftExpr.Children, right)
			return andLeftExpr, nil
		}

		andRightExpr, ok := right.(*tree.And)
		if ok {
			andRightExpr.Children = append([]tree.Expr{expr}, andRightExpr.Children...)
			return andRightExpr, nil
		}

		expr = &tree.And{
			Children: []tree.Expr{
				expr,
				right,
			},
		}
	}

	return expr, nil
}

func (p *parser) parseFactor() (tree.Expr, error) {
	switch p.peekToken.Type {
	case String:
		return p.parseQuery()

	case LeftParenthesis:
		p.nextToken()
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if p.peekToken.Type != RightParenthesis {
			return nil, wrapError(ErrClosingParenthesis, p.peekToken)
		}
		p.nextToken()
		return expr, nil

	case NOT:
		p.nextToken()
		factor, err := p.parseFactor()
		if err != nil {
			return nil, err
		}
		return &tree.Not{Child: factor}, nil
	}

	return nil, wrapError(ErrUnexpectedToken, p.peekToken)
}

func (p *parser) parseQuery() (tree.Expr, error) {
	if p.peekToken.Type != String {
		return nil, wrapError(ErrExpectedQuery, p.peekToken)
	}

	key := p.peekToken.Literal
	p.nextToken()
	if !IsSign(p.peekToken.Type) {
		return nil, wrapError(ErrExpectedSign, p.peekToken)
	}
	sign, err := lexerSignToTreeSign(p.peekToken.Type)
	if err != nil {
		return nil, err
	}

	p.nextToken()
	if p.peekToken.Type != String {
		return nil, wrapError(ErrExpectedValue, p.peekToken)
	}
	value := p.peekToken.Literal

	p.nextToken()

	return &tree.Query{Key: key, Sign: sign, Value: value}, nil
}

func lexerSignToTreeSign(sign Type) (tree.Sign, error) {
	switch sign {
	case EQ:
		return tree.Eq, nil
	case LT:
		return tree.Lt, nil
	case LTE:
		return tree.Lte, nil
	case GT:
		return tree.Gt, nil
	case GTE:
		return tree.Gte, nil
	}

	return "", ErrUnknownSign
}

func wrapError(err error, token Token) error {
	return fmt.Errorf(`%w, got "%s" between offsets: %d and %d`, err, token.Literal, token.Start, token.End)
}
