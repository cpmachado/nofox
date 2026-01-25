package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cpmachado/nofox/lex"
)

func main() {
	filename := ""
	flag.StringVar(&filename, "f", "", "bf script to load")
	flag.Parse()

	if filename == "" {
		flag.Usage()
		log.Fatalf("missing file to be loaded")
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	program, err := loadProgram(file)
	if err != nil {
		log.Fatal(err)
	}

	tape := make([]byte, 3e4)
	ptr := 0
	var stack []int
	b := make([]byte, 1)

	for ip := 0; ip < len(program); ip++ {
		switch c := program[ip]; c {
		case lex.TokenMoveRight:
			ptr++
		case lex.TokenMoveLeft:
			ptr--
		case lex.TokenIncrement:
			tape[ptr]++
		case lex.TokenDecrement:
			tape[ptr]--
		case lex.TokenPrint:
			fmt.Printf("%c", tape[ptr])
		case lex.TokenRead:
			_, err := os.Stdin.Read(b)
			if err != nil {
				if err != io.EOF {
					log.Fatal(err)
				}
				tape[ptr] = 0
			} else {
				tape[ptr] = b[0]
			}
		case lex.TokenLoopStart:
			if tape[ptr] == 0 {
				k := 0
				for program[ip] != ']' && k != 0 {
					switch program[ip] {
					case ']':
						k--
					case '[':
						k++
					}
					ip++
				}
			} else {
				stack = append(stack, ip)
			}
		case lex.TokenLoopEnd:
			n := len(stack)
			if n == 0 {
				log.Fatalf("expected '[' to precede closure at %d", ip)
			}
			k := stack[n-1]
			if tape[ptr] > 0 {
				ip = k
			} else {
				stack = stack[:n-1]
			}
		}
	}
}

func loadProgram(r io.Reader) ([]lex.Token, error) {
	var program []lex.Token

	lexChannel := make(chan lex.Token)

	err := lex.Lex(r, lexChannel)
	if err != nil {
		return nil, err
	}

	for c := <-lexChannel; c != lex.TokenEOF; c = <-lexChannel {
		program = append(program, c)
	}

	return program, nil
}
