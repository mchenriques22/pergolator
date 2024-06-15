package jen

import "github.com/dave/jennifer/jen"

func Join(codes []*jen.Statement, sep *jen.Statement) *jen.Statement {
	out := jen.Empty()
	for i, code := range codes {
		if i > 0 {
			out.Add(sep)
		}
		out.Add(code)
	}

	return out
}
