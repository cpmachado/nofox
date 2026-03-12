package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"

	"github.com/cpmachado/nofox"
)

func main() {
	filename := ""
	version := false
	flag.StringVar(&filename, "f", "", "bf script to load")
	flag.BoolVar(&version, "v", version, "display version")
	flag.Parse()

	if version {
		vsn := "dev"
		if info, ok := debug.ReadBuildInfo(); ok {
			vsn = info.Main.Version
		}
		fmt.Printf("nofox-%s\n", vsn)
		os.Exit(0)
	}

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

	vm := nofox.NewVM(3e4, os.Stdin, os.Stdout)

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
