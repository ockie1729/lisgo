package main

import (
	"fmt"
	"strings"
)


func Eval(expression []interface{}) int {
	op, _ := expression[0].(string)
	a, _ := expression[1].(int)
	b, _ := expression[2].(int)

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
	tokenized := []string{}
	// 空文字を除去
	for _, token := range split {
		if (token != "") {
			tokenized = append(tokenized, token)
		}
	}
	return tokenized
}

func main() {
	fmt.Println("hello world!")
}
