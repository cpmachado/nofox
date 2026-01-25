package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	tape := make([]byte, 3e4)
	ptr := 0
	var stack []int
	var program []rune

	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	// read program
	for _, c := range string(buf) {
		switch c {
		case '>':
			program = append(program, c)
		case '<':
			program = append(program, c)
		case '+':
			program = append(program, c)
		case '-':
			program = append(program, c)
		case '.':
			program = append(program, c)
		case ',':
			program = append(program, c)
		case '[':
			program = append(program, c)
		case ']':
			program = append(program, c)
		}
	}

	for ip := 0; ip < len(program); ip++ {
		switch c := program[ip]; c {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			tape[ptr]++
		case '-':
			tape[ptr]--
		case '.':
			fmt.Printf("%c", tape[ptr])
		case ',':
		case '[':
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
		case ']':
			n := len(stack)
			if n == 0 {
				log.Fatalf("Expected [ to precede closure at %d", ip)
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
