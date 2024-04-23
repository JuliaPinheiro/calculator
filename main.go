package main

import (
	"errors"
	"fmt"
)

func add(number1, number2 int) int {
	return number1 + number2
}

func sub(number1, number2 int) int {
	return number1 - number2
}

func mul(number1, number2 int) int {
	return number1 * number2
}

func div(number1, number2 int) (int, error) {
	if number2 == 0 {
		return 0, errors.New("error: Division by zero")
	} else {
		return number1 / number2, nil
	}
}

func main() {
	fmt.Println(add(10, 20))
	fmt.Println(sub(10, 20))
	fmt.Println(mul(10, 20))
	result, err := div(10,2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
