package lex

import "io"

type TokenType int

const (
	TokenEOF TokenType = iota
	TokenMoveRight
	TokenMoveLeft
	TokenIncrement
	TokenDecrement
	TokenPrint
	TokenRead
	TokenLoopStart
	TokenLoopEnd
)

var DefaultMapping = map[rune]TokenType{
	'>': TokenMoveRight,
	'<': TokenMoveLeft,
	'+': TokenIncrement,
	'-': TokenDecrement,
	'.': TokenPrint,
	',': TokenRead,
	'[': TokenLoopStart,
	']': TokenLoopEnd,
}

func Lex(r io.Reader, emitter chan TokenType) error {
	buf, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	go func() {
		// read program
		for _, c := range string(buf) {
			v, found := DefaultMapping[c]
			if found {
				emitter <- v
			}
		}

		close(emitter)
	}()

	return nil
}
