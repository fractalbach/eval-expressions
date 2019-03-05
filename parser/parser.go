package parser

import (
	"fmt"

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

func Parse(s string) struct {
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
| ( expr )

*/




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

