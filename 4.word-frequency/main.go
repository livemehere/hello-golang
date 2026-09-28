package main

import (
	"fmt"
	"log"
	"os"
)

type WordData struct {
	Word  string
	Count uint
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("targetFile must be provided")
	}

	targetFile := args[0]
	fmt.Printf("loading %s ...\n", targetFile)

	// create dictionary
	// words := make(map[string]uint)

	// read file from argv

	// read by word(space) save dic and count it up

	// make dic to array and pick top 3 word
}
