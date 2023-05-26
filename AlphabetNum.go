package AMJ

//import the required packages
import (
	"strings"
	"strconv"
	"fmt"
	"os"
)

// calling the main function
func AlphabetNum(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		if strings.Contains(word[i], "(cap,") {
			// initialize the capitalized
			count, err := strconv.Atoi(strings.TrimRight(word[i+1], ")"))
			if err != nil {
				fmt.Printf("Error during conversion")
				os.Exit(0)
			}else if count > i {
				fmt.Printf("Error outof range number for CAPs")
				os.Exit(0)
			}
			word[i+1] = ""
			word[i] = ""
			for j := 1 ; j <= count ; j++{
				word[i-j] = strings.Title(word[i-j])
			}

		}
		if strings.Contains(word[i], "(up,") {
			// initialize the uppercase
			count, err := strconv.Atoi(strings.TrimRight(word[i+1], ")"))
			if err != nil {
				fmt.Printf("Error during conversion")
				os.Exit(0)
			}else if count > i {
				fmt.Printf("Error outof range number for UP")
				os.Exit(0)
			}
			word[i+1] = ""
			word[i] = ""
			for j := 1 ; j <= count ; j++{
				word[i-j] = strings.ToUpper(word[i-j])
			}

		}
		if strings.Contains(word[i], "(low,") {
			// initialize the lowercase
			count, err := strconv.Atoi(strings.TrimRight(word[i+1], ")"))
			if err != nil {
				fmt.Printf("Error during conversion")
				os.Exit(0)
			}else if count > i {
				fmt.Printf("Error outof range number for LOW")
				os.Exit(0)
			}
			word[i+1] = ""
			word[i] = ""
			for j := 1 ; j <= count ; j++{
				word[i-j] = strings.ToLower(word[i-j])
			}
		}
	}
	return strings.Join(word, " ")
}
