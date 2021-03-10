package code

//go:generate stringer -type Code -linecomment -output common_string.go

// 商品 号段 [10000 - 19999]
const (
	ItemIdNotFound Code = 10000 // 商品ID不存在
)
