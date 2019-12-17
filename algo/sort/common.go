package sort

import "fmt"

func printSlice(data []int) {
	for _, num := range data {
		fmt.Printf("%d -> ", num)
	}
}

func sorted(data []int) bool {
	for i := 0; i < len(data)-1; i++ {
		if data[i] > data[i+1] {
			return false
		}
	}

	return true
}
