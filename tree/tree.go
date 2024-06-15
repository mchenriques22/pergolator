package tree

import (
	"fmt"
)

type ParseFn func(query string) (Expr, error)

type Expr interface {
	String() string
}

type Sign string

const (
	Eq  Sign = ":"
	Lt  Sign = ":<"
	Lte Sign = ":<="
	Gt  Sign = ":>"
	Gte Sign = ":>="
)

type Query struct {
	Key   string
	Sign  Sign
	Value string
}

func (x *Query) String() string {
	return x.Key + string(x.Sign) + x.Value
}

type And struct {
	Children []Expr
}

func (x *And) String() string {
	out := format(x.Children[0])
	for _, child := range x.Children[1:] {
		out = fmt.Sprintf("%s AND %s", out, format(child))
	}

	return out
}

type Or struct {
	Children []Expr
}

func (x *Or) String() string {
	out := format(x.Children[0])
	for _, child := range x.Children[1:] {
		out = fmt.Sprintf("%s OR %s", out, format(child))
	}

	return out
}

type Not struct {
	Child Expr
}

func (x *Not) String() string {
	return fmt.Sprintf("NOT(%s)", format(x.Child))
}

func format(expr Expr) string {
	switch expr.(type) {
	case *And, *Query:
		return expr.String()
	}

	return fmt.Sprintf("(%s)", expr)
}

type Empty struct{}

func (e *Empty) String() string {
	return "Empty struct, should be ignored"
}

type Modifiers func(Expr) Expr
