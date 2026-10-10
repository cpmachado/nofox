package nofox

import "errors"

var (
	ErrInvalidNode     = errors.New("invalid node")
	ErrInvalidTapeSize = errors.New("invalid tape size")
	ErrInvalidInput    = errors.New("invalid input")
	ErrInvalidOutput   = errors.New("invalid output")

	ErrDuplicatedToken = errors.New("duplicated token in mappings") // duplicated token
	ErrMissingEmitter  = errors.New("missing emitter")              // missing emitter, e.g. constructor of [Lexer]
	ErrLoopEnd         = errors.New("received loop end")
	ErrMissingLoopEnd  = errors.New("missing loop end")
)
