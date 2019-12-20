package sort

// 计数排序
type Count struct {
	data []int
}

func (this *Count) Name() string {
	return "count"
}

func (this *Count) Init(data []int) {
	this.data = data
}

//0. 找出待排序序列中最小(Min)和最大的数(Max)
//1. 构建辅助数组F[Min,Max],初始化i=0
//2. 统计,F[R[i]]++
//3. i增加1,从`2`步骤重复直到i=n
//4. 遍历F[Min,Max],填充结果数组
//5. 完成排序
func (this *Count) Sort() {
	data := this.data
	min, max := data[0], data[0]
	for _, item := range data[1:] {
		if item < min {
			min = item
		}

		if item > max {
			max = item
		}
	}

	f := make([]int, max-min+1)
	for _, item := range data {
		f[item-min]++
	}

	for index, i := 0, 0; i < len(f); {
		if f[i] <= 0 {
			i++
			continue
		}

		data[index] = i + min
		index++
		f[i]--
	}
}

func (this *Count) Print(printFunc func(format string, args ...interface{})) {
	printSlice(this.data, printFunc)
}

func (this *Count) Sorted() bool {
	return sorted(this.data)
}
