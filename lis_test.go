package main

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	got := Tokenize("(+ 1 2)")
	want := []string{"(", "+", "1", "2", ")"}

	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("got %q want %q", got[i], want[i])
		}
	}
}

func TestReadFrom(t *testing.T) {
	got := ReadFrom(Tokenize("(- 1 2)"))

	if got.childTokens[0].valString != "-" {
		t.Errorf("got %q want -", got.childTokens[0].valString)
	}

	if got.childTokens[1].valInt != 1 {
		t.Errorf("got %q want 1", got.childTokens[1].valInt)
	}

	if got.childTokens[2].valInt != 2 {
		t.Errorf("got %q want 2", got.childTokens[2].valInt)
	}
}

func TestReadFromNestedExp(t *testing.T) {
	got := ReadFrom(Tokenize("(- 1 (+ 2 3))"))

	if got.childTokens[0].valString != "-" {
		t.Errorf("got %q want -", got.childTokens[0].valString)
	}
	if got.childTokens[1].valInt != 1 {
		t.Errorf("got %q want 1", got.childTokens[1].valInt)
	}
	if got.childTokens[2].childTokens[0].valString != "+" {
		t.Errorf("got %q want +", got.childTokens[2].childTokens[0].valString)
	}
	if got.childTokens[2].childTokens[1].valInt != 2 {
		t.Errorf("got %q want 2", got.childTokens[2].childTokens[1].valString)
	}
	if got.childTokens[2].childTokens[2].valInt != 3 {
		t.Errorf("got %q want 3", got.childTokens[2].childTokens[2].valString)
	}
}

func TestAtomInt(t *testing.T) {
	input := "8"

	got := Atom(input).valInt
	want := 8

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAtomFloat(t *testing.T) {
	input := "3.14"

	got := Atom(input).valFloat
	var want float64
	want = 3.14

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAtomString(t *testing.T) {
	input := "add"

	got := Atom(input).valString
	want := "add"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestEvalAdd(t *testing.T) {
	inputExp := ReadFrom(Tokenize("(+ 1 2)"))
	got := Eval(inputExp)
	want := 3

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestEvalAdd2(t *testing.T) {
	inputExp := ReadFrom(Tokenize("(+ 2 3)"))
	got := Eval(inputExp)
	want := 5

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}


// func TestEvalSub(t *testing.T) {
// 	inputExp := []interface{}{"sub", 1, 3}
// 	got := Eval(inputExp)
// 	want := -2

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// func TestEvalMul(t *testing.T) {
// 	inputExp := []interface{}{"mul", 1, 3}
// 	got := Eval(inputExp)
// 	want := 3

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// func TestEvalDiv(t *testing.T) {
// 	inputExp := []interface{}{"div", 6, 2}
// 	got := Eval(inputExp)
// 	want := 3

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// func TestEvalNestedExp(t *testing.T) {
// 	inputExp := []interface{}{"add", 1, []interface{}{"add", 2, []interface{}{"add", 2, 2}}}
// 	got := Eval(inputExp)
// 	want := 7

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }
