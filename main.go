package main

import (
	"bufio"
	"fmt"
	"go-reloaded/helpers"
	"log"
	"os"
	"strings"
)

func main() {
	// Check if the correct number of arguments is provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		return
	}

	// Extract the input and output file names from command-line arguments
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Open the input file
	input, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Failed to open input file: %v", err)
	}
	defer input.Close()

	// Open the output file
	output, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer output.Close()

	// Create a buffered writer for the output file
	writer := bufio.NewWriter(output)

	// Read the content from the input file
	fileContent, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Make sure you have a correct input")
		os.Exit(0)
	} else if len(fileContent) == 0 {
		fmt.Println("Your sample file is empty")
		os.Exit(0)
	}

	// Split the content into lines
	lines := strings.Split(string(fileContent), "\n")

	// Process each line
	for index, line := range lines {
		// Split the line into words
		words := strings.Fields(line)

		// Process each word in the line
		for i := 0; i < len(words); i++ {
			switch {
			case strings.Contains(words[i], "(hex)"):
				helpers.HexToDecimal(&words, &i)
			case strings.Contains(words[i], "(bin)"):
				helpers.BinToDec(&words, &i)
			case (words[i] == "a") || (words[i] == "A"):
				helpers.Atoan(&words, &i)
			case strings.Contains(words[i], "(cap") ||
				strings.Contains(words[i], "(up") ||
				strings.Contains(words[i], "(low"):
				helpers.AlphabetFormat(&words, &i)
			}
		}

		// Join the modified words back into a line
		modifiedLine := strings.Join(words, " ")

		// Apply punctuation formatting to the modified line
		modifiedLine = helpers.PunctuationFormat(modifiedLine)

		// Write the modified line to the output file
		if index != len(lines)-1 {
			_, err = writer.WriteString(modifiedLine + "\n")
			if err != nil {
				log.Fatalf("Failed to write to output file: %v", err)
			}
		} else {
			_, err = writer.WriteString(modifiedLine)
			if err != nil {
				log.Fatalf("Failed to write to output file: %v", err)
			}
		}
	}

	// Flush any buffered data to the output file
	err = writer.Flush()
	if err != nil {
		log.Fatalf("Failed to flush output file: %v", err)
	}

	fmt.Printf("Modified text saved to %s\n", outputFile)
}
