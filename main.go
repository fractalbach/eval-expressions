package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/fractalbach/eval-expressions/lexer"
)

var (
	lex = lexer.New()
)

func main() {
	interactiveMode()
}

func interactiveMode() {
	s := bufio.NewScanner(os.Stdin)
	welcome()
	for {
		inputStarter()
		if !s.Scan() {
			break
		}
		eval(s.Text())
	}
	if s.Err() != nil {
		showErr(s.Err())
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
	
	err := lex.Error()

	if err != nil {
		showErr(err)
		return
	}

	// color.Set(color.FgHiMagenta)
	// fmt.Println("_________[  Info  ]__________")
	// fmt.Println(info)
	// color.Unset()

	color.Set(color.FgGreen)
	fmt.Println("_________[ Answer ]__________")
	// fmt.Println(answer)
	lex.Display()
	color.Unset()
}
