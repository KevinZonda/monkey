package lexer

func (l *Lexer) isLetter() bool {
	return 'a' <= l.curChar && l.curChar <= 'z' ||
		'A' <= l.curChar && l.curChar <= 'Z' ||
		l.curChar == '_'
}

func (l *Lexer) isDigit() bool {
	return '0' <= l.curChar && l.curChar <= '9'
}


func (l *Lexer) skipWhitespaces() {
	if l.curChar == ' ' || l.curChar == '\t' ||
		l.curChar == '\n' || l.curChar == '\r' {
		l.readChar()
		l.skipWhitespaces()
	}
}
