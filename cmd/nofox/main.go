package main

import (
	"flag"
	"log"
	"os"
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

}
