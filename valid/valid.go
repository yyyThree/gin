package valid

import "strings"

type myValid struct {
	key        string
	value      interface{}
	validations map[string]func() bool
}

var allMyValid []*myValid

func New(key string, value interface{}) *myValid {
	myValid := &myValid{
		key: key,
		value: value,
		validations: make(map[string]func() bool),
	}
	allMyValid = append(allMyValid, myValid)
	return myValid
}

// 执行所有校验方法
// TODO 语言包
func Valid() (errMsg string) {
	if len(allMyValid) == 0 {
		return
	}

	allMyValidNew := allMyValid
	allMyValid = allMyValid[0:0]
	for _, myValid := range allMyValidNew {
		for validName, validation := range myValid.validations {
			if !validation() {
				errMsg = strings.Join([]string{myValid.key, validName}, ":")
				return
			}
		}
	}
	return
}
