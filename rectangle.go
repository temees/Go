package main

import (
	"fmt"
)

func main() {
	var w float64
	var h float64
	fmt.Println("Enter width rectangle: ")
	fmt.Scanln(&w)
	fmt.Println("Enter height rectangle: ")
	fmt.Scanln(&h)
	s := w * h
	fmt.Println(s)

}
