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
			if strings.HasPrefix(word[i+1], "a") || strings.HasPrefix(word[i+1], "e") || strings.HasPrefix(word[i+1], "i") || strings.HasPrefix(word[i+1], "o") || strings.HasPrefix(word[i+1], "u") || strings.HasPrefix(word[i+1], "h") {
				word[i] = "an"
			}
		}else if word[i] == "A" {
			if strings.HasPrefix(word[i+1], "a") || strings.HasPrefix(word[i+1], "e") || strings.HasPrefix(word[i+1], "i") || strings.HasPrefix(word[i+1], "o") || strings.HasPrefix(word[i+1], "u") || strings.HasPrefix(word[i+1], "h") {
			word[i] = "An"
			}
		}
		
	}
	temp := strings.Join(word, " ")
	word = strings.Fields(temp)
	return strings.Join(word, " ")
}