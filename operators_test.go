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

func TestCar(t *testing.T) {
	res, err := Eval(ReadFrom(Tokenize("(car (quote (-1 2)))")))

	got := res.valInt
	want := -1

	if err != nil {
		t.Errorf("error: %v", err)
	} else if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCd(t *testing.T) {
	res, err := Eval(ReadFrom(Tokenize("(cdr (quote (-1 2 3)))")))
	got := res.childTokens

	resWant, _ := Eval(ReadFrom(Tokenize("(quote (2 3))")))
	want := resWant.childTokens

	if err != nil {
		t.Errorf("error: %v", err)
	} else {
		for i := 0; i < len(want); i++ {
			if got[i].valInt != want[i].valInt {
				t.Errorf("got %v want %v", got[i], want[i])
			}
		}

	}
}

func TestNullQuestion(t *testing.T) {
	res, err := Eval(ReadFrom(Tokenize("(null? (quote ()))")))
	got := res.valBool
	want := true

	if err != nil {
		t.Errorf("error: %v", err)
	} else if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCons(t *testing.T) {
	res, err := Eval(ReadFrom(Tokenize("(cons 3 (quote (2 1)))")))
	got := res.childTokens

	resWant, _ := Eval(ReadFrom(Tokenize("(quote (3 2 1))")))
	want := resWant.childTokens

	if err != nil {
		t.Errorf("error: %v", err)
	} else {
		for i := 0; i < len(want); i++ {
			if got[i].valInt != want[i].valInt {
				t.Errorf("got %v want %v", got[i], want[i])
			}
		}

	}
}
