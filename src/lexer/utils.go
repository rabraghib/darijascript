package lexer

import "unicode"

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for null character, indicating end of input
	} else {
		l.ch = rune(l.input[l.readPosition])
	}

	if l.ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readNumber() string {
	position := l.position
	for unicode.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == 0 {
			break
		}
		if l.ch == '"' {
			l.readChar()
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readComment() string {
	position := l.position
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readMultiLineComment() string {
	position := l.position
	endPosition := l.position
	for {
		l.readChar()
		if l.ch == 0 {
			endPosition = l.position
			break
		}
		if l.ch == '*' && l.peekChar() == '/' {
			endPosition = l.position
			l.readChar() // consume '/'
			break
		}
	}
	return l.input[position:endPosition]
}
