package parser

import (
	"fmt"
	"strconv"

	"github.com/fractalbach/eval-expressions/token"
)

// ======================================================================
// Notes about the Grammar
// ======================================================================

/*

dragon book info:
(pdf page, not actual page)

(72) : expresions, terms, factors
(74) : problems
(81) : traversals & semantic actions {}
(83) : problems


# Postfix Translation using Semantic Actions

Left-associative
Post-order Traversal

stmnt ->
| id = expr     {assign}
| expr          {display}

expr ->
| expr + term   {add}
| expr - term   {sub}
| term

term ->
| term * factor {mul}
| term / factor {div}
| factor

factor ->
| number        {push number}
| ( expr )


# Parsing

Top-down  : easier to write by hand
Bottom-up : better for parser generator (more powerful)

## Recursive-Decent

- Recursive-Decent is a type of top-down parsing.
- Each nonterminal in the grammar has a procedure associated with it.

### Predictive Parsing

- Predictive Parsing is a form of Recursive-Decent.
- Lookahead Symbol determines flow of control.


*/

// ======================================================================
// Kinds of Nodes
// ======================================================================

// NodeKind represents a specific kind of node that can be found in an
// Abstract Syntax Tree. Each one has a different semantic
// meaning. Terminal is a special kind of node: it hints that you
// should use the token contained within that node.
type NodeKind int

//go:generate stringer -type=NodeKind

const (
	_ NodeKind = 1 << iota
	TERMINAL
	START
	EXPR
	TERM
	FACTOR
)

// ======================================================================
// Abstract Syntax Tree
// ======================================================================

// AST is the public Abstract Syntax tree that you get when you parse
// an expression. Used to display the tree and create the answer.
type AST struct {
	Root *Node
}

type Node struct {
	Parent   *Node
	Children []*Node
	Kind     NodeKind
	Tok      token.Token
}

func NewNode(kind NodeKind) *Node {
	return &Node{
		Kind: kind,
	}
}

func NewLeafNode(tok token.Token) *Node {
	return &Node{
		Kind: TERMINAL,
		Tok:  tok,
	}
}

func (n *Node) appendChild(child *Node) {
	n.Children = append(n.Children, child)
}

// ======================================================================
// Productions
// ======================================================================

/*

stmnt ->
| id = expr     {assign}
| expr          {display}

expr ->
| expr + term   {add}
| expr - term   {sub}
| term

term ->
| term * factor {mul}
| term / factor {div}
| factor

factor ->
| number        {push number}
| word          {push word}
| ( expr )

*/

/*

==================================================
Modified Grammar to eliminate Left-Recursion
==================================================

stmnt ->
| id {push $id} stmnt_rest
| exp

stmnt_rest ->
| = exp
| exp_rest

exp ->
| term exp_rest

exp_rest ->
| + term {add} exp_rest
| - term {subtract} exp_rest
| empty

term ->
| factor term_rest

term_rest ->
| * factor {mult} term_rest
| / factor {div} term_rest
| empty

factor ->
| num {push $num}
| id {push $id}
| ( expr )

*/

// ======================================================================
// Tables and Stuff
// ======================================================================

var (
	symbolTable = map[string]float64{}
	answer      = float64(0)
	stack       = []float64{}
)

func toNum(t token.Token) float64 {
	f, err := strconv.ParseFloat(t.Content, 64)
	if err != nil {
		errtxt += "Eval Error:"
		errtxt += fmt.Sprint(err)
	}
	return f
}

func push(f float64) {
	stack = append(stack, f)
}

func pop() float64 {
	if len(stack) < 1 {
		result += "STACK EMPTY\n"
		errtxt += "eval error: stack empty."
		return 0
	}
	x := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return x
}

func add() {
	b := pop()
	a := pop()
	push(a + b)
}

func subtract() {
	b := pop()
	a := pop()
	push(a - b)
}

func multiply() {
	b := pop()
	a := pop()
	push(a * b)
}

func divide() {
	b := pop()
	a := pop()
	push(a / b)
}

func display() {
	answer = pop()
}

func assign(name string) {
	answer = pop()
	symbolTable[name] = answer
	assignDone = true
}

func lookup(name string) float64 {
	v, ok := symbolTable[name]
	if !ok {
		errtxt += fmt.Sprintf(
			"lookup error: \"%s\" not found.\n", name,
		)
		return 0
	}
	return v
}

// ======================================================================
// Interacting with the Parser
// ======================================================================

// The user calls Parse(), then checks errors using Err(), and finally
// extracts the text using Text().  Similar interface to the Scanner
// interface.

var (
	assignDone = false
	errtxt     = ""
	result     = ""
	lookahead  = 0
	list       = []token.Token{}
)

func clear() {
	assignDone = false
	errtxt = ""
	result = ""
	lookahead = 0
	answer = float64(0)
	stack = []float64{}
}

func Parse(tokenList []token.Token) {
	clear()
	list = tokenList
	start()
}

func Err() error {
	if errtxt == "" {
		return nil
	}
	return fmt.Errorf(errtxt)
}

func Text() string {
	return result
}

func Result() string {
	if assignDone {
		return fmt.Sprintf("Set %s to %v",
			savedFirstTerm, answer,
		)
	}
	return fmt.Sprint(answer)
}

// ======================================================================
// Top-Down Recursive Parser
// ======================================================================

var savedFirstTerm = ""

func start() {
	stmnt()
	if !noTokens() && !hasErr() {
		syntaxError("extra tokens at the end!\n")
		for _, v := range list[lookahead:] {
			errtxt += "\t" + fmt.Sprintln(v)
		}
	}
	if len(stack) > 0 {
		errtxt += "eval error: unevaluated items on stack.\n"
		for _, v := range stack {
			errtxt += "\t" + fmt.Sprintln(v)
		}
	}
}

func stmnt() {
	t := getToken()
	if hasErr() {
		return
	}
	switch t.Kind {
	case token.WORD:
		match(token.WORD)
		savedFirstTerm = t.Content
		stmntRest()
	default:
		expr()
		output("display")
		display()
	}
}

func stmntRest() {
	//optional
	if noTokens() {
		output("push ", savedFirstTerm)
		push(lookup(savedFirstTerm))
		output("display")
		display()
		return
	}
	t := getToken()
	if hasErr() {
		return
	}
	switch t.Kind {
	case token.EQ:
		match(token.EQ)
		expr()
		output("assign to ", savedFirstTerm)
		assign(savedFirstTerm)
	case token.ADD, token.SUB:
		output("push ", savedFirstTerm)
		push(lookup(savedFirstTerm))
		exprRest()
		output("display")
		display()
	case token.MUL, token.DIV:
		output("push ", savedFirstTerm)
		push(lookup(savedFirstTerm))
		termRest()
		output("display")
		display()
	}
}

func expr() {
	term()
	exprRest()
}

func exprRest() {
	// optional
	if noTokens() {
		return
	}
	t := getToken()
	if hasErr() {
		return
	}
	switch t.Kind {
	case token.ADD:
		match(token.ADD)
		term()
		output("add")
		add()
		exprRest()
	case token.SUB:
		match(token.SUB)
		term()
		output("subtract")
		subtract()
		exprRest()
	}
}

func term() {
	factor()
	termRest()
}

func termRest() {
	// optional
	if noTokens() {
		return
	}
	t := getToken()
	if hasErr() {
		return
	}
	switch t.Kind {
	case token.MUL:
		match(token.MUL)
		factor()
		output("mult")
		multiply()
		exprRest()
	case token.DIV:
		match(token.DIV)
		factor()
		output("div")
		divide()
		exprRest()
	}
}

func factor() {
	t := getToken()
	if hasErr() {
		return
	}
	switch t.Kind {
	case token.LEFTPAREN:
		match(token.LEFTPAREN)
		expr()
		match(token.RIGHTPAREN)
	case token.NUM:
		output("push ", t.Content)
		push(toNum(t))
		match(token.NUM)
	case token.WORD:
		output("push ", t.Content)
		match(token.WORD)
		push(lookup(t.Content))
	default:
		syntaxError("expected a number, word, or left paren")
	}
}

func output(args ...interface{}) {
	result += fmt.Sprint(args...) + "\n"
}

// ======================================================================
// Tooling Functions
// ======================================================================

func match(expected token.Kind) {
	given := getToken().Kind
	if given&expected == 0 {
		syntaxError("expected:", expected, ", got:", given)
		return
	}
	lookahead++
}

func getToken() token.Token {
	if lookahead >= len(list) {
		syntaxError("unexpected end of input")
		return token.Token{}
	}
	return list[lookahead]
}

func hasErr() bool {
	return errtxt != ""
}

func noTokens() bool {
	return lookahead >= len(list)
}

func syntaxError(args ...interface{}) {
	errtxt += "Syntax Error: "
	errtxt += fmt.Sprint(args...)
}

// ======================================================================
// Semantics
// ======================================================================

func (ast *AST) semantics() {
	ast.Root.recurseSemantics()
}

// postorder traverse, also called depth-first traverse
func (n *Node) recurseSemantics() {
	n.semantics()
	if len(n.Children) == 0 {
		return
	}
	for _, child := range n.Children {
		child.recurseSemantics()
	}
}

func (n *Node) semantics() {
	switch n.Kind {
	case TERMINAL:
		fmt.Print(n.Tok)
	case START:
		fmt.Println("program starts")
	case EXPR:
		fmt.Println("")
	case TERM:
	case FACTOR:
	}
}
