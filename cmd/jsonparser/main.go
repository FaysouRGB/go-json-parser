package main

import (
	"encoding/json"
	"fmt"
	"go-json-parser/pkg/jsonparser"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}
	content, _ := os.ReadFile(os.Args[1])
	tokens := jsonparser.Tokenize(string(content))
	result, _ := jsonparser.Parse(tokens)

	// Pretty print the result
	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Failed to generate JSON:", err)
		return
	}
	fmt.Println(string(prettyJSON))
}
