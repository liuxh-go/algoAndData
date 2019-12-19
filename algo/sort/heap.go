package sort

// 堆排序
type Heap struct {
	data []int
}

func (this *Heap) Name() string {
	return "heap"
}

func (this *Heap) Init(data []int) {
	this.data = data
}

//0. 将待排序序列构建成一个堆
//1. 将堆调整为小顶堆
//2. 将R[0]即堆顶元素和R[n]即堆底元素互换,将R[n]从堆中移除并添加到有序序列中
//3. n减少1,从`1`步骤重复直到n=0
//4. 完成排序
func (this *Heap) Sort() {
	for i := 0; i < len(this.data); i++ {
		this.adjustHeap(i)
	}
}

func (this *Heap) adjustHeap(start int) {
	data := this.data[start:]
	l := len(data)
	if l <= 0 {
		return
	}

	// 小顶堆:满足公式 R[i] < R[2*i+1] && R[i] < R[2*i+2]
	for i := l/2 - 1; i >= 0; i-- {
		minIndex := i

		if 2*i+1 < l && data[2*i+1] < data[minIndex] {
			minIndex = 2*i + 1
		}

		if 2*i+2 < l && data[2*i+2] < data[minIndex] {
			minIndex = 2*i + 2
		}

		if minIndex != i {
			data[i], data[minIndex] = data[minIndex], data[i]
		}
	}
}

func (this *Heap) Print(printFunc func(format string, args ...interface{})) {
	printSlice(this.data, printFunc)
}

func (this *Heap) Sorted() bool {
	return sorted(this.data)
}
