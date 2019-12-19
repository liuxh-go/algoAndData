package sort

import (
	"testing"

	"wshhz.com/algoAndData/algo/other"
)

const (
	sliceLength = 10
)

var (
	data = other.RandSlice(sliceLength)
)

func Test_Sort(t *testing.T) {
	//sorter := &Bubble{}
	//sorter := &Select{}
	//sorter := &Quick{}
	//sorter := &Insert{}
	//sorter := &Hill{}
	//sorter := &Heap{}
	sorter := &Merge{}
	sort(sorter, t)
}

func sort(sorter Sorter, t *testing.T) {
	sorter.Init(append([]int{}, data...))
	t.Logf("Before Sort Data: ")
	sorter.Print(t.Logf)
	sorter.Sort()
	t.Logf("After Sort Data: ")
	sorter.Print(t.Logf)
	if sorter.Sorted() {
		t.Log(sorter.Name() + " Sort Success!")
	} else {
		t.Error(sorter.Name() + " Sort Failed!")
	}
}
