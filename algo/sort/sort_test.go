package sort

import (
	"testing"
)

var (
	data = []int{1, 970, 10, 42, 50, 25, 20, 10, 6, 4, 5, 9, 24, 0, 12, 29, 38, 3, 49}
)

func Test_Sort(t *testing.T) {
	//sorter := &Bubble{}
	sorter := &Select{}
	sort(sorter, t)
}

func sort(sorter Sorter, t *testing.T) {
	sorter.Init(append([]int{}, data...))
	sorter.Sort()
	sorter.Print()
	if sorter.Sorted() {
		t.Log(sorter.Name() + " Sort Success!")
	} else {
		t.Error(sorter.Name() + " Sort Failed!")
	}
}
