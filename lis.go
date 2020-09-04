package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Token struct {
	valInt          int
	valFloat        float64
	valString       string
	valFunc         func(Token, Token) Token
	childTokens     []Token
	idxCurrentToken int
}

func Eval(expression []interface{}) int {
	// TODO expressionの型をinterface{}にすれば，[]inteface{}も受け取れるのでは

	op, _ := expression[0].(string)

	var a int
	if reflect.TypeOf(expression[1]).Kind() == reflect.Int {
		a, _ = expression[1].(int)
	} else {
		a = Eval(expression[1].([]interface{}))
	}

	var b int
	if reflect.TypeOf(expression[2]).Kind() == reflect.Int {
		b, _ = expression[2].(int)
	} else {
		b = Eval(expression[2].([]interface{}))
	}
	var res int

	if op == "add" {
		res = a + b
	} else if op == "sub" {
		res = a - b
	} else if op == "mul" {
		res = a * b
	} else if op == "div" {
		res = a / b
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
		return res
	}

	f, err := strconv.ParseFloat(tokenStr, 64)
	if err == nil {
		var res Token
		res.valFloat = f
		return res
	}

	var res Token
	res.valString = tokenStr
	return res
}

func main() {
	fmt.Println("hello world!")
}
