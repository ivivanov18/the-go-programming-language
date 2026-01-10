package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"io"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "error copying body %v", err)
			os.Exit(1)
		}

		fmt.Printf("%s", readAll(resp, url))
	}
}

func readAll(resp *http.Response, url string) []byte{
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s\n%v\n",url, err)
			os.Exit(1)
		}
		return b
}
