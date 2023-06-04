package main

import (
	"AMJ"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Make sure add the correct amount of Args")
		os.Exit(0)
	}
	input := os.Args[1]
	output := os.Args[2]

	dat, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Make sure u have a correct input")
		os.Exit(0)
	} else if len(dat) == 0 {
		fmt.Println("You sample File is Empty")
		os.Exit(0)
	}

	dat2 := AMJ.Hextodec(string(dat))
	dat2 = AMJ.Bintodec(string(dat2))
	dat2 = AMJ.Atoan(string(dat2))
	dat2 = AMJ.Alphabetformat(string(dat2))
	dat2 = AMJ.AlphabetNum(string(dat2))
	dat2 = AMJ.FormatPunctuation(string(dat2))

	err = os.WriteFile(output, []byte(dat2), 0644)
	if err != nil {
		fmt.Println("somthing wrong with the output")
		os.Exit(0)
	}

	fmt.Print(string(dat2))

}
