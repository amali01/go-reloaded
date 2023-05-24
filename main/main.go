package main

import (
	"AMJ"
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	output := os.Args[2]

	dat, err := os.ReadFile(input)
	if err != nil {
		fmt.Printf("Error1")
		return
	}

	dat2 := AMJ.Hextodec(string(dat))

	err = os.WriteFile(output, []byte(dat2), 0644)
	if err != nil {
		fmt.Printf("Error2")
		return
	}

	fmt.Print(string(dat2))

}
