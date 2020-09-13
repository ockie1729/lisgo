package main

import "testing"

func TestEvalAdd(t *testing.T) {
	inputExp := ReadFrom(Tokenize("(+ 1 2)"))
	res, _ := Eval(inputExp)
	got := res.valInt
	want := 3

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestEvalSub(t *testing.T) {
	inputExp := ReadFrom(Tokenize("(- 1 2)"))
	res, _ := Eval(inputExp)
	got := res.valInt
	want := -1

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestEvalMul(t *testing.T) {
	inputExp := ReadFrom(Tokenize("(* 2 3)"))
	res, _ := Eval(inputExp)
	got := res.valInt
	want := 6

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestEvalDiv(t *testing.T) {
	inputExp := ReadFrom(Tokenize("(/ 4 2)"))
	res, _ := Eval(inputExp)
	got := res.valInt
	want := 2

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestEvalGreater(t *testing.T) {
	inputExp := ReadFrom(Tokenize("(> 4 2)"))
	res, _ := Eval(inputExp)
	got := res.valBool
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestEvalEqualInt(t *testing.T) {
	inputExp := ReadFrom(Tokenize("(= 1 1)"))
	res, _ := Eval(inputExp)
	got := res.valBool
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
