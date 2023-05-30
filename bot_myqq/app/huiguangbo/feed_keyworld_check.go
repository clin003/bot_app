package main

import (
	"strings"

	"gitee.com/lyhuilin/util"
)

// retBool=isOk 检查订阅词和屏蔽词，有屏蔽词返回false
func feedKeyworldCheck(msgText string) (retText, retFilter string, retBool bool) {
	retBool = true

	keyworldFilter := util.KeyworldListParse(getFeedKeyworldFilter())
	for _, v := range keyworldFilter {
		vc := v
		if strings.Contains(msgText, vc) {
			retFilter = vc
			retBool = false
			return
		}
	}

	keyworldList := util.KeyworldListParse(getFeedKeyworldList())
	if len(keyworldList) <= 0 {
		retText = "无订阅词限定"
		retBool = true
		return
	} else {
		retBool = false
	}
	for _, v := range keyworldList {
		vc := v
		if strings.Contains(msgText, vc) {
			retText = vc
			retBool = true
			return
		}
	}
	return
}

// 文案级关键词替换
func feedKeyworldReplace(content string) string {
	if len(content) <= 0 {
		return ""
	}
	retText := content
	keyworldListMap := util.KeyworldListParseToMap(getFeedKeyworldReplace())
	for k, v := range keyworldListMap {
		if len(k) > 0 {
			retText = strings.ReplaceAll(retText, k, v)
		}
	}

	return retText
}
