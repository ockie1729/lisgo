package main

import (
	"fmt"
)

func main() {
	GlobalEnv = Env{}
    GlobalEnv.Init([]Token{}, []Token{}, nil)
  
	fmt.Println("hello world!")
}
