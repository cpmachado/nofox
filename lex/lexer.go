package lex

import "io"

type Token int

const (
	TokenEOF Token = iota
	TokenMoveRight
	TokenMoveLeft
	TokenIncrement
	TokenDecrement
	TokenPrint
	TokenRead
	TokenLoopStart
	TokenLoopEnd
)

var DefaultMapping = map[rune]Token{
	'>': TokenMoveRight,
	'<': TokenMoveLeft,
	'+': TokenIncrement,
	'-': TokenDecrement,
	'.': TokenPrint,
	',': TokenRead,
	'[': TokenLoopStart,
	']': TokenLoopEnd,
}

func Lex(r io.Reader, emitter chan Token) error {
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
