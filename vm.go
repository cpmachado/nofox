package nofox

import (
	"fmt"
	"io"
)

type VM struct {
	tape   []byte
	ptr    int
	input  io.Reader
	writer io.Writer
}

func NewVM(tapesize int, input io.Reader, writer io.Writer) *VM {
	return &VM{
		tape:   make([]byte, tapesize),
		input:  input,
		writer: writer,
	}
}

func (v *VM) Execute(p AST) error {
	for _, ins := range p {
		switch k := ins.(type) {
		case *NodeMove:
			v.ptr += k.Value
		case *NodeIncrement:
			v.tape[v.ptr] += byte(k.Value)
		case *NodeLoop:
			for v.tape[v.ptr] > 0 {
				err := v.Execute(k.Nodes)
				if err != nil {
					return err
				}
			}
		case *NodeRead:
			b := make([]byte, 1)
			_, err := v.input.Read(b)
			if err != nil {
				if err != io.EOF {
					return err
				}
				v.tape[v.ptr] = 0
			} else {
				v.tape[v.ptr] = b[0]
			}
		case *NodePrint:
			fmt.Fprintf(v.writer, "%c", v.tape[v.ptr])
		}
	}
	return nil
}
