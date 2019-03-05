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
	fmt.Println(title("Expression Evaluator!", 70))
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

	color.Set(color.FgHiMagenta)
	// fmt.Println("══════════ Tokens ══════════ ")
	lex.Display()
	color.Unset()

	if err != nil {
		showErr(err)
		return
	}

	fmt.Println()

	// color.Set(color.FgGreen)
	// fmt.Println("══════════ Answer ══════════ ")
	// color.Unset()
}

func title(title string, boxSize int) string {
	top := "┌"
	middle := "│"
	bottom := "╘"
	for i := 0; i < boxSize; i++ {
		top += "─"
		bottom += "═"
	}
	i := 0
	nSpaces := boxSize/2 - len(title)/2
	for ; i < nSpaces; i++ {
		middle += " "
	}
	i += len(title)
	middle += title
	for ; i < boxSize; i++ {
		middle += " "
	}
	top += "┐\n"
	middle += "│\n"
	bottom += "╛"
	return top + middle + bottom
}
