package other

import (
	"math/rand"
	"time"
)

// 随机产生[0-n)的不重复数组
func RandSlice(n int) []int {
	result := make([]int, n)

	randObj := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < n; i++ {
		randNum := randObj.Intn(i + 1)
		result[i], result[randNum] = result[randNum], i
	}

	return result
}
