package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {

	url := "http://www.google.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close();
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: Reading %s: %v\n", url, err)
	}
	fmt.Printf("%s", b)
}
