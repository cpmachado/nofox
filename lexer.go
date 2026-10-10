package nofox

import (
	"fmt"
	"io"
	"maps"
	"strings"
)

// [Token] defines an enum for tokens
type Token int

//go:generate stringer -type=Token
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

// Brainfuck default mapping tokens to [Token]
//
// Reference: https://esolangs.org/wiki/Brainfuck
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

// [Lexer] is a Brainfuck Lexer, that enables the use of other mappings
type Lexer struct {
	Mappings map[rune]Token // mappings used
	Emitter  chan Token     // emitter of symbols
}

func (l *Lexer) String() string {
	return fmt.Sprintf("Lexer{Mappings: %v, Emitter: %v}", l.Mappings, l.Emitter)
}

// Constructor for [Lexer]
func NewLexer(mappings map[rune]Token, emitter chan Token) (*Lexer, error) {
	if mappings == nil {
		mappings = DefaultMapping
	}
	if emitter == nil {
		return nil, ErrMissingEmitter
	}
	if err := validateMappings(mappings); err != nil {
		return nil, err
	}
	l := new(Lexer)
	l.Mappings = maps.Clone(mappings)
	l.Emitter = emitter

	return l, nil
}

func validateMappings(mappings map[rune]Token) error {
	tracker := map[Token]bool{
		TokenMoveRight: false,
		TokenMoveLeft:  false,
		TokenIncrement: false,
		TokenDecrement: false,
		TokenPrint:     false,
		TokenRead:      false,
		TokenLoopStart: false,
		TokenLoopEnd:   false,
	}

	for _, v := range mappings {
		if exists := tracker[v]; exists {
			return ErrDuplicatedToken
		}
		tracker[v] = true
	}

	var missing []string

	for k, v := range tracker {
		if !v {
			missing = append(missing, k.String())
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing tokens: %q", strings.Join(missing, ","))
	}

	return nil
}

// [Lexer] starts lexing and emitting to emitter defined in Constructor
func (l *Lexer) Lex(r io.Reader) error {
	buf, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	go func() {
		// read program
		for _, c := range string(buf) {
			v, found := l.Mappings[c]
			if found {
				l.Emitter <- v
			}
		}

		close(l.Emitter)
	}()

	return nil
}
