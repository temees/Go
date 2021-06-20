package main

import (
	"fmt"
	"strings"
)

func main() {
	var n string
	fmt.Println("input number (100-999): ")
	fmt.Scanln(&n)
	h := strings.Split(n, "")
	fmt.Printf("Hundreds: %v\n", h[0])
	fmt.Printf("Dozens: %v\n", h[1])
	fmt.Printf("Units: %v\n", h[2])

}
