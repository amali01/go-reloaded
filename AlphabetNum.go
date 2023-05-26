package AMJ

//import the required packages
import (
	"strings"
	"strconv"
)

// calling the main function
func AlphabetNum(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		if strings.Contains(word[i], "(cap,") {
			// initialize the capitalized
			count, err := strconv.Atoi(strings.TrimRight(word[i+1], ")"))
			if err != nil {
				return("Error during conversion")
			}else if count > i-1 {
				return("Error outof range number for CAPs")
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
				return("Error during conversion")
			}else if count > i-1 {
				return("Error outof range number for UP")
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
				return("Error during conversion")
			}else if count > i-1 {
				return("Error outof range number for LOW")
			}
			word[i+1] = ""
			word[i] = ""
			for j := 1 ; j <= count ; j++{
				word[i-j] = strings.ToLower(word[i-j])
			}
		}
	}
	temp := strings.Join(word, " ")
	word = strings.Fields(temp)
	return strings.Join(word, " ")
}
