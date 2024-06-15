package tree

import (
	"testing"
)

func TestNewTree(t *testing.T) {

	tests := []struct {
		expr     Expr
		expected string
	}{
		{
			&And{
				Children: []Expr{
					&Query{Key: "strings", Sign: Eq, Value: "my-strings"},
					&Or{
						Children: []Expr{
							&Query{Key: "mynumber", Sign: Eq, Value: "1234"},
							&Query{Key: "meta.match", Sign: Eq, Value: "true"},
						},
					},
				},
			},
			"strings:my-strings AND (mynumber:1234 OR meta.match:true)",
		},
		{
			&Query{Key: "strings", Sign: Eq, Value: "my-strings"},
			"strings:my-strings",
		},
		{
			&Or{
				Children: []Expr{
					&Query{Key: "strings", Sign: Eq, Value: "my-strings"},
					&Query{Key: "mynumber", Sign: Eq, Value: "1234"},
				},
			},
			"strings:my-strings OR mynumber:1234",
		},
		{
			&And{
				Children: []Expr{
					&Query{Key: "strings", Sign: Eq, Value: "my-strings"},
					&And{
						Children: []Expr{
							&Query{Key: "mynumber", Sign: Eq, Value: "1234"},
							&Query{Key: "meta.match", Sign: Eq, Value: "true"},
						},
					},
				},
			},
			"strings:my-strings AND mynumber:1234 AND meta.match:true",
		},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			if test.expected != test.expr.String() {
				t.Errorf("expected %s, got %s", test.expected, test.expr.String())
			}
		})
	}
}
