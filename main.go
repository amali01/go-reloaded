package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	output := os.Args[2]

	dat, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(output, []byte(dat), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(dat))

}
