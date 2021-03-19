package constant

// 商品状态
const (
	ItemStateNormal = iota
	ItemStateDeleted
	ItemStateFinalDeleted
)

// sku状态
const (
	SkuStateNormal = iota
	SkuStateDeleted
	SkuStateFinalDeleted
	SkuStateDeletedSelf
)

// 商品删除类型
const (
	ItemDelete = iota
	ItemFinalDelete
)
