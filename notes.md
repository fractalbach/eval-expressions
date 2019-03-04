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



