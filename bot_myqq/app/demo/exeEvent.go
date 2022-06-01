package main

import (
	"bot_app/bot_myqq/core"
	"bot_app/bot_myqq/onebot"
	"fmt"
	"log"
)

func onPrivateMessage(xe onebot.XEvent) {
	core.OutPutLog(fmt.Sprintf("%s:onPrivateMessage(%d):%s", PluginName, xe.UserID, xe.Message))
	// log.Println("demo", xe.UserID, xe.Message)
	log.Println(xe)
}

func onGroupMessage(xe onebot.XEvent) {
	log.Println(xe)
}

// 群文件上传
func noticeFileUpload(xe onebot.XEvent) {
	log.Println(xe)
}

// 管理员变动
func noticeAdminChange(xe onebot.XEvent, typ string) {
	log.Println(xe)
}

// 群成员减少
func noticeGroupMenberDecrease(xe onebot.XEvent, typ string) {
	log.Println(xe)
}

// 群成员增加
func noticeGroupMenberIncrease(xe onebot.XEvent) {
	log.Println(xe)
}

// 群禁言
func noticeGroupBan(xe onebot.XEvent, typ string) {
	log.Println(xe)
}

// 好友添加
func noticeFriendAdd(xe onebot.XEvent) {
	log.Println(xe)
}

// 群消息撤回
func noticGroupMsgDelete(xe onebot.XEvent) {
	log.Println(xe)
}

// 好友消息撤回
func noticFriendMsgDelete(xe onebot.XEvent) {
	log.Println(xe)
}

// 加好友请求
func requestFriendAdd(xe onebot.XEvent) {
	log.Println(xe)
}

// 加群请求
func requestGroupAdd(xe onebot.XEvent, typ string) {
	log.Println(xe)
}
