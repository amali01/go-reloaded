package AMJ

//import the required packages
import (
	"strconv"
	"strings"
	"fmt"
)

// calling the main function
func Hextodec(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		if strings.Contains(word[i], "(hex)") {
			// initialize the hexadecimal number
			hexadecimal_num := word[i-1]
			word[i] = ""

			var decimal_num int64
			var err error

			// use the parseInt() function to convert
			decimal_num, err = strconv.ParseInt(hexadecimal_num, 16, 64)
			// in case of any error
			if err != nil {
				fmt.Printf("Error4")
				return("Error")
			}
			word[i-1] = strconv.FormatInt(decimal_num, 10)

		}
	}
	temp := strings.Join(word, " ")
	word = strings.Fields(temp)
	return strings.Join(word, " ")
}
