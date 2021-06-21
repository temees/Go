package main

import (
	"fmt"
)

var res map[int]int

func main() {
	res = map[int]int{}
	ent()
}

func ent() {
	var n int
	fmt.Println("Input numeric: ")
	fmt.Scanln(&n)
	f := fib(n)
	fmt.Println(f)
	res[n] = f
	fmt.Println(res)
	ent()
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
