package sort

// 归并排序
type Merge struct {
	data []int
}

func (this *Merge) Name() string {
	return "merge"
}

func (this *Merge) Init(data []int) {
	this.data = data
}

//0. 将长度为n的序列拆分为两个长度为n/2的序列
//1. 分别对长度为n/2的序列进行排序
//2. 将排序后的两个序列进行合并得到长度为n的有序序列
//3. 递归从`0`步骤开始调用,直到序列无法拆分为止
//4. 完成排序
func (this *Merge) Sort() {
	this.data = mergeSort(this.data)
}

func mergeSort(slice []int) []int {
	l := len(slice)
	if l < 2 {
		return slice
	}

	return merge(mergeSort(slice[:l/2]), mergeSort(slice[l/2:]))
}

// 合并有序数组
func merge(sliceOne, sliceTwo []int) []int {
	totalLen := len(sliceOne) + len(sliceTwo)
	result := make([]int, totalLen)
	for i, j, k := 0, 0, 0; k < totalLen; k++ {
		if i >= len(sliceOne) {
			result[k] = sliceTwo[j]
			j++
		} else if j >= len(sliceTwo) {
			result[k] = sliceOne[i]
			i++
		} else if sliceOne[i] < sliceTwo[j] {
			result[k] = sliceOne[i]
			i++
		} else {
			result[k] = sliceTwo[j]
			j++
		}
	}

	return result
}

func (this *Merge) Print(printFunc func(format string, args ...interface{})) {
	printSlice(this.data, printFunc)
}

func (this *Merge) Sorted() bool {
	return sorted(this.data)
}
