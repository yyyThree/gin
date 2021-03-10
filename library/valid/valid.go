package valid

import (
	"strings"
)

type myValid struct {
	key         string
	value       interface{}
	validations map[string]func() bool
}

type allMyValid []*myValid

func New() *allMyValid {
	return &allMyValid{}
}

func (allMyValid *allMyValid) Append(myValid *myValid) *allMyValid {
	*allMyValid = append(*allMyValid, myValid)
	return allMyValid
}

func One(key string, value interface{}) *myValid {
	myValid := &myValid{
		key:         key,
		value:       value,
		validations: make(map[string]func() bool),
	}
	return myValid
}

// 执行所有校验方法
// TODO 语言包
func (allMyValid allMyValid) Valid() (errMsg string) {
	if len(allMyValid) == 0 {
		return
	}

	for _, myValid := range allMyValid {
		for validName, validation := range myValid.validations {
			if !validation() {
				errMsg = strings.Join([]string{myValid.key, errMsgMap[validName]}, ":")
				return
			}
		}
	}
	return
}
