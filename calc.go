package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, r float64
	var op string

	fmt.Print("Input a: ")
	if _, err := fmt.Scanln(&a); err != nil {
		fmt.Println("a must be numeric")
		return
	}

	fmt.Print("Input b: ")
	if _, err := fmt.Scanln(&b); err != nil {
		fmt.Println("b must be numeric")
		return
	}

	fmt.Print("Input op: ")
	if _, err := fmt.Scanln(&op); err != nil {

		return
	}

	switch op {
	case "+":
		r = a + b
	case "-":
		r = a - b
	case "*":
		r = a * b
	case "/":
		r = a / b

	case "^":
		r = math.Pow(a, b)
	default:
		fmt.Println("Operation must be +, -, /, *, ^")
		return
	}
	fmt.Println(a, op, b, "=", r)
}
