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
	env.inner["cons"] = Token{valFunc: Cons, tokenType: TOKEN_FUNC}
	env.inner["null?"] = Token{valFunc: NullQuestion, tokenType: TOKEN_FUNC}
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
	if len(operandsToken.childTokens[0].childTokens) == 0 {
		panic("can't car '()") // FIXME エラーを返す
	}

	return operandsToken.childTokens[0].childTokens[0]
}

func Cdr(operandsToken Token) Token {
	switch len(operandsToken.childTokens[0].childTokens) {
	case 0:
		panic("can't cdr '()") // FIXME エラーを返す
	case 1:
		return Token{childTokens: []Token{}, tokenType: TOKEN_LIST}
	default:
		return Token{childTokens: operandsToken.childTokens[0].childTokens[1:],
			tokenType: TOKEN_LIST}
	}
}

func NullQuestion(operandsToken Token) Token {
	if operandsToken.childTokens[0].tokenType == TOKEN_LIST &&
		len(operandsToken.childTokens[0].childTokens) == 0 {
		return Token{valBool: true, tokenType: TOKEN_BOOL}
	} else {
		return Token{valBool: false, tokenType: TOKEN_BOOL}
	}
}

func Cons(operandsToken Token) Token {
	car := operandsToken.childTokens[0]
	cdr := operandsToken.childTokens[1].childTokens

	cdr = append([]Token{car}, cdr...) // TODO 効率的な処理に修正

	return Token{childTokens: cdr, tokenType: TOKEN_LIST}
}
