package parser

import (
	"fmt"

	"github.com/fractalbach/eval-expressions/token"
)

// Parse takes a list of tokens and converts them into an Abstract
// Syntax Tree.
func Parse(list []token.Token) (AST, error) {
	if len(list) == 0 {
		return AST{}, nil
	}
	p := &parser{
		root:   newNode("start"),
		done:   make(chan bool),
		tokens: list,
		index:  0,
	}
	p.node = p.root
	go p.parseStart()
	<-p.done
	ast := AST{root: p.root}
	var err error
	if p.err != "" {
		err = fmt.Errorf("%s", p.err)
	}
	return ast, err
}

// ======================================================================
// The Parser and the Syntax Tree
// ======================================================================

type parser struct {
	node   *node
	root   *node
	tokens []token.Token
	index  int
	err    string
	done   chan bool
}

// ------------------------------------------------------------
//  Iterating through Tokens
// ------------------------------------------------------------

func (p *parser) token() token.Token {
	fmt.Println("current:", p.index, "  total:", len(p.tokens))
	return p.tokens[p.index]
}

func (p *parser) hasMoreTokens() bool {
	return p.index+1 < len(p.tokens)
}

func (p *parser) advance() {
	if !p.hasMoreTokens() {
		p.done <- true
		return
	}
	p.index++
	fmt.Println("advanced to:", p.index)
}

func (p *parser) peek() (token.Token, bool) {
	if p.hasMoreTokens() {
		return p.tokens[p.index+1], true
	}
	return token.Token{}, false
}

func (p *parser) consume() {
	p.node.pushChild(newNode(p.token().Content))
	fmt.Println("consumed:", p.token())
	p.advance()
}

// ------------------------------------------------------------
// Specific Parsing Functions
// ------------------------------------------------------------

func (p *parser) parseStart() {
	p.parseExpression()
	p.done <- true
}

func (p *parser) parseExpression() {
	p.pushAndFollow("expression")
	defer p.climbUp()
	//
	// Start by checking for a nested expression.
	//
	if p.token().Content == "(" {
		p.consume()
		if p.token().Content != ")" {
			p.parseExpression()	
		}
		if p.token().Content != ")" {
			p.err = "open paren must have matching close paren."
			return
		}
		p.consume()
		return
	}
	//
	// Since it doesn't start with parens, it might be a terminal
	// number or an identifier
	//
	kind := p.token().Kind
	if kind == token.NumberToken || kind == token.IDToken {
		p.consume()
		return
	}
	//
	// Since it's not a terminal, it be an expression with an
	// operator, followed by a secondary expression.
	//
	p.parseExpression()
	p.parseOperator()
	p.parseSecondaryExpression()
}

func (p *parser) parseOperator() {
	p.pushAndFollow("operator")
	defer p.climbUp()
	//
	// Don't be fooled by other symbols! they're just imposters!
	//
	switch p.token().Content {
	case "-", "+", "/", "*":
		p.consume()
		return
	}
	//
	// Reject those other foolish symbols and whatever other
	// craziness may have emerged from the depths of the parser!
	//
	p.throw(
		"expected operator, but you gave me this crap instead!:",
		p.token().Content,
	)
}

func (p *parser) parseSecondaryExpression() {
	p.pushAndFollow("2ndExpression")
	defer p.climbUp()
	//
	// If you aren't in a nest, you must be terminally falling
	// toward the end.
	//
	if p.token().Content != "(" {
		p.parseTerm()
		return
	}
	//
	// aww yeah looks like we got some parens! (((Nest it up)))!
	//
	p.consume()
	p.parseExpression()
	if p.token().Content != ")" {
		p.throw("=( open paren must have matching close paren...")
		return
	}
	p.consume() // )
}

func (p *parser) parseTerm() {
	p.pushAndFollow("term")
	defer p.climbUp()
	kind := p.token().Kind
	if !(kind == token.NumberToken || kind == token.IDToken) {
		p.throw(
			"I wanted a Number or Identifer",
			"But instead you gave me lies!\n",
			"wtf is this?!: ",
			p.token().Content,
		)
	}
}

func (p *parser) throw(v ...interface{}) {
	p.err = fmt.Sprint(v...)
	p.done <- true
}

// ======================================================================
// Abstract Syntax Tree Nodes
// ======================================================================

type node struct {
	name     string
	terminal bool
	parent   *node
	children []*node
}

func (n *node) pushChild(child *node) {
	n.children = append(n.children, child)
	child.parent = n
}

func (p *parser) pushAndFollow(name string) {
	child := newNode(name)
	p.node.pushChild(child)
	p.node = child
}

func (p *parser) climbUp() {
	p.node = p.node.parent
}

func newNode(name string) *node {
	return &node{name: name}
}

func (n *node) display(depth int) {
	spaces := ""
	for i := 0; i < depth; i++ {
		spaces += " "
	}
	fmt.Printf("%v%v\n", spaces, n.name)
	for _, next := range n.children {
		next.display(depth + 1)
	}
}

// ======================================================================
// Publicly Accesible Abstract Syntax Tree
// ======================================================================

// AST is an abstract syntax that has been produced by the parser. You
// can use this to display the tree.
type AST struct {
	root *node
}

func (ast AST) Display() {
	ast.root.display(0)
}
