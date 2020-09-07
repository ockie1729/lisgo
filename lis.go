package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	TOKEN_INT = iota
	TOKEN_FLOAT
	TOKEN_STRING
	TOKEN_FUNC
	TOKEN_CHILD_TOKENS
)

type Token struct {
	tokenType int

	valInt    int
	valFloat  float64
	valString string
	valFunc   func(Token) Token

	childTokens     []Token
	idxCurrentToken int
}

type Env struct {
    inner map[string]Token
    outer *Env
}

func (env *Env) Init(parms []string, args []Token, outer *Env) {
    env.inner = map[string]Token{}

    for i := 0; i < len(parms); i++ {
        env.inner[parms[i]] = args[i]
        env.outer = outer
    }
}

func (env *Env) Find(var_name string) *Env {
    if _, ok := env.inner[var_name]; ok {
        return env
    } else {
        return env.outer.Find(var_name)
    }
}

var GlobalEnv Env

func Eval(expression Token) Token {
	if expression.tokenType == TOKEN_STRING {
		return GlobalEnv.Find(expression.valString).inner[expression.valString]
	} else if expression.tokenType != TOKEN_CHILD_TOKENS {
		return expression
	} else if expression.childTokens[0].valString == "define" {
		var_name := expression.childTokens[1].valString
		exp := expression.childTokens[2]
		GlobalEnv.inner[var_name] = exp

		// return expression.childTokens[1] // FIXME
		return Token{}
	} else {
        operatorToken := Eval(expression.childTokens[0])

        operands := []Token{}
        for i := 1; i < len(expression.childTokens); i++ {
            operands = append(operands, Eval(expression.childTokens[i]))
        }
        operandsToken := Token{}
        operandsToken.childTokens = operands
        operandsToken.tokenType = TOKEN_CHILD_TOKENS

        return operatorToken.valFunc(operandsToken)
    }
}

func Tokenize(s string) []string {
	s = strings.Replace(s, "(", " ( ", -1)
	s = strings.Replace(s, ")", " ) ", -1)

	split := strings.Split(s, " ")

	var tokenized []string
	// 空文字を除去
	for _, s := range split {
		if s != "" {
			tokenized = append(tokenized, s)
		}
	}
	return tokenized
}

var idxCurrentToken int

func ReadFrom(inputTokens []string) Token {
	idxCurrentToken = 0
	return readFromRec(inputTokens)
}

func readFromRec(inputTokens []string) Token {
	if idxCurrentToken == len(inputTokens) {
		panic("unexpected EOF while reading")

	}

	tokenStr := inputTokens[idxCurrentToken]
	idxCurrentToken += 1

	if tokenStr == "(" {
		var newToken Token
		newToken.childTokens = make([]Token, 0)
		for inputTokens[idxCurrentToken] != ")" {

			newToken.childTokens = append(newToken.childTokens, readFromRec(inputTokens))
		}
		newToken.tokenType = TOKEN_CHILD_TOKENS
		idxCurrentToken += 1 // pop off ")"
		return newToken
	} else if tokenStr == ")" {
		panic("enexpected )")
	} else {
		return Atom(tokenStr)
	}
}

func Atom(tokenStr string) Token {
	a, err := strconv.Atoi(tokenStr)
	if err == nil {
		var res Token
		res.valInt = a
		res.tokenType = TOKEN_INT
		return res
	}

	f, err := strconv.ParseFloat(tokenStr, 64)
	if err == nil {
		var res Token
		res.valFloat = f
		res.tokenType = TOKEN_FLOAT
		return res
	}

	var res Token
	res.valString = tokenStr
	res.tokenType = TOKEN_STRING
	return res
}

func main() {
	GlobalEnv = Env{}
    GlobalEnv.Init([]string{}, []Token{}, nil)

	fmt.Println("hello world!")
}
