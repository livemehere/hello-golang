package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Input      string `json:"input"`
	Top        int    `json:"top"`
	IgnoreCase bool   `json:"ignoreCase"`
}

const FILE_NAME = "./config.json"

func main() {
	f, err := os.Open(FILE_NAME)
	if err != nil {
		panic(err)
	}

	config := Config{}

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&config); err != nil {
		panic(err)
	}
	f.Close()

	config.Input = "wow"

	wr, err := os.Create(FILE_NAME)
	if err != nil {
		panic(err)
	}
	defer wr.Close()

	encoder := json.NewEncoder(wr)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		panic(err)
	}
}
