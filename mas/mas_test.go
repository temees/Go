package main

import "testing"

/*import (
	"reflect"
	"testing"
)

func TestMas(t *testing.T) {
	table := []struct {
		arg  []int
		want []int
	}{

		{[]int{6, 3, 8, 7, 1, 2}, []int{1, 2, 3, 6, 7, 8}},
		{[]int{8, 2, 9, 7, 3, 4}, []int{2, 3, 4, 7, 8, 9}},
		{[]int{4, 3, 5, 8, 6, 7}, []int{3, 4, 5, 6, 7, 8}},
	}

	for _, entry := range table {
		got := InsertionSort(entry.arg)

		if  !reflect.DeepEqual(got, entry.want) {
			t.Errorf("for %d got %d want %d", entry.arg, got, entry.want)
		}
	}
}*/

const N = 20

var sink []int
var Ar = []int{2, 8, 3, 9, 4}

func BenchmarkMas(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = InsertionSort(Ar)
	}
	sink = res
}
