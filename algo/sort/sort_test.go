package sort

import (
	"testing"
)

var (
	data = []int{1, 10, 50, 20, 6, 4, 9, 24, 12, 38, 49}
)

func Test_Sort(t *testing.T) {
	bubble := &Bubble{}
	sort(bubble)
}

func sort(sorter Sorter) {
	sorter.Init(append([]int{}, data...))
	sorter.Sort()
	sorter.Print()
}
