package main

import (
	"bot_app/bot_keaimao/core"
	"bot_app/bot_keaimao/onebot"
	_ "embed"
	"encoding/json"

	"gitee.com/lyhuilin/util"
)

//go:embed icon.png
var iconBin []byte

func init() {
	core.AppInfo = AppInfo
	core.AppInit = AppInit
	core.AppSetting = AppSetting
	core.AppStop = AppStop
	core.AppEnable = AppEnable
	core.Login = Login

	core.OnEventGroupMsg = OnEventGroupMsg
	core.OnEventFriendMsg = OnEventFriendMsg
	core.OnEventSendOutMsg = OnEventSendOutMsg
	core.OnEventScanCashMoney = OnEventScanCashMoney
	core.OnEventReceivedTransfer = OnEventReceivedTransfer
	core.OnEventFriendVerify = OnEventFriendVerify
	core.OnEventContactsChange = OnEventContactsChange
	core.OnEventGroupMemberAdd = OnEventGroupMemberAdd
	core.OnEventGroupMemberDecrease = OnEventGroupMemberDecrease
	core.OnEventSysMsg = OnEventSysMsg
}

var (
	//插件名称
	PluginName = "huiguangbo"
	//插件版本
	PluginVer = "0.0.7"
	//插件作者
	PluginAuthor = "白菜林"
	//插件说明
	PluginDesc = "慧林淘友软件交流2群:153690156"
	// 文件名
	PluginFileName = "huiguangbo.cat.dll"
	// 以上4者可自行修改
	//插件Skey
	PluginSkey = "153690156"
	//插件SDK
	// PluginSDK             = "S1"
	// PluginMenuButtonTitle = "设置"
	//以上两个变量请勿修改
)

type PluginInfo struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"desc"`
	Version     string `json:"version"`
	// FileName    string `json:"文件名"`
	ApiVersion   string `json:"api_version"`
	DeveloperKey string `json:"developer_key"`
	MenuTitle    string `json:"menu_title"`
	CoverBase64  string `json:"cover_base64"`
}

// var iconBase64 string

func AppInfo() string {
	/*pluginInfo := fmt.Sprintf("插件名称{%s}\n插件版本{%s}\n插件作者{%s}\n插件说明{%s}\n插件skey{%s}\n插件sdk{%s}",
	PluginName, PluginVer, PluginAuthor, PluginDesc, PluginSkey, PluginSDK)*/
	iconBase64 := util.Base64StdEncodingEncodeToString(iconBin)
	pluginInfo := PluginInfo{
		Name:        PluginName,
		Author:      PluginAuthor,
		Description: PluginDesc,
		Version:     PluginVer,
		// FileName:     PluginFileName,
		ApiVersion:   "5.0",
		DeveloperKey: PluginSkey,
		MenuTitle:    "设置",
		CoverBase64:  iconBase64,
	}
	bytes, _ := json.Marshal(pluginInfo)
	return string(bytes)
}

func AppInit(appConfPath string) int64 {
	go onebot.ProtectRun(func() {
		core.OutPutLog("AppInit:" + appConfPath)
		// core.OutPutLog("iconBase64:" + iconBase64)
	}, "outPutLog")
	// go onebot.ProtectRun(func() {
	// 	once2.Do(func() {
	// 		Init()
	// 		InitHGBConf()
	// 	})
	// }, "Init()")
	return 0
}

func AppSetting() int64 {
	go onStart()
	go onebot.ProtectRun(func() {
		core.OutPutLog("AppSetting by 白菜林")
	}, "outPutLog")
	return 0
}
func AppStop() int64 {
	go onebot.ProtectRun(func() {
		core.OutPutLog("AppStop by 白菜林")
	}, "outPutLog")
	return 0
}

func AppEnable() int64 {
	go onebot.ProtectRun(func() {
		core.OutPutLog("AppEnable by 白菜林")
	}, "outPutLog")

	go onebot.ProtectRun(func() {
		once2.Do(func() {
			Init()
			InitHGBConf()
		})
	}, "Init()")
	return 0
}

func Login(robotId, robotName string, eType int64, msg string) int64 {
	go onebot.ProtectRun(func() {
		core.OutPutLog("Login:" + robotName + "(" + robotId + "):" + msg)
	}, "outPutLog")
	go onebot.ProtectRun(func() {
		once2.Do(func() {
			Init()
			InitHGBConf()
		})
	}, "Init()")
	return 0
}

// func OnEvent(eventName, robotId string, msgType, sendOutScene int64, fromWxId, fromName, finalFromWxId, finalFromName, toWxid, toName, msg, msgId, rawMsg string) int64 {
// 	go onebot.ProtectRun(func() {
// 		onebot.DEBUG("eventName:%s (%s)", eventName, robotId)
// 	}, eventName)
// 	return 0
// }

func OnEventGroupMsg(eventName, robotId string, msgType int64, fromWxId, fromName, finalFromWxId, finalFromName, toWxid, toName, msg, msgId, rawMsg string) int64 {

	go onebot.ProtectRun(func() {
		onebot.INFO("%s %s %s(%s):%s\r\n%s", eventName, robotId, fromName, fromWxId, msg, rawMsg)
		core.SendTextMsg(robotId, finalFromWxId, msg)
		onebot.DEBUG("sendmsg(%s,%s,%s)", robotId, finalFromWxId, msg)
	}, "OnEventGroupMsg")
	return 0
}

// 私聊消息
func OnEventFriendMsg(eventName, robotId string, msgType int64, fromWxId, fromName, toWxid, toName, msg, msgId, rawMsg string) int64 {

	return 0
}

// 本人发出的消息都在这里
func OnEventSendOutMsg(eventName, robotId string, msgType, sendOutScene int64, toWxid, toName, msg, msgId, rawMsg string) int64 {

	return 0
}

// 面对面收款
func OnEventScanCashMoney(eventName, robotId, fromWxId, fromName, msg, rawMsg string) int64 {

	return 0
}

// 收到转账事件
func OnEventReceivedTransfer(eventName, robotId, fromWxId, fromName, toWxid, money, rawMsg string) int64 {

	return 0
}

// 好友请求事件
func OnEventFriendVerify(eventName, robotId, fromWxId, fromName, toWxid, rawMsg string) int64 {

	return 0
}

//朋友变动事件
func OnEventContactsChange(eventName, robotId string, msgType int64, fromWxId, fromName, msg string) int64 {

	return 0
}

// 群成员增加事件
func OnEventGroupMemberAdd(eventName, robotId, fromWxId, fromName, rawMsg string) int64 {

	return 0
}

// 群成员减少
func OnEventGroupMemberDecrease(eventName, robotId, fromWxId, fromName, rawMsg string) int64 {

	return 0
}

// 系统消息事件
func OnEventSysMsg(eventName, robotId string, msgType int64, rawMsg string) int64 {

	return 0
}
