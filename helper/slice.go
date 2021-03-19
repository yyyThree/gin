package helper

import (
	"reflect"
)

// 判断某个值是否含在切片中
func InSlice(searchVal interface{}, searchSlice interface{}) (exist bool) {
	exist = false
	// index = -1
	sValue := reflect.ValueOf(searchVal)

	switch reflect.TypeOf(searchSlice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(searchSlice)
		for i := 0; i < s.Len(); i++ {
			// 如果都是广义的int类型 则只进行值比较
			ssValue := reflect.ValueOf(s.Index(i).Interface())
			if sValue.Kind() != ssValue.Kind() && IsInt(sValue.Interface()) && IsInt(ssValue.Interface()) {
				tmpSearchValue := sValue.Int()
				tmpSearchSliceValue := s.Index(i).Int()
				if reflect.DeepEqual(tmpSearchValue, tmpSearchSliceValue) == true {
					// index = i
					exist = true
					return
				}
			} else {
				if reflect.DeepEqual(searchVal, s.Index(i).Interface()) == true {
					// index = i
					exist = true
					return
				}
			}
		}
	}
	return
}

func IsSlice(s interface{}) bool {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// uninitialized zero value of a struct
	if v.Kind() == reflect.Invalid {
		return false
	}

	return v.Kind() == reflect.Slice
}

// slice 去重
func SliceUnique(sliceData interface{}) (ret []interface{}) {
	if !IsSlice(sliceData) {
		return ret
	}
	sliceValue := reflect.ValueOf(sliceData)
	var isExist = false
	for i := 0; i < sliceValue.Len(); i++ {
		isExist = false
		if len(ret) == 0 {
			ret = append(ret, sliceValue.Index(i).Interface())
			continue
		}
		for j := 0; j < len(ret); j++ {
			if reflect.DeepEqual(sliceValue.Index(i).Interface(), ret[j]) {
				isExist = true
				continue
			}
		}
		if !isExist {
			ret = append(ret, sliceValue.Index(i).Interface())
		}
	}
	return ret
}

// 剔除slice元素值，如果存在多个相同则全部剔除
func SliceDelByValue(s interface{}, slice interface{}) (newSlice []interface{}) {
	if !IsSlice(slice) {
		return newSlice
	}

	sliceData := reflect.ValueOf(slice)

	for i := 0; i < sliceData.Len(); i++ {
		if sliceData.Index(i).Interface() != s {
			newSlice = append(newSlice, sliceData.Index(i).Interface())
		}
	}
	return
}

// 比较两个slice，返回差集
func SliceDiff(slice1, slice2 interface{}) (newSlice []interface{}) {
	if !IsSlice(slice1) || !IsSlice(slice2) {
		return newSlice
	}

	slice1Data := reflect.ValueOf(slice1)
	slice2Data := reflect.ValueOf(slice2)

	for i := 0; i < slice1Data.Len(); i++ {
		newSlice = append(newSlice, slice1Data.Index(i).Interface())
	}

	for i := 0; i < slice2Data.Len(); i++ {
		if InSlice(slice2Data.Index(i).Interface(), slice1) {
			newSlice = SliceDelByValue(slice2Data.Index(i).Interface(), newSlice)
		}
	}
	return
}

// 比较两个slice，返回交集
func SliceIntersect(slice1, slice2 interface{}) (newSlice []interface{}) {
	if !IsSlice(slice1) || !IsSlice(slice2) {
		return newSlice
	}

	slice2Data := reflect.ValueOf(slice2)

	for i := 0; i < slice2Data.Len(); i++ {
		if InSlice(slice2Data.Index(i).Interface(), slice1) {
			newSlice = append(newSlice, slice2Data.Index(i).Interface())
		}
	}
	return
}
