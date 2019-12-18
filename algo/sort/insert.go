package sort

// 插入排序
type Insert struct {
	data []int
}

func (this *Insert) Name() string {
	return "insert"
}

func (this *Insert) Init(data []int) {
	this.data = data
}

//0. 取出R[0],默认为有序序列,初始化n=1
//1. 取出R[n],将R[n]循环和R[0,n)中的元素比较,若R[n]大于比较的元素,则将R[n]插入在比较元素的下一个位置
//2. n增加1,重复`1`步骤直到遍历完整个序列
//3. 完成排序
func (this *Insert) Sort() {
	data, l := this.data, len(this.data)

	for i := 1; i < l; i++ {
		for j := i; j > 0; j-- {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			} else {
				break
			}
		}
	}
}

func (this *Insert) Print(printFunc func(format string, args ...interface{})) {
	printSlice(this.data, printFunc)
}

func (this *Insert) Sorted() bool {
	return sorted(this.data)
}
