package main

import "fmt"

func main()  {
	if 7 % 2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	fmt.Println()
	if 8%4 == 0 {
		fmt.Println("8 is divisiable by 4")
	}
	fmt.Println()
	if 8%4 == 0 || 7%2==0 {
		fmt.Println("Either 8 or 7 are even")
	}
	fmt.Println()
	if num:= 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}