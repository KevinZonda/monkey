package lexer

import "github.com/KevinZonda/monkey/token"

func (l *Lexer) readIdentifier() string {
	pos := l.curPos
	for l.isLetter() {
		l.readChar()
	}
	return l.input[pos:l.curPos]
}

var reservedKeywords = map[string]token.TokenKind{
	"fn":  token.FUNC,
	"let": token.LET,
}

func mapIdentifier(ident string) token.TokenKind {
	if tk, ok := reservedKeywords[ident]; ok {
		return tk
	}
	return token.IDENTITY
}
