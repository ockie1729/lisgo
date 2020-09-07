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

func (env *Env) Init(parms []Token, args []Token, outer *Env) {
    env.inner = map[string]Token{}

    for i := 0; i < len(parms); i++ {
        env.inner[parms[i].valString] = args[i]
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

func Eval(expression Token) Token{
    return evalRec(expression, &GlobalEnv)
}

func evalRec(x Token, env *Env) Token {
	if x.tokenType == TOKEN_STRING {
		return env.Find(x.valString).inner[x.valString]
	} else if x.tokenType != TOKEN_CHILD_TOKENS {
		return x
	} else if x.childTokens[0].valString == "define" {
		var_name := x.childTokens[1].valString
		exp := x.childTokens[2]
		env.inner[var_name] = exp

		return Token{}  // FIXME
	} else if x.childTokens[0].valString == "lambda" {
        vars := x.childTokens[1]
        exp := x.childTokens[2]

        res := Token{}
        res.valFunc = func(token Token) Token {
            newEnv := &Env{}
            newEnv.Init(vars.childTokens, token.childTokens, env)
            return evalRec(exp, newEnv)}
        return res
    } else {
        operatorToken := evalRec(x.childTokens[0], env)

        operands := []Token{}
        for i := 1; i < len(x.childTokens); i++ {
            operands = append(operands, evalRec(x.childTokens[i], env))
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
    GlobalEnv.Init([]Token{}, []Token{}, nil)

	fmt.Println("hello world!")
}
