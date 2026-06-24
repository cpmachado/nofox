package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/debug"

	"github.com/cpmachado/nofox"
)

func main() {
	filename := "stdin"
	version := false
	tapeSize := int(3e4)
	flag.StringVar(&filename, "f", filename, "bf script to load")
	flag.IntVar(&tapeSize, "l", tapeSize, "tapesize, must be a positive integer")
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

	if tapeSize <= 0 {
		fmt.Fprintln(os.Stderr, "invalid tapesize")
		flag.Usage()
		os.Exit(1)
	}

	var err error
	source := os.Stdin

	switch filename {
	case "":
		fmt.Fprintln(os.Stderr, "missing file to be loaded")
		flag.Usage()
		os.Exit(1)
	case "stdin":
	default:
		source, err = os.Open(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}

	program, err := loadProgram(source)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	vm, err := nofox.NewVM[int](tapeSize, os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}

	err = vm.Execute(program)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
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
