package helper

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gin/constant"
	"gin/model/entity"
	"gin/model/param"
	mapSet "github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
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

// 断言是否是string类型
func IsString(i interface{}) bool {
	if _, ok := i.(string); ok {
		return true
	}
	return false
}

// 断言是否是int类型 广义
func IsInt(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
		//case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		//	data = int(v.Uint())
	}
	return false
}

// 校验字段
func GetVerifyField(fields []string, getField string) (verifyField string) {
	if len(fields) == 0 || getField == "" {
		return
	}
	if getField == "*" {
		return "*"
	}
	getField = strings.Replace(getField, "，", ",", -1)
	getField = strings.Replace(getField, " ", "", -1)
	getFieldSlice := strings.Split(getField, ",")
	fieldsSet := mapSet.NewSet()
	for _, v := range fields {
		fieldsSet.Add(v)
	}
	getFieldSet := mapSet.NewSet()
	for _, v := range getFieldSlice {
		getFieldSet.Add(v)
	}
	intersectSet := fieldsSet.Intersect(getFieldSet)

	var verifyFieldSet []string
	intersectSet.Each(func(i interface{}) bool {
		iToStr := fmt.Sprintf("%v", i)
		if tmp := strings.Split(iToStr, "."); len(tmp) == 2 {
			verifyFieldSet = append(verifyFieldSet, tmp[1])
		} else {
			verifyFieldSet = append(verifyFieldSet, iToStr)
		}
		return false
	})

	verifyField = strings.Join(verifyFieldSet, ",")
	return
}

// 合并多个map
func MergeMap(maps ...constant.BaseMap) constant.BaseMap {
	newMap := make(constant.BaseMap)
	switch len(maps) {
	case 0:
		return newMap
	case 1:
		return maps[0]
	default:
	}
	merge := func(map1 constant.BaseMap, map2 constant.BaseMap) constant.BaseMap {
		for k, v := range map2 {
			map1[k] = v
		}
		return map1
	}
	for _, m := range maps {
		newMap = merge(newMap, m)
	}
	return newMap
}

func CopyFields(a interface{}, b interface{}, fields ...string) (err error) {
	at := reflect.TypeOf(a)
	av := reflect.ValueOf(a)
	bt := reflect.TypeOf(b)
	bv := reflect.ValueOf(b)

	// 简单判断下
	if at.Kind() != reflect.Ptr {
		err = fmt.Errorf("a must be a struct pointer")
		return
	}
	av = reflect.ValueOf(av.Interface())

	// 要复制哪些字段
	_fields := make([]string, 0)
	if len(fields) > 0 {
		_fields = fields
	} else {
		for i := 0; i < bv.NumField(); i++ {
			_fields = append(_fields, bt.Field(i).Name)
		}
	}

	if len(_fields) == 0 {
		fmt.Println("no fields to copy")
		return
	}

	// 复制
	for i := 0; i < len(_fields); i++ {
		name := _fields[i]
		f := av.Elem().FieldByName(name)
		bValue := bv.FieldByName(name)
		// a中有同名的字段并且类型一致才复制
		if f.IsValid() && f.Kind() == bValue.Kind() && f.Type() == bValue.Type() {
			f.Set(bValue)
		} else {
			fmt.Printf("no such field or different kind, fieldName: %s\n", name)
		}
	}
	return
}

func IsStruct(s interface{}) bool {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// uninitialized zero value of a struct
	if v.Kind() == reflect.Invalid {
		return false
	}

	return v.Kind() == reflect.Struct
}

// 结构体转map
func StructToMap(s interface{}) constant.BaseMap {
	data := make(constant.BaseMap)
	if !IsStruct(s) {
		return data
	}

	js, err := json.Marshal(s)
	if err != nil {
		return data
	}

	var toData interface{}
	if err := json.Unmarshal(js, &toData); err != nil {
		return data
	}
	return toData.(map[string]interface{})
}

// JSON转map
func JsonToMap(jsonString []byte) constant.BaseMap {
	var toData interface{}
	if err := json.Unmarshal(jsonString, &toData); err != nil {
		return make(constant.BaseMap)
	}
	return toData.(map[string]interface{})
}

func GenerateUuid() string {
	return uuid.NewV4().String()
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 基于传入的fields对struct字段进行过滤，最终转为map
func FilterStructByFields(s interface{}, fields string, allFields constant.FieldMap) constant.BaseMap {
	data := make(constant.BaseMap)

	if !IsStruct(s) {
		return data
	}

	if fields == "*" {
		fields = ""
		for _, fieldSlice := range allFields {
			fields += strings.Join(fieldSlice, ",")
		}
	}

	sType := reflect.TypeOf(s)
	sData := reflect.ValueOf(s)
	for i := 0; i < sData.NumField(); i++ {
		sValue := sData.Field(i)
		if sValue.Kind() == reflect.Ptr {
			sValue = sValue.Elem()
		}
		if !sValue.IsValid() || sValue.IsZero() {
			continue
		}

		// 读取子结构体tag
		sTag := sType.Field(i).Tag
		jsonTag := sTag.Get("json")
		exTableName := sTag.Get("exTableName")   // 前端传入的fields前缀
		exName := strings.Split(jsonTag, ",")[0] // 输出至前端的key

		// 获取匹配的字段
		if exTableName == "" {
			exTableName = "base"
		}
		getFields := strings.Split(GetVerifyField(allFields[exTableName], fields), ",")
		if len(getFields) == 0 {
			continue
		}
		// 存在子结构体和子结构体切片两种可能
		switch sValue.Kind() {
		case reflect.Struct:
			structValues := getStructKeyValues(sValue.Interface(), getFields)
			if exName != "" {
				data[exName] = structValues
			} else {
				data = MergeMap(data, structValues)
			}
		case reflect.Slice:
			sliceData := make([]constant.BaseMap, 0)
			sliceReflect := reflect.ValueOf(sValue.Interface())
			for i := 0; i < sliceReflect.Len(); i++ {
				sliceData = append(sliceData, getStructKeyValues(sliceReflect.Index(i).Interface(), getFields))
			}
			data[exName] = sliceData
		default:
			continue
		}
	}
	return data
}

// 列表，基于传入的fields对struct字段进行过滤，最终转为map
func FilterStructsByFields(s interface{}, fields string, allFields constant.FieldMap) []constant.BaseMap {
	data := make([]constant.BaseMap, 0)
	if !IsSlice(s) {
		return data
	}

	sData := reflect.ValueOf(s)
	for i := 0; i < sData.Len(); i++ {
		singleData := FilterStructByFields(sData.Index(i).Interface(), fields, allFields)
		if IsEmpty(singleData) {
			continue
		}
		data = append(data, singleData)
	}

	return data
}

// 读取结构体字段名 切片
// 优先读取json:xxx，其次读取结构体本身的字段名
func GetStructKeys(s interface{}, prefix string) (data []string) {
	if !IsStruct(s) {
		return data
	}

	sType := reflect.TypeOf(s)

	for i := 0; i < sType.NumField(); i++ {
		jsonTag := sType.Field(i).Tag.Get("json")
		columnName := strings.Split(jsonTag, ",")[0]
		if columnName == "" {
			columnName = sType.Field(i).Name
		}
		if prefix != "" {
			columnName = strings.Join([]string{prefix, columnName}, ".")
		}
		data = append(data, columnName)
	}

	return data
}

// 基于给定的字段名，读取结构体中对应的字段kv键值对
func getStructKeyValues(s interface{}, fields []string) constant.BaseMap {
	data := make(constant.BaseMap)
	if !IsStruct(s) {
		return data
	}

	sData := reflect.ValueOf(s)
	if sData.Kind() == reflect.Ptr {
		sData = sData.Elem()
	}
	sType := sData.Type()

	for i := 0; i < sType.NumField(); i++ {
		jsonTag := sType.Field(i).Tag.Get("json")
		columnName := strings.Split(jsonTag, ",")[0]
		if columnName == "" {
			columnName = sType.Field(i).Name
		}
		if !InSlice(columnName, fields) {
			continue
		}
		data[columnName] = sData.Field(i).Interface()
		if fmtDateTime, ok := data[columnName].(entity.DateTime); ok {
			data[columnName] = fmtDateTime.String()
		}
	}
	return data
}

// 读取string，为空返回默认值
func GetString(s, defaultS string) string {
	if len(s) == 0 {
		return defaultS
	}
	return s
}

// 读取int，为空返回默认值
func GetInt(i, defaultI int) int {
	if i == 0 {
		return defaultI
	}
	return i
}

func GetInterfaceSliceByString(i []string) (data []interface{}) {
	for _, v := range i {
		data = append(data, v)
	}
	return
}

func GetInterfaceSliceByInt(i []int) (data []interface{}) {
	for _, v := range i {
		data = append(data, v)
	}
	return
}

// 添加token参数
func AppendTokenParams(c *gin.Context, params *param.Common) {
	params.AppKey = c.GetString("AppKey")
	params.Channel = c.GetInt("Channel")

	return
}

// 任意形式的请求，读取对应的key值
func GetRequestStringByKey(c *gin.Context, key string) string {
	if queryData := c.Query(key); queryData != "" {
		return queryData
	}
	if postFormData := c.PostForm(key); postFormData != "" {
		return postFormData
	}

	data, _ := c.GetRawData()
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	var jsonData interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return ""
	}

	if data, ok := jsonData.(map[string]interface{})[key]; ok {
		return fmt.Sprintf("%v", data)
	}
	return ""
}

func InterfaceToInt(i interface{}) (data int) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.String:
		data, _ = strconv.Atoi(v.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		data = int(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		data = int(v.Uint())
	case reflect.Float32, reflect.Float64:
		data = int(v.Float())
	}
	return
}
