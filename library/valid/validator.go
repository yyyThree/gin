package valid

import (
	"gin/constant"
	"gin/helper"
	"github.com/asaskevich/govalidator"
	"reflect"
)

var errMsgMap = constant.StringMap{
	"NotEmpty":          "校验不为空失败",
	"IsIn":              "校验参数范围失败",
	"IsAllInStruct":     "校验参数存在于结构体内失败",
	"Contains":          "校验字符串包含关系失败",
	"HasLowerCase":      "校验存在小写字母失败",
	"HasUpperCase":      "校验存在大写字母失败",
	"IsLowerCase":       "校验小写字符串失败",
	"IsUpperCase":       "校验大写字符串失败",
	"HasWhitespace":     "校验存在空格符失败",
	"HasWhitespaceOnly": "校验仅存在空格符失败",
	"IsAlpha":           "校验仅存在字母失败",
	"IsNumeric":         "校验仅存在数字失败",
	"IsAlphanumeric":    "校验仅存在字母和数字失败",
	"IsASCII":           "校验ascii码失败",
	"IsBase64":          "校验Base64格式失败",
	"IsEmail":           "校验邮箱格式失败",
	"IsFilePath":        "校验文件路径失败",
	"IsFloat":           "校验浮点数失败",
	"IsHash":            "校验哈希格式失败",
	"IsIP":              "校验IP格式失败",
	"IsDNSName":         "校验DNS地址格式失败",
	"IsJSON":            "校验JSON格式失败",
	"IsLatitude":        "校验纬度格式失败",
	"IsLongitude":       "校验经度格式失败",
	"IsMD5":             "校验MD5格式失败",
	"IsURL":             "校验URL格式失败",
	"InRange":           "校验参数范围失败",
	"StringLength":      "校验字符串长度失败",
	"MinStringLength":   "校验字符串长度最小值失败",
	"MaxStringLength":   "校验字符串长度最大值失败",
}

// =============自定义方法=============
// 任意值不为空
func (myValid *myValid) NotEmpty() *myValid {
	myValid.validations["NotEmpty"] = func() bool {
		return !helper.IsEmpty(myValid.value)
	}
	return myValid
}

// 校验值在切片内
func (myValid *myValid) IsIn(slice interface{}) *myValid {
	myValid.validations["IsIn"] = func() bool {
		if !helper.IsSlice(slice) {
			return false
		}
		sData := reflect.ValueOf(slice)
		for i := 0; i < sData.Len(); i++ {
			if sData.Index(i).Interface() == myValid.value {
				return true
			}
		}
		return false
	}
	return myValid
}

// 是否含有某个值
func (myValid *myValid) IsAllInStruct(containStruct interface{}) *myValid {
	myValid.validations["IsAllInStruct"] = func() bool {
		if !helper.IsStruct(containStruct) {
			return false
		}
		if !helper.IsSlice(myValid.value) {
			return false
		}
		sValue := reflect.ValueOf(myValid.value)
		// sType := reflect.TypeOf(myValid.value)
		cValue := reflect.ValueOf(containStruct)
		// cType := reflect.TypeOf(containStruct)
		for i := 0; i < sValue.Len(); i++ {
			f := cValue.FieldByName(sValue.Index(i).String())
			if f.IsValid() {
				continue
			}
			return false
		}
		return true
	}
	return myValid
}

// =====================================

// =============goValidator方法=========
// 字符串包含
func (myValid *myValid) Contains(subString string) *myValid {
	myValid.validations["Contains"] = func() bool {
		return govalidator.Contains(govalidator.ToString(myValid.value), subString)
	}
	return myValid
}

// 存在小写字母
func (myValid *myValid) HasLowerCase() *myValid {
	myValid.validations["HasLowerCase"] = func() bool {
		return govalidator.HasLowerCase(govalidator.ToString(myValid.value))
	}
	return myValid
}

// 存在大写字母
func (myValid *myValid) HasUpperCase() *myValid {
	myValid.validations["HasUpperCase"] = func() bool {
		return govalidator.HasUpperCase(govalidator.ToString(myValid.value))
	}
	return myValid
}

func (myValid *myValid) IsLowerCase() *myValid {
	myValid.validations["IsLowerCase"] = func() bool {
		return govalidator.IsLowerCase(govalidator.ToString(myValid.value))
	}
	return myValid
}

func (myValid *myValid) IsUpperCase() *myValid {
	myValid.validations["IsUpperCase"] = func() bool {
		return govalidator.IsUpperCase(govalidator.ToString(myValid.value))
	}
	return myValid
}

// 存在空格
func (myValid *myValid) HasWhitespace() *myValid {
	myValid.validations["HasWhitespace"] = func() bool {
		return govalidator.HasWhitespace(govalidator.ToString(myValid.value))
	}
	return myValid
}

// 仅空格
func (myValid *myValid) HasWhitespaceOnly() *myValid {
	myValid.validations["HasWhitespaceOnly"] = func() bool {
		return govalidator.HasWhitespaceOnly(govalidator.ToString(myValid.value))
	}
	return myValid
}

// 仅字母
func (myValid *myValid) IsAlpha() *myValid {
	myValid.validations["IsAlpha"] = func() bool {
		return govalidator.IsAlpha(govalidator.ToString(myValid.value))
	}
	return myValid
}

// 仅数字
func (myValid *myValid) IsNumeric() *myValid {
	myValid.validations["IsNumeric"] = func() bool {
		return govalidator.IsNumeric(govalidator.ToString(myValid.value))
	}
	return myValid
}

// 仅字母+数组
func (myValid *myValid) IsAlphanumeric() *myValid {
	myValid.validations["IsAlphanumeric"] = func() bool {
		return govalidator.IsAlphanumeric(govalidator.ToString(myValid.value))
	}
	return myValid
}

// 仅ascii码
func (myValid *myValid) IsASCII() *myValid {
	myValid.validations["IsASCII"] = func() bool {
		return govalidator.IsASCII(govalidator.ToString(myValid.value))
	}
	return myValid
}

func (myValid *myValid) IsBase64() *myValid {
	myValid.validations["IsBase64"] = func() bool {
		return govalidator.IsBase64(govalidator.ToString(myValid.value))
	}
	return myValid
}

func (myValid *myValid) IsEmail() *myValid {
	myValid.validations["IsEmail"] = func() bool {
		return govalidator.IsEmail(govalidator.ToString(myValid.value))
	}
	return myValid
}

func (myValid *myValid) IsFilePath() *myValid {
	myValid.validations["IsFilePath"] = func() bool {
		ok, _ := govalidator.IsFilePath(govalidator.ToString(myValid.value))
		return ok
	}
	return myValid
}

// string是否可以转化为float
func (myValid *myValid) IsFloat() *myValid {
	myValid.validations["IsFloat"] = func() bool {
		return govalidator.IsFloat(govalidator.ToString(myValid.value))
	}
	return myValid
}

func (myValid *myValid) IsHash(algorithm string) *myValid {
	myValid.validations["IsHash"] = func() bool {
		return govalidator.IsHash(govalidator.ToString(myValid.value), algorithm)
	}
	return myValid
}

func (myValid *myValid) IsIP() *myValid {
	myValid.validations["IsIP"] = func() bool {
		return govalidator.IsIP(govalidator.ToString(myValid.value))
	}
	return myValid
}

func (myValid *myValid) IsDNSName() *myValid {
	myValid.validations["IsDNSName"] = func() bool {
		return govalidator.IsDNSName(govalidator.ToString(myValid.value))
	}
	return myValid
}

func (myValid *myValid) IsJSON() *myValid {
	myValid.validations["IsJSON"] = func() bool {
		return govalidator.IsJSON(govalidator.ToString(myValid.value))
	}
	return myValid
}

func (myValid *myValid) IsLatitude() *myValid {
	myValid.validations["IsLatitude"] = func() bool {
		return govalidator.IsLatitude(govalidator.ToString(myValid.value))
	}
	return myValid
}

func (myValid *myValid) IsLongitude() *myValid {
	myValid.validations["IsLongitude"] = func() bool {
		return govalidator.IsLongitude(govalidator.ToString(myValid.value))
	}
	return myValid
}

func (myValid *myValid) IsMD5() *myValid {
	myValid.validations["IsMD5"] = func() bool {
		return govalidator.IsMD5(govalidator.ToString(myValid.value))
	}
	return myValid
}

func (myValid *myValid) IsURL() *myValid {
	myValid.validations["IsURL"] = func() bool {
		return govalidator.IsURL(govalidator.ToString(myValid.value))
	}
	return myValid
}

// 校验值是否在范围内，支持 int, float32, float64 and string
func (myValid *myValid) InRange(left interface{}, right interface{}) *myValid {
	myValid.validations["InRange"] = func() bool {
		return govalidator.InRange(myValid.value, left, right)
	}
	return myValid
}

// 字符串长度范围
func (myValid *myValid) StringLengthRange(min, max int) *myValid {
	myValid.validations["StringLengthRange"] = func() bool {
		return govalidator.IsByteLength(govalidator.ToString(myValid.value), min, max)
	}
	return myValid
}

// 字符串长度
func (myValid *myValid) StringLength(length int) *myValid {
	myValid.validations["StringLength"] = func() bool {
		return govalidator.IsByteLength(govalidator.ToString(myValid.value), length, length)
	}
	return myValid
}

// 字符串最小长度，不小于此长度
func (myValid *myValid) MinStringLength(length int) *myValid {
	myValid.validations["MinStringLength"] = func() bool {
		return govalidator.MinStringLength(govalidator.ToString(myValid.value), govalidator.ToString(length))
	}
	return myValid
}

// 字符串最大长度，不大于此长度
func (myValid *myValid) MaxStringLength(length int) *myValid {
	myValid.validations["MaxStringLength"] = func() bool {
		return govalidator.MaxStringLength(govalidator.ToString(myValid.value), govalidator.ToString(length))
	}
	return myValid
}

// =====================================
