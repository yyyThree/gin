package helper

import (
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

// 基于反射，校验任意值是否为空
func IsEmpty(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

// 校验传入值是否至少存在一个空值
func HasAnyEmpty(list ...interface{}) bool {
	for _, v := range list {
		if IsEmpty(v) {
			return true
		}
	}
	return false
}

func Rand(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max - min) + min
}

// 创建唯一ID
// 时间戳 + 随机值
func Uuid() string {
	return time.Now().Format("2006012150405") + strconv.Itoa(Rand(1000, 9999))
}
