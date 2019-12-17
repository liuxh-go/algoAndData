package sort

// 一、冒泡排序
type Bubble struct {
	data []int
}

func (this *Bubble) Name() string {
	return "bubble"
}

func (this *Bubble) Init(data []int) {
	this.data = data
}

//0. 初始化n为0
//1. 第n个数和第n+1个数进行比较,如果第n+1个数小于第n个数,那么把第n个数和第n+1个数进行交换
//2. 递增n(即n加1),重复`1`步骤直到遍历完整个数组
//3. 根据数组长度重复`0`步骤(长度为l则重复l次)
//4. 完成排序
func (this *Bubble) Sort() {
	data := this.data
	l := len(data)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1; j++ {
			if data[j+1] < data[j] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func (this *Bubble) Print() {
	printSlice(this.data)
}

func (this *Bubble) Sorted() bool {
	return sorted(this.data)
}
