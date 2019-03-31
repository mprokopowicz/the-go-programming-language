package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		response, httpError := http.Get(url)
		if httpError != nil {
			fmt.Fprintf(os.Stderr, "%v\n", httpError)
			os.Exit(1)
		}

		_, err := io.Copy(os.Stdout, response.Body)

		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
		}
		response.Body.Close()
	}
}
