package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, fileName := range files {
			file, openError := os.Open(fileName)
			if openError != nil {
				fmt.Fprintf(os.Stderr, "%v\n", openError)
			} else {
				countLines(file, counts)
				file.Close()
			}
		}
	}

	for lineText, occurrences := range counts {
		if occurrences > 1 {
			fmt.Printf("%d\t%s\n", occurrences, lineText)
		}
	}
}

func countLines(input *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		counts[scanner.Text()] += 1
	}
}
