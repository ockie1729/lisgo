package main

import (
	"testing"
)

func TestTokenString(t *testing.T) {
	tokens := []Token{
		Token{valInt: 1, tokenType: TOKEN_INT},
		Token{valFloat: 1.2, tokenType: TOKEN_FLOAT},
		Token{valBool: false, tokenType: TOKEN_BOOL},
		Token{valString: "Go", tokenType: TOKEN_STRING},
		Token{valFunc: nil, tokenType: TOKEN_FUNC},
	}
	wants := []string{"1", "1.2", "false", "Go", "func"}

	for i := 0; i < len(tokens); i++ {
		if tokens[i].String() != wants[i] {
			t.Errorf("got %q want %q", tokens[i].String(), wants[i])
		}
	}
}
