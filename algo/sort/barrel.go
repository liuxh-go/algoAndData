package sort

// 桶排序
type Barrel struct {
	data []int
}

func (this *Barrel) Name() string {
	return "barrel"
}

func (this *Barrel) Init(data []int) {
	this.data = data
}

//0. 初始化桶的大小为size
//1. 找到待排序序列R[0,n)中的最小数Min和最大数Max,计算桶的个数为(Max-Min)/size+1
//2. 遍历R[0,n),根据映射函数将R[i]填入对应的桶中
//3. 遍历每个桶,如果桶的size为1,将桶中元素添加到结果序列中
//4. 如果桶的个数为1但桶的size大于1,则将size减1
//5. 递归调用`1`步骤将非空桶中的序列作为待排序序列继续进行分配
//6. 结果序列即为排序后的序列
func (this *Barrel) Sort() {
	this.data = barrelSort(this.data, 10)
}

func barrelSort(data []int, barrelSize int) []int {
	if data == nil || len(data) <= 1 {
		return data
	}

	min, max := data[0], data[0]
	for _, item := range data {
		if item < min {
			min = item
		}

		if item > max {
			max = item
		}
	}

	barrelCount := (max-min)/barrelSize + 1
	barrelSlice := make([][]int, barrelCount)
	resultSlice := make([]int, 0)

	// 填充桶
	for _, item := range data {
		barrelIndex := (item - min) / barrelSize
		barrelSlice[barrelIndex] = append(barrelSlice[barrelIndex], item)
	}

	for i := 0; i < barrelCount; i++ {
		if barrelSize == 1 {
			resultSlice = append(resultSlice, barrelSlice[i]...)
			continue
		}

		if barrelCount == 1 {
			barrelSize--
		}

		resultSlice = append(resultSlice, barrelSort(barrelSlice[i], barrelSize)...)
	}

	return resultSlice
}

func (this *Barrel) Print(printFunc func(format string, args ...interface{})) {
	printSlice(this.data, printFunc)
}

func (this *Barrel) Sorted() bool {
	return sorted(this.data)
}
