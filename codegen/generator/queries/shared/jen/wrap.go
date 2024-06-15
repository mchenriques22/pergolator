package jen

import "github.com/dave/jennifer/jen"

func Wrap(strs []string, wrapper func(v interface{}) *jen.Statement) []jen.Code {
	wrapped := make([]jen.Code, len(strs))
	for i, str := range strs {
		wrapped[i] = wrapper(str)
	}

	return wrapped
}
