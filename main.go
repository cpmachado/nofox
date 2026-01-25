package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
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
			_, err := os.Stdin.Read(b)
			if err != nil {
				if err != io.EOF {
					log.Fatal(err)
				}
				tape[ptr] = 0
			} else {
				tape[ptr] = b[0]
			}
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

func loadProgram(r io.Reader) ([]rune, error) {
	var program []rune

	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, err
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

	return program, nil
}
