package calc

import "fmt"

func Div(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return x / y, nil
}