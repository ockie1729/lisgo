package main

import "strconv"

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

func (t Token) String() string {
    return strconv.Itoa(t.valInt)
}
