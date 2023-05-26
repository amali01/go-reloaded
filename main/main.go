package main

import (
	"AMJ"
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	output := os.Args[2]
	args := os.Args[1:]
	if len(args) !=2 { 
		fmt.Printf("Make sure add the correct amount of Args")
		return
	}

	dat, err := os.ReadFile(input)
	if err != nil {
		fmt.Printf("Error1")
		return
	}

	dat2 := AMJ.Hextodec(string(dat))
	dat2 = AMJ.Bintodec(string(dat2))
	dat2 = AMJ.Alphabetformat(string(dat2))
	dat2 = AMJ.AlphabetNum(string(dat2))
	dat2 = AMJ.FormatPunctuation(string(dat2))
	dat2 = AMJ.Atoan(string(dat2))

	err = os.WriteFile(output, []byte(dat2), 0644)
	if err != nil {
		fmt.Printf("Error2")
		return
	}

	fmt.Print(string(dat2))

}
