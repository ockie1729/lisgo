package main

import (
	"testing"
)

func TestTokenString(t *testing.T) {
	token := Token{valInt: 1, tokenType: TOKEN_INT}
	got := token.String()
	want := "1"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
