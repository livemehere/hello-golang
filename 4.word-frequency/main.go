package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type WordData struct {
	Word  string
	Count uint
}

func main() {
	// read file from argv
	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("first argument must be provided (read file)")
	}

	targetFile := args[0]
	fmt.Printf("loading %s ...\n", targetFile)

	// read file
	file, err := os.Open(targetFile)
	if err != nil {
		log.Fatal("faild to open file", err)
	}
	defer file.Close()

	fmt.Println(file)

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	counts := make(map[string]int)

	for scanner.Scan() {
		word := scanner.Text()
		counts[word]++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for word, count := range counts {
		fmt.Println(word, count)
	}
}
