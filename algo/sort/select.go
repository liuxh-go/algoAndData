package sort

// 选择排序
type Select struct {
	data []int
}

func (this *Select) Name() string {
	return "select"
}

func (this *Select) Init(data []int) {
	this.data = data
}

//0. 初始化i=0
//1. 划分有序序列R[0,i)和无序序列R[i,n),R为待排序序列
//2. 从无序序列R[i,n)中选出最小的数,下标为min,将R[min]和R[i]进行交换
//3. i增加1,继续步骤`2`直到i=n
//4. 完成排序
func (this *Select) Sort() {
	data := this.data
	l := len(data)
	for i := 0; i < l-1; i++ {
		min := i
		for j := i; j < l; j++ {
			if data[j] < data[min] {
				min = j
			}
		}

		data[min], data[i] = data[i], data[min]
	}
}

func (this *Select) Print() {
	printSlice(this.data)
}

func (this *Select) Sorted() bool {
	return sorted(this.data)
}
