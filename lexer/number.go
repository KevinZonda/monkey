package lexer

func (l *Lexer) readInt() string {
	pos := l.curPos
	for l.isDigit() {
		l.readChar()
	}
	return l.input[pos:l.curPos]
}
