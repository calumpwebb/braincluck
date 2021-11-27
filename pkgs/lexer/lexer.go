package lexer

import (
	"braincluck/pkgs/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	l.readChar()
	return l
}

// readChar: give us the next character and advance our position in the input string
func (l *Lexer) readChar() {
	// in ASCII code, 0 represents "NUL" and we've reached the end of the input
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	l.skipNonValidChars()

	var tok token.Token

	switch l.ch {
	case '>':
		tok = token.New(token.MovePtrRight, string(l.ch))
	case '<':
		tok = token.New(token.MovePtrLeft, string(l.ch))
	case '+':
		tok = token.New(token.IncrementCell, string(l.ch))
	case '-':
		tok = token.New(token.DecrementCell, string(l.ch))
	case '.':
		tok = token.New(token.OutputCell, string(l.ch))
	case ',':
		tok = token.New(token.InputCell, string(l.ch))
	case '[':
		tok = token.New(token.WhileStart, string(l.ch))
	case ']':
		tok = token.New(token.WhileEnd, string(l.ch))
	case 0:
		// END OF FILE
		tok = token.New(token.EOF, "")
	}

	l.readChar()
	return tok
}

// skipNonValidChars: skips any non valid tokens and sets the read head to the right position
func (l *Lexer) skipNonValidChars() {
	// skip over until either end of file or a valid token is found!
	for l.ch != 0 && !token.IsValidToken(l.ch) {
		l.readChar()
	}
}
