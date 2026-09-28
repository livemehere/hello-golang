package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type WordData struct {
	Word  string
	Count int
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
		log.Fatal("failed to open file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	counts := make(map[string]int)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		word = strings.Trim(word, ",.?!{}()[]")
		if word == "" {
			continue
		}
		counts[word]++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("failed to scan Scanner")
	}

	var words []WordData

	for word, count := range counts {
		words = append(words, WordData{
			Word:  word,
			Count: count,
		})
	}

	sort.Slice(words, func(i, j int) bool {
		if words[i].Count == words[j].Count {
			return words[i].Word < words[j].Word
		}

		return words[i].Count > words[j].Count
	})

	for _, item := range words[:min(3, len(words))] {
		fmt.Println(item.Word, item.Count)
	}
}
