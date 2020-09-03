package main

import (
	"fmt"
	"reflect"
	"strings"
)

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

func Tokenize(s string) []Token {
	s = strings.Replace(s, "(", " ( ", -1)
	s = strings.Replace(s, ")", " ) ", -1)

	split := strings.Split(s, " ")
	var tokenized []Token
	// 空文字を除去
	for _, s := range split {
		if s != "" {
			var token Token
			token.valString = s
			tokenized = append(tokenized, token)
		}
	}
	return tokenized
}

type Token struct {
	valInt    int
	valFloat  float64
	valString string
	valFunc   func(Token, Token) Token
}

func main() {
	fmt.Println("hello world!")
}
