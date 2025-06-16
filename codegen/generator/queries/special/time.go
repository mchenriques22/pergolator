package special

import (
	"go/types"

	. "github.com/dave/jennifer/jen"

	"github.com/mchenriques22/pergolator/codegen/generator/queries/shared"
	"github.com/mchenriques22/pergolator/codegen/utils"
)

func GenerateTimeQuery(typeName *types.TypeName, optionsGetter utils.FieldOptionsGetter) *Statement {
	caseStmts := Empty()

	// 	parsed, err := time.Parse(time.RFC3339, query.Value)
	//	if err != nil {
	//		return pTimeFalseFn
	//	}
	caseStmts.
		List(Id("parsed"), Err()).Op(":=").Qual("time", "Parse").Call(Qual("time", "RFC3339"), Id("query").Dot("Value")).
		Line().If(Err().Op("!=").Nil()).Block(shared.ReturnFalseQuery(typeName)).
		Line()

	// switch query.Sign {
	caseStmts.Switch(Id("query").Dot("Sign")).Block(
		//	case tree.Eq:
		//		return func(t *time.Time) bool {
		//			return t.Equal(parsed)
		//		}
		Case(Id("tree").Dot("Eq")).Block(Return(Func().Params(Id("t").Op("*").Qual("time", "Time")).Bool().Block(
			Return().Id("t").Dot("Equal").Call(Id("parsed")),
		))),

		// 	case tree.Lt:
		//		return func(t *time.Time) bool {
		//			return t.Before(parsed)
		//		}
		Case(Id("tree").Dot("Lt")).Block(Return(Func().Params(Id("t").Op("*").Qual("time", "Time")).Bool().Block(
			Return().Id("t").Dot("Before").Call(Id("parsed")),
		))),

		// 	case tree.Lte:
		//		return func(t *time.Time) bool {
		//			return t.Before(parsed) || t.Equal(parsed)
		//		}
		Case(Id("tree").Dot("Lte")).Block(Return(Func().Params(Id("t").Op("*").Qual("time", "Time")).Bool().Block(
			Return().Id("t").Dot("Before").Call(Id("parsed")).Op("||").Id("t").Dot("Equal").Call(Id("parsed")),
		))),

		// 	case tree.Gt:
		//		return func(t *time.Time) bool {
		//			return parsed.Before(*t)
		//		}
		Case(Id("tree").Dot("Gt")).Block(Return(Func().Params(Id("t").Op("*").Qual("time", "Time")).Bool().Block(
			Return().Id("parsed").Dot("Before").Call(Op("*").Id("t")),
		))).Line(),

		// 	case tree.Gte:
		//		return func(t *time.Time) bool {
		//			return parsed.Before(*t) || parsed.Equal(*t)
		//		}
		Case(Id("tree").Dot("Gte")).Block(Return(Func().Params(Id("t").Op("*").Qual("time", "Time")).Bool().Block(
			Return().Id("parsed").Dot("Before").Call(Op("*").Id("t")).Op("||").Id("parsed").Dot("Equal").Call(Op("*").Id("t")),
		))),
	)

	// return pTimeFalseFn
	caseStmts.Add(Line(), shared.ReturnFalseQuery(typeName))

	return caseStmts
}
