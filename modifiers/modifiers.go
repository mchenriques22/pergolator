package modifiers

import (
	"github.com/iancoleman/strcase"

	"github.com/mchenriques22/pergolator/tree"
)

func FormatKeys(format func(string) string) tree.Modifiers {
	return func(expr tree.Expr) tree.Expr {
		return formatKeys(expr, format)
	}
}

func formatKeys(expr tree.Expr, format func(string) string) tree.Expr {
	switch expr := expr.(type) {
	case *tree.Query:
		expr.Key = format(expr.Key)
	case *tree.And:
		for i, child := range expr.Children {
			expr.Children[i] = formatKeys(child, format)
		}
	case *tree.Or:
		for i, child := range expr.Children {
			expr.Children[i] = formatKeys(child, format)
		}
	case *tree.Not:
		expr.Child = formatKeys(expr.Child, format)
	}

	return expr
}

func FormatValues(format func(string) string) tree.Modifiers {
	return func(expr tree.Expr) tree.Expr {
		return formatValues(expr, format)
	}
}

func formatValues(expr tree.Expr, format func(string) string) tree.Expr {
	switch expr := expr.(type) {
	case *tree.Query:
		expr.Value = format(expr.Value)
	case *tree.And:
		for i, child := range expr.Children {
			expr.Children[i] = formatValues(child, format)
		}
	case *tree.Or:
		for i, child := range expr.Children {
			expr.Children[i] = formatValues(child, format)
		}
	case *tree.Not:
		expr.Child = formatValues(expr.Child, format)
	}

	return expr
}

func IgnoreSomeKeys(keys []string) tree.Modifiers {
	return func(expr tree.Expr) tree.Expr {
		return ignoreSomeKeys(expr, keys)
	}
}

func ignoreSomeKeys(expr tree.Expr, keys []string) tree.Expr {
	switch expr := expr.(type) {
	case *tree.Query:
		for _, key := range keys {
			if expr.Key == key {
				return &tree.Empty{}
			}
		}
	case *tree.And:
		var children []tree.Expr
		for _, child := range expr.Children {
			newChild := ignoreSomeKeys(child, keys)
			if _, isEmpty := newChild.(*tree.Empty); !isEmpty {
				children = append(children, newChild)
			}
		}
		expr.Children = children
	case *tree.Or:
		var children []tree.Expr
		for _, child := range expr.Children {
			newChild := ignoreSomeKeys(child, keys)
			if _, isEmpty := newChild.(*tree.Empty); !isEmpty {
				children = append(children, newChild)
			}
		}
		expr.Children = children
	case *tree.Not:
		newChild := ignoreSomeKeys(expr.Child, keys)
		if _, isEmpty := newChild.(*tree.Empty); !isEmpty {
			expr.Child = newChild
		} else {
			return &tree.Empty{}
		}
	}

	return expr
}

func FormatKeysToSnakeCase(expr tree.Expr) tree.Expr {
	return formatKeys(expr, strcase.ToSnake)
}

func FormatKeysToCamelCase(expr tree.Expr) tree.Expr {
	return formatKeys(expr, strcase.ToCamel)
}
