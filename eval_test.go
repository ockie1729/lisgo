package main

import "testing"

func TestEvalNestedExp(t *testing.T) {
	tokenized := Tokenize("(* (+ 4 (+ 1 1)) (/ 6 (* 1 3)))")
	inputExp := ReadFrom(tokenized)
	res, _ := Eval(inputExp)
	got := res.valInt
	want := 12

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDefineVar(t *testing.T) {
	Eval(ReadFrom(Tokenize("(define a 2)")))
	res, _ := Eval(ReadFrom(Tokenize("(* a 2)")))
	got := res.valInt
	want := 4

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestUndefinedVar(t *testing.T) {
	_, err := Eval(ReadFrom(Tokenize("(+ hoge 2)")))

	if err == nil {
		t.Errorf("got %q want err", err)
	}
}

func TestDefineFunc(t *testing.T) {
	Eval(ReadFrom(Tokenize("(define double (lambda (x) (* x 2)))")))
	res, _ := Eval(ReadFrom(Tokenize("(double 4)")))
	got := res.valInt
	want := 8

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestVariable(t *testing.T) {
	Eval(ReadFrom(Tokenize("(define a 2)")))
	Eval(ReadFrom(Tokenize("(define b 3)")))
	res, _ := Eval(ReadFrom(Tokenize("(+ a b)")))
	got := res.valInt
	want := 5

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestLambdaExpression(t *testing.T) {
	res, _ := Eval(ReadFrom(Tokenize("((lambda (x) (* x 2)) 2)")))
	got := res.valInt
	want := 4

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestIfStatement(t *testing.T) {
	res, _ := Eval(ReadFrom(Tokenize("(if (> 4 3) (+ 2 3) (- 3 1))")))
	got := res.valInt
	want := 5

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestBeginStatement(t *testing.T) {
	res, err := Eval(ReadFrom(Tokenize("(begin (define n 2) (= n 2))")))
	got := res.valBool
	want := true

	if err != nil {
		t.Errorf("error: %v", err)
	} else if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
