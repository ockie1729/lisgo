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
	valFunc   func(Token, Token) Token

	childTokens     []Token
	idxCurrentToken int
}

func Eval(expression Token) int {
    if expression.tokenType == TOKEN_STRING {
        return GlobalEnv[expression.valString].valInt
    } else if expression.tokenType != TOKEN_CHILD_TOKENS {
        return expression.valInt
    }

    op := expression.childTokens[0].valString

    var res int


    a := Eval(expression.childTokens[1])
    b := Eval(expression.childTokens[2])

	var res int
	if op == "+" {
		res = a + b
	} else if op == "-" {
		res = a - b
	} else if op == "*" {
		res = a * b
	} else if op == "/" {
		res = a / b
	} else {
        panic("unknown op")
    }

    return res
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
	fmt.Println("hello world!")
}
