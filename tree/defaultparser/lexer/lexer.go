package lexer

import (
	"unicode"

	"github.com/mchenriques22/pergolator/tree/defaultparser/token"
)

// Lexer performs lexical analysis/scanning of the input string
type Lexer struct {
	Input        []rune
	char         rune // current char
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
}

// New creates and returns a pointer to the Lexer
func New(input string) *Lexer {
	l := &Lexer{Input: []rune(input)}
	l.readChar()
	return l
}

// NewFromRunes creates and returns a pointer to the Lexer
func NewFromRunes(input []rune) *Lexer {
	l := &Lexer{Input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.Input) {
		// End of input (haven't read anything yet or EOF)
		// 0 is ASCII code for "NUL" character
		l.char = 0
	} else {
		l.char = l.Input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

// NextToken switches through the lexer's current char and creates a new token.
// It then it calls readChar() to advance the lexer and it returns the token
func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.char {
	case '-':
		t = newToken(token.NOT, l.position, l.position+1, l.char)
	case '(':
		t = newToken(token.LeftParenthesis, l.position, l.position+1, l.char)
	case ')':
		t = newToken(token.RightParenthesis, l.position, l.position+1, l.char)
	case '"':
		t.Start = l.position
		t.Literal = l.readString()
		t.End = l.position
		t.Type = token.String
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		t.Start = l.position
		ident := l.readIdentifier()
		t.Literal = ident
		t.End = l.position
		t.Type = token.LookupIdentifier(ident)
		return t
	}

	l.readChar()
	return t
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.char) {
		l.readChar()
	}
}

func newToken(tokenType token.Type, start, end int, char ...rune) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(char),
		Start:   start,
		End:     end,
	}
}

// readString sets a start position and reads through characters
// When it finds a closing `"`, it stops consuming characters and
// returns the string between the start and end positions.
func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		prevChar := l.char
		l.readChar()
		if (l.char == '"' && prevChar != '\\') || l.char == 0 {
			break
		}
	}
	return string(l.Input[position:l.position])
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	lexingSign := l.char == ':'

	continueFn := func(c rune) bool {
		if lexingSign {
			return c == ':' || unicode.IsSymbol(c)
		}
		return !unicode.IsSpace(c) && c != '(' && c != ')' && c != 0 && c != ':'
	}

	for continueFn(l.char) {
		l.readChar()
	}

	return string(l.Input[position:l.position])
}
