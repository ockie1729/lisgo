package main

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

func (env *Env) Find(var_name string) *Env {
    if _, ok := env.inner[var_name]; ok {
        return env
    } else {
        return env.outer.Find(var_name)
    }
}

var GlobalEnv Env
