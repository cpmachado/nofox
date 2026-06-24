package nofox

import (
	"io"
	"os"
	"testing"
)

func TestNewVM(t *testing.T) {
	tests := []struct {
		description string
		tapesize    int
		input       io.Reader
		output      io.Writer
		err         error
	}{
		{description: "normal", tapesize: 3e5, input: os.Stdin, output: os.Stdout},
		{description: "no tape", tapesize: 0, input: os.Stdin, output: os.Stdout, err: ErrInvalidTapeSize},
		{description: "negative tape", tapesize: -1, input: os.Stdin, output: os.Stdout, err: ErrInvalidTapeSize},
		{description: "input nil", tapesize: 3e5, input: nil, output: os.Stdout, err: ErrInvalidInput},
		{description: "output nil", tapesize: 3e5, input: os.Stdin, output: nil, err: ErrInvalidOutput},
		{description: "io nil", tapesize: 3e5, input: nil, output: nil, err: ErrInvalidInput},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			// given
			tapesize := test.tapesize
			input := test.input
			output := test.output
			err := test.err

			// when
			vm, gerr := NewVM[byte](tapesize, input, output)

			// then
			if err != gerr {
				t.Errorf("Expected error to be %v, but got %v\n", err, gerr)
			}
			if vm != nil {
				if got := vm.TapeSize(); tapesize != got {
					t.Errorf("Expected tape size %d, but got %d\n", tapesize, got)
				}
				if got := vm.Input(); input != got {
					t.Errorf("Expected input to be %v, but got %v\n", input, got)
				}
				if got := vm.Output(); output != got {
					t.Errorf("Expected output to be %v, but got %v\n", output, got)
				}
			}
		})
	}
}

func Test_DefaultVM_TapeSize(t *testing.T) {
	tests := []struct {
		description string
		tapesize    int
		input       io.Reader
		output      io.Writer
	}{
		{description: "1", tapesize: 1, input: os.Stdin, output: os.Stdout},
		{description: "3", tapesize: 3, input: os.Stdin, output: os.Stdout},
		{description: "30", tapesize: 30, input: os.Stdin, output: os.Stdout},
		{description: "300", tapesize: 300, input: os.Stdin, output: os.Stdout},
		{description: "3k", tapesize: 3e4, input: os.Stdin, output: os.Stdout},
		{description: "30k", tapesize: 3e5, input: os.Stdin, output: os.Stdout},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			// given
			tapesize := test.tapesize
			input := test.input
			output := test.output

			// when
			vm, err := NewVM[byte](tapesize, input, output)
			// then
			if err != nil {
				t.Errorf("Expected error to be nil, but got %v\n", err)
				return
			}

			if got := vm.TapeSize(); tapesize != got {
				t.Errorf("Expected tape size %d, but got %d\n", tapesize, got)
			}
		})
	}
}
