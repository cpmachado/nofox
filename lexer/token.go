// lexer is a software package defining the lexer used in nofox

//go:generate stringer -type=Type -output token_string.go -linecomment
package scanner

// [Type] is a Token Type
type Type int

const (
	EOF       Type = iota // eof
	MoveLeft              // move_left
	MoveRight             // move_right
	Decrement             // decrement
	Increment             // increment
	Output                // output
	Input                 // input
	JumpStart             // jump_start
	JumpBack              // jump_back
	Ignored               // ignored
)
