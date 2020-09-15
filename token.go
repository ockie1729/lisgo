package main

import "strconv"

const (
	TOKEN_INT = iota
	TOKEN_FLOAT
	TOKEN_BOOL
	TOKEN_STRING
	TOKEN_FUNC
	TOKEN_LIST
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
	switch t.tokenType {
	case TOKEN_INT:
		return strconv.Itoa(t.valInt)
	case TOKEN_FLOAT:
		return strconv.FormatFloat(t.valFloat, 'f', -1, 64)
	case TOKEN_BOOL:
		return strconv.FormatBool(t.valBool)
	case TOKEN_STRING:
		return t.valString
	case TOKEN_FUNC:
		return "func"
	case TOKEN_LIST:
		return "child"
	default:
		panic("unknown token type ")  // FIXME エラーを返す; LIST表示
	}
}
