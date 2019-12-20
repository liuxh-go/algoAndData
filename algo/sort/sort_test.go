package sort

import (
	"testing"
	"time"

	"wshhz.com/algoAndData/algo/other"
)

const (
	startNum     = 0
	randSliceLen = 100
)

var (
	data = append(other.RandSlice(startNum, randSliceLen), other.RandSlice(startNum, randSliceLen)...)
)

func Test_Sort(t *testing.T) {
	sort(&Bubble{}, t)
	sort(&Select{}, t)
	sort(&Quick{}, t)
	sort(&Insert{}, t)
	sort(&Hill{}, t)
	sort(&Heap{}, t)
	sort(&Merge{}, t)
	sort(&Count{}, t)
	sort(&Barrel{}, t)
	sort(&Base{}, t)
}

func sort(sorter Sorter, t *testing.T) {
	sorter.Init(append([]int{}, data...))

	t.Logf("--------------------%s Sort Start--------------------", sorter.Name())
	t.Logf("Before Sort Data: ")
	sorter.Print(t.Logf)
	dtStart := time.Now()
	t.Logf("Start Time : %s", dtStart.String())
	sorter.Sort()
	dtEnd := time.Now()
	t.Logf("End Time : %s,Consume Time : %s", dtEnd.String(), dtEnd.Sub(dtStart))
	t.Logf("After Sort Data: ")
	sorter.Print(t.Logf)
	if sorter.Sorted() {
		t.Log(sorter.Name() + " Sort Success!")
	} else {
		t.Error(sorter.Name() + " Sort Failed!")
	}
	t.Logf("--------------------%s Sort End----------------------", sorter.Name())
}
