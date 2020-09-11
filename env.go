package main

import "fmt"

type Env struct {
	inner map[string]Token
	outer *Env
}

func (env *Env) Init(parms []Token, args []Token, outer *Env) {
	env.inner = map[string]Token{}

	for i := 0; i < len(parms); i++ {
		env.inner[parms[i].valString] = args[i]
		env.outer = outer
	}
}

func (env *Env) Find(var_name string) (*Env, error) {
	if _, ok := env.inner[var_name]; ok {
		return env, nil
	} else if env.outer == nil {
		return nil, fmt.Errorf("%q not defined", var_name)
	} else {
		found_env, err := env.outer.Find(var_name)
		if err != nil {
			return nil, err
		} else {
			return found_env, nil
		}
	}
}

var GlobalEnv Env
