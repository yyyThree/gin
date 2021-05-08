package helper

import (
	"math/rand"
	"time"
)

func Max(i int, j int) int {
	if i < j {
		return j
	}
	return i
}

// 生成范围随机值
func Rand(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}
