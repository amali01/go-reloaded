package helpers

//import the required packages
import (
	"strings"
)

// isVowel checks if a string starts with a vowel (a, e, i, o, u, or their uppercase counterparts)
func isVowel(s string) bool {
	return strings.HasPrefix(s, "a") ||
		strings.HasPrefix(s, "A") ||
		strings.HasPrefix(s, "e") ||
		strings.HasPrefix(s, "E") ||
		strings.HasPrefix(s, "i") ||
		strings.HasPrefix(s, "I") ||
		strings.HasPrefix(s, "o") ||
		strings.HasPrefix(s, "O") ||
		strings.HasPrefix(s, "u") ||
		strings.HasPrefix(s, "U") ||
		strings.HasPrefix(s, "h") ||
		strings.HasPrefix(s, "H")
}

// Atoan converts "a" or "A" to "an" or "An" respectively if the following word starts with a vowel
func Atoan(words *[]string, i *int) {
	// Check if there is a word following the current index
	if *i+1 < len(*words) {
		word := (*words)[*i]
		nextWord := (*words)[*i+1]

		// Check if the current word is "a" and the next word starts with a vowel
		if (word == "a") && isVowel(nextWord) {
			(*words)[*i] = "an"
		} else if (word == "A") && isVowel(nextWord) {
			(*words)[*i] = "An"
		}
	}
}
