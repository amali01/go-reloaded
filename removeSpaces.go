package AMJ

import (
	"fmt"
	"strings"
)

// removeSpaces removes any spaces between two ' marks
func removeSpaces(s string) string {
	// split the string by ' marks
	parts := strings.Split(s, "'")
	// loop over the parts and trim any spaces
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	// join the parts back with ' marks
	return strings.Join(parts, "'")
}

func main() {
	// test cases
	test1 := "I am exactly how they describe me: ' awesome '"
	test2 := "As Elton John said: ' I am the most well-known homosexual in the world '"
	test3 := "She said: ' I don't like you ' and left"

	fmt.Println(removeSpaces(test1)) // I am exactly how they describe me: 'awesome'
	fmt.Println(removeSpaces(test2)) // As Elton John said: 'I am the most well-known homosexual in the world'
	fmt.Println(removeSpaces(test3)) // She said: 'I don't like you' and left
}
