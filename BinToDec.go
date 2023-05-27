package AMJ

import (
	"strconv"
	"strings"
	"fmt"
	"os"
)
func Bintodec(s string) string {
	word := strings.Fields(s)
	for i := 0; i < len(word); i++ {
		if strings.Contains(word[i], "(bin)") {
			// initialize the binary number
			binary_num := word[i-1]
			word[i] = ""

			var decimal_num int64
			var err error

			// use the parseInt() function to convert
			decimal_num, err = strconv.ParseInt(binary_num, 2, 64)
			// in case of any error
			if err != nil {
				fmt.Println("wrong BIN entry")
				os.Exit(0)
			}
			word[i-1] = strconv.FormatInt(decimal_num, 10)

		}
	}
	return strings.Join(word, " ")
}