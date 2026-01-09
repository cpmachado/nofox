package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cpmachado/nofox"
)

func main() {
	filename := ""
	flag.StringVar(&filename, "f", "", "Input file")
	flag.Parse()

	if filename == "" {
		flag.CommandLine.Usage()
		log.Fatal("missing input file")
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	l := nofox.Lexer{TokenMap: nofox.DefaultTokens, Source: file}

	tokens := l.Lex()

	for _, token := range tokens {
		if token.Type != nofox.Ignored {
			fmt.Println(token.String())
		}
	}
}
