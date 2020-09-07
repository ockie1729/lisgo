package main

func AddGlobals() {
	GlobalEnv.inner["+"] = Token{valFunc: Add, tokenType: TOKEN_FUNC}
	GlobalEnv.inner["-"] = Token{valFunc: Sub, tokenType: TOKEN_FUNC}
    GlobalEnv.inner["*"] = Token{valFunc: Mul, tokenType: TOKEN_FUNC}
    GlobalEnv.inner["/"] = Token{valFunc: Div, tokenType: TOKEN_FUNC}
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
