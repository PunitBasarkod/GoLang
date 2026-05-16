package main

import (
	"math"
	"fmt"
)

const s string = "Hello"

func main() {
	fmt.Println(s)

	const n = 500000

	const d = 3e20 / n
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}