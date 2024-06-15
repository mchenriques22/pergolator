package token

type Type string

const (
	// Token/character we don't know about
	Illegal Type = "ILLEGAL"

	// End of file
	EOF Type = "EOF"

	// Literals
	String Type = "String"

	// The structural tokens
	LeftParenthesis  Type = "("
	RightParenthesis Type = ")"

	// Operators
	AND Type = "AND"
	OR  Type = "OR"
	NOT Type = "NOT"

	// Signs
	EQ  Type = ":"
	LTE Type = ":<="
	GTE Type = ":>="
	LT  Type = ":<"
	GT  Type = ":>"
)

type Token struct {
	Type    Type
	Literal string
	Start   int
	End     int
}

func LookupIdentifier(ident string) Type {
	switch ident {
	case "AND":
		return AND
	case "OR":
		return OR
	case "-", "NOT":
		return NOT
	case ":", "=", ":=":
		return EQ
	case ":<=":
		return LTE
	case ":>=":
		return GTE
	case ":<":
		return LT
	case ":>":
		return GT
	}

	return String
}

func IsSign(t Type) bool {
	switch t {
	case EQ, LT, LTE, GT, GTE:
		return true
	}

	return false
}
