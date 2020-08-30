package main

import "testing"

import "github.com/google/go-cmp/cmp"

func TestTokenize(t *testing.T) {
	got := Tokenize("(+ 1 2)")
	want := []string{"(", "+", "1", "2", ")"}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("got %q want %q", got, want)
	}
}
