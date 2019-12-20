package other

import (
	"math/rand"
	"time"
)

// 随机产生[start,start+n)的不重复数组
func RandSlice(start, n int) []int {
	result := make([]int, n)

	randObj := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < n; i++ {
		randNum := randObj.Intn(i + 1)
		result[i], result[randNum] = result[randNum], i
	}

	if start != 0 {
		for i := 0; i < n; i++ {
			result[i] += start
		}
	}

	return result
}
