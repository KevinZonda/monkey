package token

type Token struct {
	Kind    TokenKind
	Literal string
}

func NewToken(kind TokenKind, liter string) Token {
	return Token{Kind: kind, Literal: liter}
}

func NewTokenCh(kind TokenKind, liter byte) Token {
	return NewToken(kind, string(liter))
}
