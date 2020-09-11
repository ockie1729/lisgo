package main

import "testing"

func TestEvalNestedExp(t *testing.T) {
	tokenized := Tokenize("(* (+ 4 (+ 1 1)) (/ 6 (* 1 3)))")
	inputExp := ReadFrom(tokenized)
	got := Eval(inputExp).valInt
	want := 12

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDefineVar(t *testing.T) {
	Eval(ReadFrom(Tokenize("(define a 2)")))
	got := Eval(ReadFrom(Tokenize("(* a 2)"))).valInt
	want := 4

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDefineFunc(t *testing.T) {
	Eval(ReadFrom(Tokenize("(define double (lambda (x) (* x 2)))")))
	got := Eval(ReadFrom(Tokenize("(double 4)"))).valInt
	want := 8

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestVariable(t *testing.T) {
	Eval(ReadFrom(Tokenize("(define a 2)")))
	Eval(ReadFrom(Tokenize("(define b 3)")))
	got := Eval(ReadFrom(Tokenize("(+ a b)"))).valInt
	want := 5

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestLambdaExpression(t *testing.T) {
	got := Eval(ReadFrom(Tokenize("((lambda (x) (* x 2)) 2)"))).valInt
	want := 4

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestIfStatement(t *testing.T) {
	got := Eval(ReadFrom(Tokenize("(if (> 4 3) (+ 2 3) (- 3 1))"))).valInt
	want := 5

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestIfStatement2(t *testing.T) {
	got := Eval(ReadFrom(Tokenize("(if (> 3 4) (+ 2 3) (- 3 1))"))).valInt
	want := 2

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
