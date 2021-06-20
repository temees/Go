package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num string
	fmt.Println("Input numerics with comma: ")
	fmt.Scanln(&num)
	mas := strings.Split(num, ",")
	var sliceNum []int
	for _, s := range mas {
		n, err := strconv.Atoi(s)
		if err == nil {
			sliceNum = append(sliceNum, n)
		}
	}

	InsertionSort(sliceNum)
	fmt.Println(sliceNum)
}
func InsertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
}
