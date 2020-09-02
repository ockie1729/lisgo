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

func TestEvalAdd(t *testing.T) {
	inputExp := []interface{}{"add", 1, 2}
	got := Eval(inputExp)
	want := 3

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestEvalSub(t *testing.T) {
	inputExp := []interface{}{"sub", 1, 3}
	got := Eval(inputExp)
	want := -2

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestEvalMul(t *testing.T) {
	inputExp := []interface{}{"mul", 1, 3}
	got := Eval(inputExp)
	want := 3

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestEvalDiv(t *testing.T) {
	inputExp := []interface{}{"div", 6, 2}
	got := Eval(inputExp)
	want := 3

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestEvalNestedExp(t *testing.T) {
	inputExp := []interface{}{"add", 1, []interface{}{"add", 2, []interface{}{"add", 2, 2}}}
	got := Eval(inputExp)
	want := 7

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
