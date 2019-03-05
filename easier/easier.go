package easier

import (
	"unicode"
	"fmt"
)

type node struct {
	label    string
	children []*node
}

type evaluator struct {
}



// Evaluate the expression given by the string. Returns 2 strings.
// The 1st string has the solution, and the 2nd string has extra
// information.  The error will be non-nill and contain an error
// message if there were any syntax errors or stuff like that.
func Eval(s string) (string, string, error) {

	var ans string
	var info string
	var err error

	var num string   // accumulates numbers into a string.
	var depth int // depth caused by parens
	var word string // keeps track of the current word

	runes := []rune(s)
	limit := len(runes)

	// Iterate through the characters (called runes in Go), and do
	// some preliminary operations on them.  This will make it
	// easier to deal with later.

	for i := 0; i < limit; i++ {

		// Remove all whitespace, we don't want them at all in
		// our expression solver.

		if unicode.IsSpace(runes[i]) {
			continue
		}

		// If you start with a number, then assume that you
		// have a number, and keep iterating through it until
		// you hit either (1: the end) or (2: an operator) or
		// (3: close paren)
		//
		// Cases for syntax errors: (1. a letter followed by a
		// number) and (2: a paren followed by a number)

		if !unicode.IsNumber(runes[i]) {
			goto afterNumber
		}
		for unicode.IsNumber(runes[i]) {
			num += string(runes[i])
			i++
			if i >= limit {
				ans += term(depth, "num", num) + "\n"
				goto ret
			}
		}
		ans += term(depth, "num", num) + "\n"
		
	afterNumber:

		// --------------------
		// Check for Operators
		// --------------------
		if isOperator(runes[i]) {
			ans += term(depth, "op", string(runes[i])) + "\n"
			continue
		}

		// --------------------
		// Check for Parens
		// --------------------
		if runes[i] == '(' {
			ans += opentag(depth, "group")
			ans += "\n"
			depth++
			continue
		}
		if runes[i] == ')' {
			depth--
			ans += closetag(depth, "group")
			ans += "\n"
			continue
		}

		// --------------------
		// Words
		// --------------------
		// The final thing to do is check for words, which
		// will just be strings of non-number and non-symbol
		// characters put together.
		for unicode.IsLetter(runes[i]) {
			word += string(runes[i])
			i++
			if i >= limit {
				ans += term(depth, "word", word)
				ans += "\n"
				goto ret
			}
		}
		if word != "" {
			ans += term(depth, "word", word)
			ans += "\n"
			continue
		}
		
		ans += term(depth, "unid", string(runes[i])) + "\n"
		
	}
	goto ret
ret:
	return ans, info, err
}

// isOperator simply returns true if the rune is one of the main
// operators that we want in our evaluation program.  The equal sign
// (=) is NOT going to count as an operator, we are going to treat
// that as it's own special thing called assignment.
func isOperator(r rune) bool {
	switch r {
	case '-', '*', '/', '+':
		return true
	}
	return false
}


// spaces
func spaces(depth int) string {
	s := ""
	for i := 0; i<depth; i++ {
		s += "  "
	}
	return s
}

// term returns a string in the form of <name>contents</name>
func term(depth int, name, contents string) string {
	return fmt.Sprintf(
		"%s<%s> %s </%s>",
		spaces(depth), name, contents, name)
}

// opentag
func opentag(depth int, name string) string {
	return fmt.Sprintf(
		"%s<%s>",
		spaces(depth), name)
}

// close tag
func closetag(depth int, name string) string {
	return fmt.Sprintf(
		"%s</%s>",
		spaces(depth), name)
}
