package main

func Eval(expression Token) Token {
	return evalRec(expression, &GlobalEnv)
}

func evalRec(x Token, env *Env) Token {
	if x.tokenType == TOKEN_STRING {
		return env.Find(x.valString).inner[x.valString]
	} else if x.tokenType != TOKEN_CHILD_TOKENS {
		return x
	} else if x.childTokens[0].valString == "define" {
		var_name := x.childTokens[1].valString
		exp := x.childTokens[2]
		env.inner[var_name] = evalRec(exp, env)

		return Token{} // FIXME
	} else if x.childTokens[0].valString == "if" {
		test := x.childTokens[1]
		conseq := x.childTokens[2]
		alt := x.childTokens[3]

		if evalRec(test, env).valBool {
			return evalRec(conseq, env)
		} else {
			return evalRec(alt, env)
		}
	} else if x.childTokens[0].valString == "lambda" {
		vars := x.childTokens[1]
		exp := x.childTokens[2]

		res := Token{}
		res.valFunc = func(token Token) Token {
			newEnv := &Env{}
			newEnv.Init(vars.childTokens, token.childTokens, env)
			return evalRec(exp, newEnv)
		}
		return res
	} else {
		operatorToken := evalRec(x.childTokens[0], env)

		operands := []Token{}
		for i := 1; i < len(x.childTokens); i++ {
			operands = append(operands, evalRec(x.childTokens[i], env))
		}
		operandsToken := Token{}
		operandsToken.childTokens = operands
		operandsToken.tokenType = TOKEN_CHILD_TOKENS

		return operatorToken.valFunc(operandsToken)
	}
}
