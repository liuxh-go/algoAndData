package sort

// 快速排序
type Quick struct {
	data []int
}

func (this *Quick) Name() string {
	return "quick"
}

func (this *Quick) Init(data []int) {
	this.data = data
}

//0. 分区:从序列中选出一个关键字作为基准值,遍历序列,将比基准值小的项全部放到基准值左边,比基准值大的项放到右边
//1. 将`0`步骤得到的左分区和右分区序列递归的再进行`0`步骤的操作,直到分区为1
//2. 完成排序
func (this *Quick) Sort() {
	quickSort(this.data, 0, len(this.data))
}

func quickSort(data []int, start, end int) {
	if end <= start {
		return
	}

	index := partition(data, start, end)

	// 递归排序左右分区
	quickSort(data, start, index)
	quickSort(data, index+1, end)
}

func partition(data []int, start, end int) int {
	// 将比基准值小的项都放到左边
	benchmarkValue := data[start]
	for i := start + 1; i < end; i++ {
		if data[i] > benchmarkValue {
			continue
		}

		j := i - 1
		for ; j > start; j-- {
			if data[j] <= benchmarkValue {
				break
			}
		}

		data[i], data[j+1] = data[j+1], data[i]
	}

	// 找到基准值应该在的位置
	index := start
	for ; index < end; index++ {
		if data[index] > benchmarkValue {
			break
		}
	}
	data[start], data[index-1] = data[index-1], data[start]

	return index - 1
}

func (this *Quick) Print(printFunc func(format string, args ...interface{})) {
	printSlice(this.data, printFunc)
}

func (this *Quick) Sorted() bool {
	return sorted(this.data)
}
