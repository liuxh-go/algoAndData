package sort

type Sorter interface {
	Name() string
	Init([]int)
	Sort()
	Print(func(format string, args ...interface{}))
	Sorted() bool
}
