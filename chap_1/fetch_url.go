package main

import (
	"strings"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"io"
)

const (
	HTTP_PREFIX = "http://"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, HTTP_PREFIX) {
			url = HTTP_PREFIX + url
			fmt.Printf("\nurl is now: %s\n", url)
		}

		resp, err := http.Get(url)
		fmt.Printf("\n\n\nResponse: %s\n\n\n", resp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "error copying body %v", err)
			os.Exit(1)
		}
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
