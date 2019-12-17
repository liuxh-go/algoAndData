package sort

import "fmt"

// 一、冒泡排序
type Bubble struct {
	data []int
}

func (b *Bubble) Init(data []int) {
	b.data = data
}

func (b *Bubble) Sort() {
	l := len(b.data)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1; j++ {
			if b.data[j+1] < b.data[j] {
				b.data[j], b.data[j+1] = b.data[j+1], b.data[j]
			}
		}
	}
}

func (b *Bubble) Print() {
	for _, num := range b.data {
		fmt.Printf("%d -> ", num)
	}
}
