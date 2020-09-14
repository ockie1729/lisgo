package main

func (env *Env) AddOperators() {
	env.inner["+"] = Token{valFunc: Add, tokenType: TOKEN_FUNC}
	env.inner["-"] = Token{valFunc: Sub, tokenType: TOKEN_FUNC}
	env.inner["*"] = Token{valFunc: Mul, tokenType: TOKEN_FUNC}
	env.inner["/"] = Token{valFunc: Div, tokenType: TOKEN_FUNC}
	env.inner[">"] = Token{valFunc: Greater, tokenType: TOKEN_FUNC}
	env.inner["="] = Token{valFunc: Equal, tokenType: TOKEN_FUNC}
	env.inner["car"] = Token{valFunc: Car, tokenType: TOKEN_FUNC}
	env.inner["cdr"] = Token{valFunc: Cdr, tokenType: TOKEN_FUNC}
}

func Add(operandsToken Token) Token {
	a := operandsToken.childTokens[0].valInt
	b := operandsToken.childTokens[1].valInt

	ans := Token{}
	ans.valInt = a + b
	ans.tokenType = TOKEN_INT
	return ans
}

func Sub(operandsToken Token) Token {
	a := operandsToken.childTokens[0].valInt
	b := operandsToken.childTokens[1].valInt

	ans := Token{}
	ans.valInt = a - b
	ans.tokenType = TOKEN_INT
	return ans
}

func Mul(operandsToken Token) Token {
	a := operandsToken.childTokens[0].valInt
	b := operandsToken.childTokens[1].valInt

	ans := Token{}
	ans.valInt = a * b
	ans.tokenType = TOKEN_INT
	return ans
}

func Div(operandsToken Token) Token {
	a := operandsToken.childTokens[0].valInt
	b := operandsToken.childTokens[1].valInt

	ans := Token{}
	ans.valInt = a / b
	ans.tokenType = TOKEN_INT
	return ans
}

func Greater(operandsToken Token) Token {
	a := operandsToken.childTokens[0].valInt
	b := operandsToken.childTokens[1].valInt

	ans := Token{}
	ans.valBool = a > b
	ans.tokenType = TOKEN_BOOL
	return ans
}

func Equal(operandsToken Token) Token {
	// TODO 数値以外も比較できるようにする
	a := operandsToken.childTokens[0].valInt
	b := operandsToken.childTokens[1].valInt

	return Token{valBool: a == b, tokenType: TOKEN_BOOL}
}

func Car(operandsToken Token) Token {
	return operandsToken.childTokens[0].childTokens[0]
}

func Cdr(operandsToken Token) Token {
	return Token{childTokens: operandsToken.childTokens[0].childTokens[1:],
		tokenType: TOKEN_CHILD_TOKENS}
}
