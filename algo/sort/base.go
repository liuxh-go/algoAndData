package sort

import "math"

// 基数排序
type Base struct {
	data []int
}

func (this *Base) Name() string {
	return "base"
}

func (this *Base) Init(data []int) {
	this.data = data
}

//0. 找到最大的数Max,计算Max的位数为Digit,初始化辅助数组F[0,10),初始化i=1
//1. 遍历待排序序列,计算对应位的数R[d]=R[i]%10^i/10^(i-1),根据R[d]将数字填入辅助数组中F[R[d]]=R[i]
//2. 遍历辅助数组F[0,10),依次取出数字放入临时数组
//3. i增加1,将临时数组传入`1`步骤重复直到i=Digit
//4. 最后得到的临时数组即为排序好的序列
func (this *Base) Sort() {
	data := this.data
	max := data[0]
	for _, item := range data {
		if item > max {
			max = item
		}
	}

	digit := 0
	for max > 0 {
		max /= 10
		digit++
	}

	for i := 1; i <= digit; i++ {
		f := make([][]int, 10)
		// 填充辅助数组
		for _, item := range data {
			index := (item % int(math.Pow10(i))) / int(math.Pow10(i-1))
			f[index] = append(f[index], item)
		}

		// 从辅助数组中取出数字替换data
		for dataIndex, j := 0, 0; j < 10; j++ {
			if f[j] == nil || len(f[j]) <= 0 {
				continue
			}

			for k := 0; k < len(f[j]); k++ {
				data[dataIndex] = f[j][k]
				dataIndex++
			}
		}
	}
}

func (this *Base) Print(printFunc func(format string, args ...interface{})) {
	printSlice(this.data, printFunc)
}

func (this *Base) Sorted() bool {
	return sorted(this.data)
}
