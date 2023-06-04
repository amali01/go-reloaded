package AMJ

//import the required packages
import (
	"strings"
)

// calling the main function
func Atoan(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		if word[i] == "a" {
			if strings.HasPrefix(word[i+1], "a") || strings.HasPrefix(word[i+1], "A") || strings.HasPrefix(word[i+1], "e") || strings.HasPrefix(word[i+1], "E") || strings.HasPrefix(word[i+1], "i") || strings.HasPrefix(word[i+1], "I") || strings.HasPrefix(word[i+1], "o") || strings.HasPrefix(word[i+1], "O") || strings.HasPrefix(word[i+1], "u") || strings.HasPrefix(word[i+1], "U") || strings.HasPrefix(word[i+1], "h") || strings.HasPrefix(word[i+1], "H") {
				word[i] = "an"
			}
		} else if word[i] == "A" {
			if strings.HasPrefix(word[i+1], "a") || strings.HasPrefix(word[i+1], "A") || strings.HasPrefix(word[i+1], "e") || strings.HasPrefix(word[i+1], "E") || strings.HasPrefix(word[i+1], "i") || strings.HasPrefix(word[i+1], "I") || strings.HasPrefix(word[i+1], "o") || strings.HasPrefix(word[i+1], "O") || strings.HasPrefix(word[i+1], "u") || strings.HasPrefix(word[i+1], "U") || strings.HasPrefix(word[i+1], "h") || strings.HasPrefix(word[i+1], "H") {
				word[i] = "An"
			}
		}

	}
	return strings.Join(word, " ")
}
