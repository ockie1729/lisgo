package main

import "strconv"

const (
	TOKEN_INT = iota
	TOKEN_FLOAT
    TOKEN_BOOL
	TOKEN_STRING
	TOKEN_FUNC
	TOKEN_CHILD_TOKENS
)

type Token struct {
	tokenType int

	valInt    int
	valFloat  float64
    valBool   bool
	valString string
	valFunc   func(Token) Token

	childTokens     []Token
	idxCurrentToken int
}

func (t Token) String() string {
    return strconv.Itoa(t.valInt)
}
