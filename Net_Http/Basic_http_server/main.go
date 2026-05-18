package main

import (
	"fmt"
	"net/http"
)

//http.ResponseWriter - write the response back to the client

func helloHandler(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodGet {
		http.Error(w, "Only Get is allowed", http.StatusMethodNotAllowed)
		return
	}

	_, _ = w.Write([]byte("Hello from Go net/http server"))

}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("try going to 8080 port")

	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err)

}