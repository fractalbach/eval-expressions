package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/fractalbach/eval-expressions/lexer"
	// "github.com/Knetic/govaluate"
	"github.com/fractalbach/eval-expressions/parser"
)

var (
	lex        = lexer.New()
	simpleFlag = false
)

func init() {
	flag.BoolVar(&simpleFlag, "s", false, "simple mode. not interactive.")
}

func main() {
	flag.Parse()
	if simpleFlag {
		plainMode()
		return
	}
	interactiveMode()
}

func interactiveMode() {
	s := bufio.NewScanner(os.Stdin)
	welcome()
	inputStarter()
	for s.Scan() {
		eval(s.Text())
		inputStarter()
	}
	if s.Err() != nil {
		showErr(s.Err())
	}
}

func plainMode() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		plainEval(s.Text())
	}
	if s.Err() != nil {
		fmt.Fprintln(os.Stdout, s.Err())
	}
}

func welcome() {
	color.Set(color.FgHiMagenta)
	fmt.Println("══════════( Expression Evaluator )══════════")
	color.Unset()
}

func inputStarter() {
	color.Set(color.FgCyan)
	fmt.Print(">>> ")
	color.Unset()
}

func showErr(err interface{}) {
	color.Set(color.FgHiRed)
	fmt.Fprintln(os.Stderr, err)
	color.Unset()
}

const lolz = `
 __________________________________________________
/                                                  \
|                    R   I   P                     |
|                                                  |
|           Here lies the Tree of Parsing          |
|                                                  |
|             It's final words were....            |
|~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|`

func eval(s string) {
	if s == "" {
		return
	}
	lex.Tokenize(s)
	if lex.Error() != nil {
		showErr(lex.Error())
		return
	}

	// expr, _ := govaluate.NewEvaluableExpression(s)
	// result, _ := expr.Evaluate(nil)

	list := lex.Tokens()
	ast, err := parser.Parse(list)
	if err != nil {
		showErr(lolz)
		showErr(err)
		return
	}

	color.Set(color.FgGreen)
	// fmt.Println(result)
	fmt.Println("_________ The Tree of Parsing __________")
	ast.Display()
	color.Unset()
}

func plainEval(s string) {
	lex.Tokenize(s)
	if lex.Error() != nil {
		fmt.Fprintln(os.Stdout, lex.Error())
		return
	}
	lex.Display()
}
