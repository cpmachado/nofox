package nofox

import (
	"errors"
	"fmt"
	"io"
)

type VM struct {
	tape   []byte
	ptr    int
	Input  io.Reader
	Writer io.Writer
}

func NewVM(tapesize int, input io.Reader, writer io.Writer) *VM {
	return &VM{
		tape:   make([]byte, tapesize),
		Input:  input,
		Writer: writer,
	}
}

func (v *VM) Execute(p AST) error {
	for _, ins := range p {
		switch ins.Type() {
		case NodeTypeMove:
			nmove, ok := ins.(*NodeMove)
			if !ok {
				return errors.New("invalid node")
			}
			v.ptr += nmove.Value
		case NodeTypeIncrement:
			ninc, ok := ins.(*NodeIncrement)
			if !ok {
				return errors.New("invalid node")
			}
			v.tape[v.ptr] += byte(ninc.Value)
		case NodeTypeLoop:
			nloop, ok := ins.(*NodeLoop)
			if !ok {
				return errors.New("invalid node")
			}
			for v.tape[v.ptr] > 0 {
				err := v.Execute(nloop.Nodes)
				if err != nil {
					return err
				}
			}
		case NodeTypeRead:
			b := make([]byte, 1)
			_, err := v.Input.Read(b)
			if err != nil {
				if err != io.EOF {
					return err
				}
				v.tape[v.ptr] = 0
			} else {
				v.tape[v.ptr] = b[0]
			}
		case NodeTypePrint:
			fmt.Fprintf(v.Writer, "%c", v.tape[v.ptr])
		}
	}
	return nil
}
