// ======================================================================
// Token Structure
// ======================================================================

package token

import (
	"fmt"
)

const (
	NoKind      = ""       // default token kind
	NumberToken = "number" // token holding a number.
	IDToken     = "id"     // token holding an identifier.
	SymbolToken = "symbol" // token holding a symbol.
)

// Token holds a string and some useful information about that string.
type Token struct {
	Kind    string // the kind of info the token holds.
	Content string // the actual string inside the token.
}

// New creates and returns a new Token.
func New(kind, content string) Token {
	return Token{kind, content}
}

func (t Token) String() string {
	return fmt.Sprintf("<%s> %s </%s>", t.Kind, t.Content, t.Kind)
}
