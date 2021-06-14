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

	ShellSort(sliceNum)
	fmt.Println(sliceNum)
}
func ShellSort(ar []int) {
	for gap := len(ar) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(ar); i++ {
			x := ar[i]
			j := i
			for ; j >= gap && ar[j-gap] > x; j -= gap {
				ar[j] = ar[j-gap]
			}
			ar[j] = x
		}
	}
}
