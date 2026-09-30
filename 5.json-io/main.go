package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Input      string `json:"input"`
	Top        int    `json:"top"`
	IgnoreCase bool   `json:"ignoreCase"`
}

const FILE_NAME = "./config.json"

func main() {
	data, err := os.ReadFile(FILE_NAME)
	if err != nil {
		panic(err)
	}

	config := Config{}
	json.Unmarshal(data, &config)

	config.Input = "hello world!!"

	output, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(FILE_NAME, output, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
	fmt.Printf("type is %T\n", output)
}
