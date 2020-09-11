package main

import (
	"testing"
)

func TestMain(m *testing.M) {
	GlobalEnv = Env{}
	GlobalEnv.Init([]Token{}, []Token{}, nil)
	(&GlobalEnv).AddOperators()

	m.Run()
}
