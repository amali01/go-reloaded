package AMJ

//import the required packages
import (
	"strings"
)

// calling the main function
func Alphabetformat(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		if strings.Contains(word[i], "(cap)") {
			// initialize the capitalized
			word[i-1] = strings.Title(word[i-1])
			word[i] = ""
		}
		if strings.Contains(word[i], "(up)") {
			// initialize the uppercase
			word[i-1] = strings.ToUpper(word[i-1])
			word[i] = ""
		}
		if strings.Contains(word[i], "(low)") {
			// initialize the lowercase
			word[i-1] = strings.ToLower(word[i-1])
			word[i] = ""
		}
	}
	return strings.Join(word, " ")
}
