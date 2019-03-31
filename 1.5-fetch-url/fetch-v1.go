package main

import (
	"fmt"
	"io/ioutil"
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

		body, err := ioutil.ReadAll(response.Body)
		response.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%s", body)
	}
}
