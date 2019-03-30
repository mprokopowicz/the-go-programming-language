package main

import (
	"fmt"
	"os"
)

func main() {
	var finalResult, separator string

	for i := 1; i < len(os.Args); i++ {
		finalResult += separator + os.Args[i]
		separator = " "
	}

	fmt.Println(finalResult)
}
