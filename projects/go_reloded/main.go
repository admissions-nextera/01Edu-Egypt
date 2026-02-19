package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: go run . <input_file> <output_file>")
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]
	fmt.Printf("Input: %s, Output: %s\n", inputFile, outputFile)

}
