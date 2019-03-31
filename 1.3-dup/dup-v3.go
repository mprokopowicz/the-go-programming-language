package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, fileName := range os.Args[1:] {
		fileContent, readFileError := ioutil.ReadFile(fileName)

		if readFileError != nil {
			fmt.Fprintf(os.Stderr, "%v\n", readFileError)
			continue
		}

		for _, line := range strings.Split(string(fileContent), "\n") {
			counts[line] += 1
		}
	}

	for line, numberOfOccurrences := range counts {
		if numberOfOccurrences > 1 {
			fmt.Printf("%d\t%s\n", numberOfOccurrences, line)
		}
	}
}
