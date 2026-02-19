package nofox

import "errors"

var (
	ErrInvalidNode     = errors.New("invalid node")
	ErrInvalidTapeSize = errors.New("invalid tape size")
	ErrInvalidInput    = errors.New("invalid input")
	ErrInvalidOutput   = errors.New("invalid output")
)
