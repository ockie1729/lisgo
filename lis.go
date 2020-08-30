package main

import (
	"fmt"
	"strings"
)

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
