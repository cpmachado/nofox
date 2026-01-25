package main

import (
	"fmt"
	"os"

	"github.com/cpmachado/nofox"
)

func main() {
	l := nofox.Lexer{TokenMap: nofox.DefaultTokens, Source: os.Stdin}

	tokens := l.Lex()

	for _, token := range tokens {
		if token.Type != nofox.Ignored {
			fmt.Println(token.String())
		}
	}
}
