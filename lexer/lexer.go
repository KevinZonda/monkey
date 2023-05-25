package lexer

import "github.com/KevinZonda/monkey/token"

type Lexer struct {
	input     string
	curPos    int // Current pos
	toReadPos int // next Char we need to read
	curChar   byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.toReadPos >= len(l.input) {
		l.curChar = 0
	} else {
		l.curChar = l.input[l.toReadPos]
	}
	l.curPos = l.toReadPos
	l.toReadPos++
}

func (l *Lexer) NextToken() token.Token {
	var tk token.Token
	l.skipWhitespaces()
	switch l.curChar {
	case '=':
		tk = token.NewTokenCh(token.ASSIGN, l.curChar)
	case ';':
		tk = token.NewTokenCh(token.SEMICOLON, l.curChar)

	case '(':
		tk = token.NewTokenCh(token.L_PAREN, l.curChar)
	case ')':
		tk = token.NewTokenCh(token.R_PAREN, l.curChar)

	case ',':
		tk = token.NewTokenCh(token.COMMA, l.curChar)
	case '+':
		tk = token.NewTokenCh(token.PLUS, l.curChar)

	case '{':
		tk = token.NewTokenCh(token.L_BRACE, l.curChar)
	case '}':
		tk = token.NewTokenCh(token.R_BRACE, l.curChar)

	case 0:
		tk = token.NewToken(token.EOF, "")
	default:
		if l.isLetter() {
			tk.Literal = l.readIdentifier()
			tk.Kind = mapIdentifier(tk.Literal)
			return tk
		} else if l.isDigit() {
			tk.Literal = l.readInt()
			tk.Kind = token.INT
			return tk
		} else {
			tk = token.NewTokenCh(token.BAD, l.curChar)
		}
	}
	l.readChar()
	return tk
}
