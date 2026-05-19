package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body failed", err)
		return
	}

	bodyText := string(bodyBytes)

	max := 250
	if(len(bodyText) < max) {
		max = len(bodyText)
	}
	//print

	fmt.Println()
}