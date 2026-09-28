package main

import (
	"fmt"
	"io"
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

	buf := make([]byte, 8)

	for {

		n, err := file.Read(buf)

		if err == io.EOF {
			break
		}

		if n > 0 {
			fmt.Println("read", n, "bytes", string(buf[:n]))
		}

		if err != nil {
			log.Fatal("faild to read file", err)
		}
	}

	// scanner := bufio.NewScanner(file)

	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// create dictionary
	// words := make(map[string]uint)

	// read by word(space) save dic and count it up

	// make dic to array and pick top 3 word
}
