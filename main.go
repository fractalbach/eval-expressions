package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/fractalbach/eval-expressions/lexer"
)

var (
	lex             = lexer.New()
	interactiveFlag = false
)

func init () {
	flag.BoolVar(&interactiveFlag, "i", false, "use interactive mode.")
}

func main() {
	flag.Parse()
	if interactiveFlag {
		interactiveMode()
		return
	}
	plainMode()
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

func eval(s string) {
	lex.Tokenize(s)
	err := lex.Error()
	if err != nil {
		showErr(err)
	} else {
		color.Set(color.FgGreen)
		lex.Display()
		color.Unset()
	}
}

func plainEval(s string) {
	lex.Tokenize(s)
	if lex.Error() != nil {
		fmt.Fprintln(os.Stdout, lex.Error())
		return
	}
	lex.Display()
}
