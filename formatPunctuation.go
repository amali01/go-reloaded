package AMJ

import (
	"regexp"
)

// formatPunctuation formats the punctuation marks in a string
func FormatPunctuation(s string) string {
	// use regular expressions to match and replace the punctuation marks
	re1 := regexp.MustCompile(`\s*([.,!?:;])`)           // match any spaces around single punctuation marks
	re2 := regexp.MustCompile(`'\s*(.*?)\s*'`)           // match any spaces between ' marks
	re3 := regexp.MustCompile(`([.,!?:;])([[:alpha:]])`) // match any spaces around single punctuation marks

	s = re1.ReplaceAllString(s, "$1")    // replace the spaces with one space after the punctuation mark
	s = re2.ReplaceAllString(s, "'$1'")  // replace the spaces with no spaces around the ' marks
	s = re3.ReplaceAllString(s, "$1 $2") // replace the spaces with one space after the punctuation mark
	return s
}
