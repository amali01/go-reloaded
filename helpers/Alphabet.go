package helpers

// Import the required packages
import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// AlphabetFormat formats words based on specified alphabet modifications
func AlphabetFormat(words *[]string, i *int) {

	// Check if the word contains "(cap)"
	if strings.Contains((*words)[*i], "(cap)") {
		// Convert the first character to uppercase and the rest to lowercase
		(*words)[*i-1] = strings.ToLower((*words)[*i-1])
		firstChar := strings.Title(string((*words)[*i-1][0]))
		(*words)[*i-1] = firstChar + (*words)[*i-1][1:]

		// Remove the current word
		(*words)[*i] = ""
	}
	if strings.Contains((*words)[*i], "(up)") {
		// Convert the entire word to uppercase
		(*words)[*i-1] = strings.ToUpper((*words)[*i-1])

		// Remove the current word
		(*words)[*i] = ""
	}
	if strings.Contains((*words)[*i], "(low)") {
		// Convert the entire word to lowercase
		(*words)[*i-1] = strings.ToLower((*words)[*i-1])

		// Remove the current word
		(*words)[*i] = ""
	}
	if strings.Contains((*words)[*i], "(cap,") {
		// Convert a specified number of preceding words to capitalized form
		count, err := strconv.Atoi(strings.TrimRight((*words)[*i+1], ")"))
		if err != nil {
			fmt.Println("Error during conversion")
			os.Exit(0)
		} else if count > *i {
			fmt.Println("Error out of range number for CAPs")
			os.Exit(0)
		}
		(*words)[*i+1] = ""
		(*words)[*i] = ""
		for j := 1; j <= count; j++ {
			(*words)[*i-j] = strings.ToLower((*words)[*i-j])
			firstChar := strings.Title(string((*words)[*i-j][0]))
			(*words)[*i-j] = firstChar + (*words)[*i-j][1:]
		}
	}
	if strings.Contains((*words)[*i], "(up,") {
		// Convert a specified number of preceding words to uppercase
		count, err := strconv.Atoi(strings.TrimRight((*words)[*i+1], ")"))
		if err != nil {
			fmt.Println("Error during conversion")
			os.Exit(0)
		} else if count > *i {
			fmt.Println("Error out of range number for UP")
			os.Exit(0)
		}
		(*words)[*i+1] = ""
		(*words)[*i] = ""
		for j := 1; j <= count; j++ {
			(*words)[*i-j] = strings.ToUpper((*words)[*i-j])
		}
	}
	if strings.Contains((*words)[*i], "(low,") {
		// Convert a specified number of preceding words to lowercase
		count, err := strconv.Atoi(strings.TrimRight((*words)[*i+1], ")"))
		if err != nil {
			fmt.Println("Error during conversion")
			os.Exit(0)
		} else if count > *i {
			fmt.Println("Error out of range number for LOW")
			os.Exit(0)
		}
		(*words)[*i+1] = ""
		(*words)[*i] = ""
		for j := 1; j <= count; j++ {
			(*words)[*i-j] = strings.ToLower((*words)[*i-j])
		}
	}

	// Clean up the array by removing empty strings
	CleanedArr(words, i)
}
