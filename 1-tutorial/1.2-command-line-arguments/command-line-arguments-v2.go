package main

import (
	"fmt"
	"os"
)

func main() {
	finalResult, separator := "", ""

	for _, argument := range os.Args[1:] {
		finalResult += separator + argument
		separator = " "
	}

	fmt.Println(finalResult)
}
