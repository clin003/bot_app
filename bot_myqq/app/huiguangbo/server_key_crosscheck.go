package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"gitee.com/lyhuilin/util"
)

// {
//     "code": 0,
//     "message": "OK",
//     "data": true
// }
type ServerKeyCrosscheckResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    bool   `json:"data"`
}

// //是否已记录 true 真 存在，false 不存在，不存在存储
// func groupMsgKeyCrosscheck(groupId int64, msgText string) bool {
// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	time.Sleep(time.Millisecond * time.Duration(r.Intn(1000)))
// 	_, msgHash := util.GetTextCleaner(msgText)
// 	keyStr := fmt.Sprintf("%s_%d", msgHash, groupId)
// 	return ServerKeyCrosscheck(keyStr)
// }

// func ServerKeyCrosscheck(key string) (ret bool) {
// 	apiUrl := fmt.Sprintf("%s/apis/v1/SignatureCheckHandler", getServerURL())
// 	serverToken := getSenderWsserverToken()
// 	postData := map[string]string{
// 		"key":     key,
// 		"botoken": serverToken,
// 	}
// 	retData, err := util.PostToServer(apiUrl, serverToken, postData)
// 	if err != nil {
// 		return
// 	}
// 	var serverKeyCrosscheckResponse ServerKeyCrosscheckResponse

// 	if err := json.Unmarshal(retData, &serverKeyCrosscheckResponse); err != nil {
// 		return
// 	}
// 	return serverKeyCrosscheckResponse.Data
// }

//是否已记录 true 真 存在，false 不存在，不存在存储
func groupMsgKeyCrosscheckV2(groupId int64, msgText, msgId string) bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Millisecond * time.Duration(r.Intn(1000)))
	_, msgHash := util.GetTextCleaner(msgText)
	// keyStr := fmt.Sprintf("%s_%d", msgHash, groupId)
	return ServerKeyCrosscheckV2(string(groupId), msgHash, msgId)
}
func ServerKeyCrosscheckV2(keyBase, key, key2 string) (ret bool) {
	apiUrl := fmt.Sprintf("%s/apis/v2/SignatureCheck", getServerURL())
	serverToken := getSenderWsserverToken()
	postData := map[string]string{
		"key_base": keyBase,
		"key":      key,
		"key2":     key2,
		"botoken":  serverToken,
	}
	retData, err := util.PostToServer(apiUrl, serverToken, postData)
	if err != nil {
		return
	}
	var serverKeyCrosscheckResponse ServerKeyCrosscheckResponse

	if err := json.Unmarshal(retData, &serverKeyCrosscheckResponse); err != nil {
		return
	}
	return serverKeyCrosscheckResponse.Data
}
