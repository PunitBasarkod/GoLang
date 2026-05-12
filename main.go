package main

import (
	"fmt"
	"github.com/PunitBasarkod/GoLang/prac"
)

func main() {
	fmt.Println("Hello from Main!!")
	result := sum(4,4)
	fmt.Println("Sum : ", result)


	rs := prac.Abc("Hi", "Hello")
	fmt.Println("words are : ", rs)
}
