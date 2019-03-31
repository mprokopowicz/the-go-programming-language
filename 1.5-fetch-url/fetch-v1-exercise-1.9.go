package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if strings.HasPrefix(url, "http://") == false {
			url = "http://" + url
		}

		response, httpError := http.Get(url)
		if httpError != nil {
			fmt.Fprintf(os.Stderr, "%v\n", httpError)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "Response status HTTP %d - %s\n\n", response.StatusCode, response.Status)

		_, err := io.Copy(os.Stdout, response.Body)

		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
		}
		response.Body.Close()
	}
}
