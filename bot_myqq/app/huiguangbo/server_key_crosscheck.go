package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"gitee.com/lyhuilin/util"
)

//true 存在
func groupMsgKeyCrosscheck(groupId int64, msgText string) bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Millisecond * time.Duration(r.Intn(1000)))
	_, msgHash := util.GetTextCleaner(msgText)
	keyStr := fmt.Sprintf("%s_%d", msgHash, groupId)
	return ServerKeyCrosscheck(keyStr)
}

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

func ServerKeyCrosscheck(key string) (ret bool) {
	apiUrl := fmt.Sprintf("%s/apis/v1/SignatureCheckHandler", getServerURL())
	serverToken := getSenderWsserverToken()
	postData := map[string]string{
		"key":     key,
		"botoken": serverToken,
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
