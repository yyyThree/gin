// Code generated by "stringer -type Code -linecomment -output common_string.go"; DO NOT EDIT.

package code

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OK-1]
	_ = x[ApiNotFound-404]
	_ = x[NoAuthorization-1001]
	_ = x[AuthorizationErr-1002]
	_ = x[TokenNotFound-1003]
	_ = x[TokenMalformed-1004]
	_ = x[TokenExpired-1005]
	_ = x[TokenNotValidYet-1006]
	_ = x[TokenNotValid-1007]
	_ = x[ParamBindErr-2001]
	_ = x[IllegalParams-2002]
	_ = x[IllegalJsonTypeString-2003]
	_ = x[RecordNotFound-3001]
	_ = x[MySqlErr-4001]
	_ = x[ServerErr-5001]
	_ = x[ItemInsertFail-10000]
	_ = x[SkuInsertFail-10001]
	_ = x[ItemNoFound-10002]
	_ = x[ItemStateError-10003]
	_ = x[ItemUpdateFail-10004]
	_ = x[SkuUpdateFail-10005]
	_ = x[SkuDelFail-10006]
	_ = x[ItemDelFail-10007]
	_ = x[ItemRecoverFail-10008]
	_ = x[ItemSearchInsertFail-10009]
	_ = x[ItemSearchUpdateFail-10010]
}

const (
	_Code_name_0 = "成功"
	_Code_name_1 = "接口不存在"
	_Code_name_2 = "未获取到AuthorizationAuthorization非法token不存在token解析失败token已失效token未生效token无效"
	_Code_name_3 = "参数绑定失败参数非法非法json字符串"
	_Code_name_4 = "未查询到记录"
	_Code_name_5 = "mysql执行错误"
	_Code_name_6 = "服务器错误"
	_Code_name_7 = "商品添加失败sku添加失败商品不存在商品状态不正确商品更新失败sku更新失败sku删除失败商品删除失败商品恢复失败商品搜索数据添加失败商品搜索数据更新失败"
)

var (
	_Code_index_2 = [...]uint8{0, 25, 44, 58, 75, 89, 103, 114}
	_Code_index_3 = [...]uint8{0, 18, 30, 49}
	_Code_index_7 = [...]uint8{0, 18, 33, 48, 69, 87, 102, 117, 135, 153, 183, 213}
)

func (i Code) String() string {
	switch {
	case i == 1:
		return _Code_name_0
	case i == 404:
		return _Code_name_1
	case 1001 <= i && i <= 1007:
		i -= 1001
		return _Code_name_2[_Code_index_2[i]:_Code_index_2[i+1]]
	case 2001 <= i && i <= 2003:
		i -= 2001
		return _Code_name_3[_Code_index_3[i]:_Code_index_3[i+1]]
	case i == 3001:
		return _Code_name_4
	case i == 4001:
		return _Code_name_5
	case i == 5001:
		return _Code_name_6
	case 10000 <= i && i <= 10010:
		i -= 10000
		return _Code_name_7[_Code_index_7[i]:_Code_index_7[i+1]]
	default:
		return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
