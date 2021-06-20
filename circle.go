package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	fmt.Println("Input S:")
	fmt.Scanln(&s)
	d := 2 * math.Sqrt(s/math.Pi)
	fmt.Println(d)
	l := d * math.Pi
	fmt.Println(l)
}
