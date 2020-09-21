package main

import (
	"strconv"
	"strings"
)

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
		panic("unexpected EOF while reading") // FIXME エラーを返すように

	}

	tokenStr := inputTokens[idxCurrentToken]
	idxCurrentToken += 1

	if tokenStr == "(" {
		var newToken Token
		newToken.childTokens = make([]Token, 0)
		for inputTokens[idxCurrentToken] != ")" {

			newToken.childTokens = append(newToken.childTokens, readFromRec(inputTokens))
		}
		newToken.tokenType = TOKEN_LIST
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
