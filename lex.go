package nofox

import (
	"fmt"
	"io"
	"text/scanner"
)

type TokenType byte

//go:generate stringer -type=TokenType -output lex_string.go -linecomment
const (
	EOF       TokenType = iota // eof
	MoveLeft                   // move_left
	MoveRight                  // move_right
	Decrement                  // decrement
	Increment                  // increment
	Output                     // output
	Input                      // input
	JumpStart                  // jump_start
	JumpBack                   // jump_back
	Ignored                    // ignored
)

type Token struct {
	Type TokenType
	Text rune
}

func (t *Token) String() string {
	return fmt.Sprintf("%s,%s", t.Type, string(t.Text))
}

type TokenMap map[rune]TokenType

var DefaultTokens = TokenMap{
	'<': MoveLeft,
	'>': MoveRight,
	'-': Decrement,
	'+': Increment,
	'.': Output,
	',': Input,
	'[': JumpStart,
	']': JumpBack,
}

type Lexer struct {
	TokenMap TokenMap
	Source   io.Reader
}

func (l *Lexer) Lex() []Token {
	var s scanner.Scanner
	var tokens []Token

	s.Init(l.Source)

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		v, found := l.TokenMap[tok]
		if !found {
			v = Ignored
		}
		tokens = append(tokens, Token{Type: v, Text: tok})
	}

	return tokens
}
