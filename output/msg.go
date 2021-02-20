package output

import (
	"gin/config"
	"gin/constant"
	"strconv"
	"strings"

	"gin/output/language/en"
	"gin/output/language/zh"
)

type msgListType map[string]constant.MsgMap

type languageListType map[string]msgListType

// 语言包列表
var languageList = languageListType{
	"zh": {
		"Common": zh.Common,
		"Item":   zh.Item,
	},
	"en": {
		"Common": en.Common,
		"Item":   en.Item,
	},
}

// 基于配置的语言，获取完整的语言包列表
func getMsgList() (msgListType, error) {
	if _, ok := languageList[config.Config.Language]; !ok {
		return make(msgListType), nil
	}
	return languageList[config.Config.Language], nil
}

// 查找语言包msg
func getMsg(msgCode string) (msg string) {
	if len(msgCode) == 0 {
		return
	}

	// 解析 msgCode
	msgCodeSlice := strings.Split(msgCode, ".")
	if len(msgCodeSlice) != 3 {
		return
	}

	// 获取所有语言包列表
	msgList, err := getMsgList()
	if err != nil {
		return
	}

	for name, subMsgList := range msgList {
		if name != msgCodeSlice[0] {
			continue
		}
		// string转int，得到最终状态码
		msgState, err := strconv.Atoi(msgCodeSlice[2])
		if err != nil {
			continue
		}
		if _, ok := subMsgList[msgCodeSlice[1]][msgState]; ok {
			msg = subMsgList[msgCodeSlice[1]][msgState]
			break
		}
	}
	return
}
