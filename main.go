package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	"github.com/fatih/color"
)

func main() {
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
	tokenize(s)
}

type token struct {
	kind    string
	content string
}

var (
	content     = ""
	kind        = ""
	insideToken = false
	dotUsedAlready = false
)

func tokenize(s string) {
	resetTokenizer()
	for _, r := range s {
		switch r {
		case '(', ')', '+', '-', '*', '/', '=':
			clearToken()
			xml("symbol", r)
			continue
			
		case '.':
			if kind == "" {
				tokErr("cannot use \".\" to start a number")
				return
			}
			if  kind != "number" {
				tokErr("cannot use \".\" outside number")
				return
			}
			if dotUsedAlready {
				tokErr("too many \".\"s in a number")
				return
			} 
			dotUsedAlready = true
			content += "."
			continue
		}
		switch {
		case unicode.IsSpace(r):
			// clearToken()
			continue

		case unicode.IsDigit(r):
			if !insideToken {
				startToken("number")
			}
			content += string(r)
			continue

		case unicode.IsLetter(r):
			if !insideToken {
				startToken("identifier")
			}
			content += string(r)
			continue
		}
		tokErr(`Invalid symbol "` + string(r) + `"`)
		return
	}
	clearToken()
}

func tokErr(i ...interface{}) {
	showErr("Error: That's a mean expression you have.")
	errMessage := fmt.Sprint(i...)
	showErr(errMessage)
}

func startToken(name string) {
	insideToken = true
	kind = name
}

func clearToken() {
	if insideToken {
		xml(kind, content)
	}
	resetTokenizer()
}

func resetTokenizer() {
	insideToken = false
	kind = ""
	content = ""
	dotUsedAlready = false
}

func spaces(depth int) {
	s := ""
	for i := 0; i < depth; i++ {
		s += " "
	}
	fmt.Print(s)
}

func xml(name string, i interface{}) {
	r, ok := i.(rune)
	if ok {
		fmt.Printf("<%s> %c </%s>\n", name, r, name)
	} else {
		fmt.Printf("<%s> %s </%s>\n", name, i, name)
	}
}
