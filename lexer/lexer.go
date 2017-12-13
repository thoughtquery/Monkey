// take source code as input and output the tokens that represent the source code
// go through input and output the next token it recognizes

package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (after current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Give us the next character and advance our input in the input string
func (l *Lexer) readChar() {
	// Are we at the end of the input?
	if l.readPosition >= len(l.input) {
		// ASCII code for 'NUL' byte (Haven't read anything yet/End Of File)
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	// Always point from position were going to read from next
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACKET, l.ch)
	case '}':
		tok = newToken(token.RBRACKET, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.reachChar()
	return tok
}

func newToken(tokenType token.tokenType, ch byte) token.tokenType {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
