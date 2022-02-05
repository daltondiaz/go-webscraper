package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := "http://go.dev/doc/tutorial/getting-started/"
	getHtmlFromSomeSite(url)
}

func getHtmlFromSomeSite(url string) {

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	data, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
