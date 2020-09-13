package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	REPL_PROMPT_STR = "lis.go> "
)

func Repl() {
	stdin := bufio.NewScanner(os.Stdin)

	fmt.Print(REPL_PROMPT_STR)
	for stdin.Scan() {
		exp := stdin.Text()
		val, err := Eval(ReadFrom(Tokenize(exp)))

		if err != nil {
			fmt.Println(err)
			fmt.Print(REPL_PROMPT_STR)
			continue
		}

		fmt.Println(val.String())

		fmt.Print(REPL_PROMPT_STR)
	}
}

func main() {
	GlobalEnv = Env{}
	GlobalEnv.Init([]Token{}, []Token{}, nil)
	(&GlobalEnv).AddOperators()

	Repl()
}
