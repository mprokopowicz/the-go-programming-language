package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()] += 1
	}

	for lineText, occurrences := range counts {
		if occurrences > 1 {
			fmt.Printf("%d\t%s\n", occurrences, lineText)
		}
	}
}
