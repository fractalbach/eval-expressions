// ======================================================================
// Token Structure
// ======================================================================

package token

import (
	"fmt"
)

// ======================================================================
// Token Definitions
// ======================================================================

// Token holds a string and some useful information about that string.
type Token struct {
	Kind    Kind   // the kind of info the token holds.
	Content string // the actual string inside the token.
}

// Kind represents the token's kind in the form of a bitmask.
type Kind int

const (
	EMPTY Kind = 1 << iota
	NUM
	WORD
	ADD
	SUB
	MUL
	DIV
	LEFTPAREN
	RIGHTPAREN
	EQ
)

var runeToKind = map[rune]Kind{
	'(': LEFTPAREN,
	')': RIGHTPAREN,
	'+': ADD,
	'-': SUB,
	'*': MUL,
	'/': DIV,
	'=': EQ,
}

var kindToString = map[Kind]string{
	EMPTY:      "empty",
	NUM:        "num",
	WORD:       "word",
	ADD:        "add",
	SUB:        "sub",
	MUL:        "mul",
	DIV:        "div",
	LEFTPAREN:  "Lparen",
	RIGHTPAREN: "Rparen",
	EQ:         "equals",
}

// ======================================================================
// Token Creation
// ======================================================================

// New creates and returns a new Token.
func New(kind Kind, contents ...string) Token {
	if len(contents) == 0 {
		return Token{Kind: kind}
	}
	contentString := ""
	for _, v := range contents {
		contentString += fmt.Sprint(v)
	}
	return Token{
		Kind:    kind,
		Content: contentString,
	}
}

// NewFromSymbol returns a newly made token that has the correct Kind
// and Content for one of the symbolic tokens.
func NewFromSymbol(r rune) Token {
	kind, ok := runeToKind[r]
	if !ok {
		kind = EMPTY
	}
	return Token{
		Kind:    kind,
		Content: string(r),
	}
}

// NewNumber returns a new token that holds a number. Since numbers
// are often multiple charaters, the input string should hold the
// number.
func NewNum(num string) Token {
	return Token{
		Kind:    NUM,
		Content: num,
	}
}

// NewWord creates a token that holds the word (also called identifier).
func NewWord(word string) Token {
	return Token{
		Kind:    WORD,
		Content: word,
	}
}

// ======================================================================
// Token Methods and Printing
// ======================================================================

// IsOperator returns ture if the token holds a  +, -, *, / symbol
func (t Token) IsOperator() bool {
	return (t.Kind & (ADD | MUL | SUB | DIV)) != 0
}

func (t Token) String() string {
	if t.Content == "" {
		return fmt.Sprintf("<%s/>", t.Kind)
	}
	return fmt.Sprintf("<%s> %s </%s>", t.Kind, t.Content, t.Kind)
}

func (kind Kind) String() string {
	s, ok := kindToString[kind]
	if !ok {
		s = "?"
	}
	return s
}
