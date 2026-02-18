package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/cpmachado/nofox"
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

	vm, err := nofox.NewVM[int](3e4, os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	err = vm.Execute(program)
	if err != nil {
		log.Fatal(err)
	}
}

func loadProgram(r io.Reader) (nofox.AST, error) {
	nofoxChannel := make(chan nofox.Token)

	err := nofox.Lex(r, nofoxChannel)
	if err != nil {
		return nil, err
	}

	program, err := nofox.Parse(nofoxChannel)

	return program, err
}
