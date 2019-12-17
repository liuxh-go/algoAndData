package sort

type Sorter interface {
	Name() string
	Init([]int)
	Sort()
	Print()
	Sorted() bool
}
