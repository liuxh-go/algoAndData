package sort

// 希尔排序
type Hill struct {
	data []int
}

func (this *Hill) Name() string {
	return "hill"
}

func (this *Hill) Init(data []int) {
	this.data = data
}

//0. 选择增量序列,一般选择gap=length/2,gap/2,gap/4...
//1. 按照增量拆分待排序序列,对每组序列分别进行插入排序
//2. 选择下一个增量,重复`1`过程直到增量为1
//3. 完成排序
func (this *Hill) Sort() {
	data, l := this.data, len(this.data)

	for gap := l / 2; gap > 0; gap /= 2 {
		insertSort(data, gap, l/gap)
	}
}

func insertSort(data []int, gap, count int) {
	for i := gap; i < count; i += gap {
		for j := i; j > 0; j -= gap {
			if data[j] < data[j-gap] {
				data[j], data[j-gap] = data[j-gap], data[j]
			} else {
				break
			}
		}
	}
}

func (this *Hill) Print(printFunc func(format string, args ...interface{})) {
	printSlice(this.data, printFunc)
}

func (this *Hill) Sorted() bool {
	return sorted(this.data)
}
