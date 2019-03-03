# Creating an Eval function

Write a program that can evaluate expressions.  Include the basic
operators +-*/ and allow for nested parentheses.  Include the ability
to save values to variables using the = operator, and use them in
later expressions.

## Approach

1. Identify the grammar.
2. Create a parse tree.
3. Add operator precendence.
4. Add Symbol tables and variables.


## The Grammar

Start with a Context Free Grammar that can accept valid strings for
expressions.  This one has nested parenthesis, infinite length
integers, and the basic operators.

~~~
	Start symbol: S
	S → E
	E → (E) | N | E+e | E-e | E*e | E/e
	n → 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0
	N → n | nN
	e → (E) | N
~~~

I found a tool online that helps with messing around with grammars:
https://web.stanford.edu/class/archive/cs/cs103/cs103.1156/tools/cfg/
It will show you some example strings for the grammar, and you can add
your own test strings.

One of the things to notice in this grammar is `E → E+e`.  This is
included because if you were to have the rule: `E → E+E`, it would
lead to an [Ambiguous grammar](https://en.wikipedia.org/wiki/Ambiguous_grammar).

![Picture showing the multiple parse trees produced by an ambiguous grammar](https://upload.wikimedia.org/wikipedia/commons/0/0b/Leftmostderivations_jaredwf.png)
