package main

func Eval(expression Token) (Token, error) {
	token, err := evalRec(expression, &GlobalEnv)

	return token, err
}

func evalRec(x Token, env *Env) (Token, error) {
	if x.tokenType == TOKEN_STRING {
		found_env, err := env.Find(x.valString)
		if err != nil {
			return Token{}, err
		}
		return found_env.inner[x.valString], nil
	} else if x.tokenType != TOKEN_CHILD_TOKENS {
		return x, nil
	} else if x.childTokens[0].valString == "quote" {
		return Token{childTokens: x.childTokens[1].childTokens,
			tokenType: TOKEN_CHILD_TOKENS}, nil
	} else if x.childTokens[0].valString == "define" {
		var_name := x.childTokens[1].valString
		exp := x.childTokens[2]

		token, err := evalRec(exp, env)
		if err != nil {
			return Token{}, err
		}
		env.inner[var_name] = token

		return Token{}, nil // FIXME
	} else if x.childTokens[0].valString == "if" {
		test := x.childTokens[1]
		conseq := x.childTokens[2]
		alt := x.childTokens[3]

		testRes, err := evalRec(test, env)
		if err != nil {
			return Token{}, err
		}
		if testRes.valBool {
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
			t, _ := evalRec(exp, newEnv) // FIXME 現状はerrorを握りつぶしている
			return t
		}
		return res, nil
	} else if x.childTokens[0].valString == "begin" {
		var val Token
		for i := 1; i < len(x.childTokens); i++ {
			var err error
			val, err = evalRec(x.childTokens[i], env)
			if err != nil {
				return Token{}, err
			}
		}
		return val, nil
	} else {
		operatorToken, err := evalRec(x.childTokens[0], env)

		if err != nil {
			return Token{}, err
		}

		operands := []Token{}
		for i := 1; i < len(x.childTokens); i++ {
			op, err := evalRec(x.childTokens[i], env)
			if err != nil {
				return Token{}, err
			}
			operands = append(operands, op)
		}
		operandsToken := Token{}
		operandsToken.childTokens = operands
		operandsToken.tokenType = TOKEN_CHILD_TOKENS

		return operatorToken.valFunc(operandsToken), nil
	}
}
