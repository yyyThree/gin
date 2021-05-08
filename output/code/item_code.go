package code

//go:generate stringer -type Code -linecomment -output common_string.go

// 商品 号段 [10000 - 19999]
const (
	ItemInsertFail       Code = 10000 // 商品添加失败
	SkuInsertFail        Code = 10001 // sku添加失败
	ItemNoFound          Code = 10002 // 商品不存在
	ItemStateError       Code = 10003 // 商品状态不正确
	ItemUpdateFail       Code = 10004 // 商品更新失败
	SkuUpdateFail        Code = 10005 // sku更新失败
	SkuDelFail           Code = 10006 // sku删除失败
	ItemDelFail          Code = 10007 // 商品删除失败
	ItemRecoverFail      Code = 10008 // 商品恢复失败
	ItemSearchInsertFail Code = 10009 // 商品搜索数据添加失败
	ItemSearchUpdateFail Code = 10010 // 商品搜索数据更新失败
)
