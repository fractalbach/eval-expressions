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


# Context Free Grammar

Context Free Grammars are awesome.  They let you do a lot of neat
things, and they are great for describing programming languages.
I think designing one gives you a better understanding of how to evaluate
expressions.

I found a tool online that helps with messing around with grammars:
https://web.stanford.edu/class/archive/cs/cs103/cs103.1156/tools/cf
It's helpful for experimenting with different grammars quickly. It
shows a few example input strings that the grammar will accept.  You
can also add your own test strings, and see the derivations.


## The Expression

A good starting place is to focus on the overall structure of an
expression.  We want parentheses and operators.

~~~
Start symbol: S
S → E
E → n | (E) | E+e | E-e | E*e | E/e
e → n | (E)
~~~

One of the things to notice in this grammar is `E → E+e`.  This is
included because if you were to have the rule: `E → E+E`, it would
lead to an [Ambiguous grammar](https://en.wikipedia.org/wiki/Ambiguous_grammar).


## Numbers

Grammar for integers and floating point numbers.  Infinite length is allowed.

~~~
Start symbol: Q
Z → N | NZ
Q → Z | Z.Z
N → 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
~~~


## Identifiers

This grammar shows identifiers.  In many programming languages,
identifiers need to begin with a letter, and then can have numbers
anywhere in the name.  Since accessing properties of objects often use
a dot notation, we want those numbers to be integers instead of
floating points.

~~~
Start symbol: W
W → L | LW | LZ | LZW
Z → N | NZ
Q → Z | Z.Z
N → 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
L → a | b | c | d | e | f | g | h | i | j | k | l | m | n | o | p | q | r | s | t | u | v | w | x | y | z
~~~

This example only uses the 26 lowercase letters because I don't want
to list them all out.  When actually programming this, you can define
"letter" or "character" to mean whatever you want.



## Variable Assignments

Once you have identifiers and expressions worked out, it's relatively
straight forward to add variable assignment on to everything else. An
assignment consists of an identifer to save the value into, an equal
sign (=), and an expression.

~~~
Start symbol: S
S -> E | V
V -> W=E
E -> ...
~~~


## Putting it Together

After putting together all of the parts I think will be needed to
evaluate expressions, I came up with this grammar:

~~~
Start symbol: S
S → E | V
E → T | (E) | E+X | E-X | E/X | E*X
V → W=E
Z → N | NZ
N → 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
F → Z | Z.Z
L → a | b | c | d | e | f | g | h | i | j | k | l | m | n | o | p | q | r | s | t | u | v | w | x | y | z
W → L | LW | LZ | LZW
T → W | F
X → T | (E)
~~~


It doesn't allow any whitespace at all, but should work fine if you just ignore/delete all whitespace.
