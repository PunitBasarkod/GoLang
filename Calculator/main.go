package main

import (
	"fmt"

	"github.com/PunitBasarkod/Calculator/calc"
)

func main() {
	for {
		var a, b int
		var op string
		fmt.Println("Enter 1st numbers : ")
		fmt.Scan(&a)

		fmt.Println("Enter 2nd numbers : ")
		fmt.Scan(&b)

		fmt.Println("Enter operation (+, -, *, /) or q to quit: ")
		fmt.Scan(&op)

    if op == "q" || op == "exit" {
        fmt.Println("Exiting program. Goodbye!")
          break
    }

		switch op {
		case "+":
			fmt.Println("The add is : ", calc.Add(a, b))
		case "-":
			fmt.Println("The subtract is : ", calc.Sub(a, b))
		case "*":
			fmt.Println("The multiply is : ", calc.Mul(a, b))
		case "/":
			result, err := calc.Div(a, b)
			if err != nil {
				fmt.Println("Error : ", err)
			} else {
				fmt.Println("Result : ", result)
			}
		}

		var choice string
		fmt.Println("Do you want to continue ? (y/n):")
		fmt.Scan(&choice)

		if choice == "n" || choice == "N" {
			fmt.Println("Exiting the pgm")
			break
		}
		fmt.Println()
	}
}