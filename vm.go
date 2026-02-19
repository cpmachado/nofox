package nofox

import (
	"fmt"
	"io"
)

type Int interface{ int64 | int | byte }

type VM[T Int] interface {
	Execute(program AST) error
	ValueAt(idx int) (v T, found bool)
	TapeSize() int
	Input() io.Reader
	Output() io.Writer
}

type defaultVM[T Int] struct {
	tape   []T
	ptr    int
	input  io.Reader
	output io.Writer
}

// NewVM creates an instance of [VM]
//   - tapesize must be a positive value
//   - input must be not nil
//   - output must be not nil
//
// If successful it returns the parametized [VM], if not it validates the
// errors per the order of parameters.
// Returning, if failing:
//   - [ErrInvalidTapeSize]
//   - [ErrInvalidInput]
//   - [ErrInvalidOutput]
func NewVM[T Int](tapesize int, input io.Reader, output io.Writer) (VM[T], error) {
	if tapesize <= 0 {
		return nil, ErrInvalidTapeSize
	}
	if input == nil {
		return nil, ErrInvalidInput
	}
	if output == nil {
		return nil, ErrInvalidOutput
	}
	return &defaultVM[T]{
		tape:   make([]T, tapesize),
		input:  input,
		output: output,
	}, nil
}

func (v *defaultVM[T]) Execute(p AST) error {
	for _, ins := range p {
		switch ins.Type() {
		case NodeTypeMove:
			nmove, ok := ins.(*NodeMove)
			if !ok {
				return ErrInvalidNode
			}
			v.ptr += nmove.Value
		case NodeTypeIncrement:
			ninc, ok := ins.(*NodeIncrement)
			if !ok {
				return ErrInvalidNode
			}
			v.tape[v.ptr] += T(ninc.Value)
		case NodeTypeLoop:
			nloop, ok := ins.(*NodeLoop)
			if !ok {
				return ErrInvalidNode
			}
			for v.tape[v.ptr] > 0 {
				err := v.Execute(nloop.Nodes)
				if err != nil {
					return err
				}
			}
		case NodeTypeRead:
			b := make([]byte, 1)
			_, err := v.input.Read(b)
			if err != nil {
				if err != io.EOF {
					return err
				}
				v.tape[v.ptr] = 0
			} else {
				v.tape[v.ptr] = T(b[0])
			}
		case NodeTypePrint:
			_, _ = fmt.Fprintf(v.output, "%c", rune(v.tape[v.ptr]))
		}
	}
	return nil
}

func (v *defaultVM[T]) ValueAt(idx int) (T, bool) {
	if v.ptr < 0 || len(v.tape) >= idx {
		return T(0), false
	}
	return v.tape[idx], true
}

// TapeSize returns the tape size
func (v *defaultVM[T]) TapeSize() int {
	return len(v.tape)
}

func (v *defaultVM[T]) Input() io.Reader {
	return v.input
}

func (v *defaultVM[T]) Output() io.Writer {
	return v.output
}
