---
documentclass: scrartcl
pagesize: a5
fontsize: 12pt
<!-- geometry: landscape -->
<!-- classoption: twocolumn -->
indent: yes
geometry: margin=1in, bottom=1in, top=0.5in
---


# Dragon Book

(pdf page, not actual page)
<!-- TODO: write down actual pages. -->

- (72) : expresions, terms, factors
- (74) : problems
- (81) : traversals & semantic actions {}
- (83) : problems


##  Our Grammar

-	Left-associative means the rules on the left side take priority.
	In the example, they are written beneath each other.

- 	Post-order Traversal (Depth-first) is the order you evaluate the tree
	after it's been built.

- 	Semantic Actions can be written in {brackets},
	in YACC that's what they look like.

Example: Postfix Translation using Semantic Actions

~~~
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
~~~

# Parsing

- Top-down  : easier to write by hand
- Bottom-up : better for parser generator (more powerful)

## Recursive-Decent

- Recursive-Decent is a type of top-down parsing.
- Each non-terminal in the grammar has a procedure associated with it.

### Predictive Parsing

- Predictive Parsing is a form of Recursive-Decent.
- Lookahead Symbol determines flow of control.

### Some PseudoCode

Here's an example grammar defintion:

~~~
	stmt ->
		| if ( expr ) stmt ;
		| for ( optexpr ; optexpr ; optexpr ) stmt ;
~~~

And here's what some of the code for the parser might look like.
See (pdf 88) for the code from the book.

~~~
	lookahead = the currently scanned symbol

	match(t):
		if lookahead == t:
			lookahead = nextTerminal()
		else:
			SyntaxError()

	stmt():
		switch(lookahead):
			case "if":
				match("if")
				match("(")
				expr()
				match(")")
				stmt()

			case "for":
				...

			default:
				SyntaxError()
~~~


Extra: It can be helpful to have a function that tells you
what all of those first terminals are (if, for)

### Empty Strings

When a production has an empty string.
It essentially makes it optional.

~~~
	optexpr ->
		| expr
		| empty
~~~


## Translation vs Parsing

- When writing the grammar, you usually think about how it translates.
- When trying to parse it, you can get stuck.

~~~
	expr ->  
	| expr + term   {add}
	| expr - term   {sub}
	| term
~~~

This one sucks, because how do you decide if you are doing addition
or subtraction when you haven't seen that far yet?

~~~
	expr ->
	| exp1
~~~

------------------------------------------------------------

Lexer vs Parser

Lexical Rules
Syntatic Rules



Scanning
- deletes comments, removes whitespace.

## Lexical Analyzer

Lexical Analysis (in terms of computer programming) is the process of
scanning a string of source code as input,
and producing a sequence of tokens as output

In theory, you don't actually need to create Tokens.
Since Grammars are more powerful than regular expressions,
you could define everything in terms of a grammar.
But this would put a lot of work on the Parser, and make it more
confusing to write, understand, and extend.

### Why

Regular expressions are sufficient for describing the structure of
things like identifiers, constants, keywords, numbers, etc.

Grammars are more powerful, and can more easily handle nested
structures like balanced parentheses.

### Tokens

Tokens usually contain a name and a value.

~~~go
	type Token struct {
		Name string
		Value string
	}
~~~

When displayed, they can easily be represented using xml

~~~xml
	<TokenName> Value </TokenName>
~~~


### Patterns



### Lexeme

A Lexeme is a string that matches the pattern you are looking for.

In my opinion, it's a fancy word that sounds impressive.
To me, it looks like the word Legume.  So I imagine it as finding
a magic bean in your program.  "Ah! I found the bean!"

The only difference is... that bean contains a string of characters.


### What's it all mean

~~~c
printf("Hello, %s is my name.", myName)
~~~
