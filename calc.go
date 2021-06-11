package main

import "fmt"

func main() {
	var a float64
	if _, err := fmt.Scanln(&a); err != nil {
		return
	}

}
