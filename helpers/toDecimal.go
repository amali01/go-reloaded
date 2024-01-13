package helpers

import (
	"fmt"
	"os"
	"strconv"
)

// HexToDecimal converts a hexadecimal number to decimal
func HexToDecimal(words *[]string, i *int) {
	// Check if the index is out of range
	if *i-1 < 0 || *i-1 >= len(*words) {
		fmt.Println("Hex index out of range")
		os.Exit(0)
	}

	// Extract the hexadecimal number
	hexNum := (*words)[*i-1]

	// Remove the current word containing the hexadecimal number
	(*words)[*i] = ""

	// Convert hexadecimal to decimal
	decNum, err := strconv.ParseInt(hexNum, 16, 64)
	if err != nil {
		fmt.Println("Wrong HEX entry", hexNum)
		os.Exit(0)
	}

	// Update the previous word with the decimal representation
	(*words)[*i-1] = strconv.FormatInt(decNum, 10)

	// Clean up the array by removing empty strings
	CleanedArr(words, i)
}

// BinToDec converts a binary number to decimal
func BinToDec(words *[]string, i *int) {
	// Check if the index is out of range
	if *i-1 < 0 || *i-1 >= len(*words) {
		fmt.Println("Bin index out of range")
		os.Exit(0)
	}

	// Extract the binary number
	binNum := (*words)[*i-1]

	// Remove the current word containing the binary number
	(*words)[*i] = ""

	// Convert binary to decimal
	decNum, err := strconv.ParseInt(binNum, 2, 64)
	if err != nil {
		fmt.Println("Wrong BIN entry", binNum)
		os.Exit(0)
	}

	// Update the previous word with the decimal representation
	(*words)[*i-1] = strconv.FormatInt(decNum, 10)

	// Clean up the array by removing empty strings
	CleanedArr(words, i)
}
