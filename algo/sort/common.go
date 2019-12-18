package sort

import (
	"fmt"
	"strings"
)

func printSlice(data []int, printFunc func(format string, args ...interface{})) {
	dataStrSlice := make([]string, 0, len(data))
	for _, num := range data {
		dataStrSlice = append(dataStrSlice, fmt.Sprintf("%d", num))
	}

	printFunc(strings.Join(dataStrSlice, "->"))
}

func sorted(data []int) bool {
	for i := 0; i < len(data)-1; i++ {
		if data[i] > data[i+1] {
			return false
		}
	}

	return true
}
