package core

import "C"

var AppInfo func() string
var AppInit func(path string) int64

var AppSetting func() int64
var AppStop func() int64
var AppEnable func() int64
var Login func(robotId, robotName string, eType int64, msg string) int64

// 群消息
var OnEventGroupMsg func(eventName, robotId string, msgType int64, fromWxId, fromName, finalFromWxId, finalFromName, toWxid, toName, msg, msgId, rawMsg string) int64

// 私聊消息
var OnEventFriendMsg func(eventName, robotId string, msgType int64, fromWxId, fromName, toWxid, toName, msg, msgId, rawMsg string) int64

// 本人发出的消息都在这里
var OnEventSendOutMsg func(eventName, robotId string, msgType, sendOutScene int64, toWxid, toName, msg, msgId, rawMsg string) int64

// 面对面收款
var OnEventScanCashMoney func(eventName, robotId, fromWxId, fromName, msg, rawMsg string) int64

// 收到转账事件
var OnEventReceivedTransfer func(eventName, robotId, fromWxId, fromName, toWxid, money, rawMsg string) int64

// 好友请求事件
var OnEventFriendVerify func(eventName, robotId, fromWxId, fromName, toWxid, rawMsg string) int64

//朋友变动事件
var OnEventContactsChange func(eventName, robotId string, msgType int64, fromWxId, fromName, msg string) int64

// 群成员增加事件
var OnEventGroupMemberAdd func(eventName, robotId, fromWxId, fromName, rawMsg string) int64

// 群成员减少
var OnEventGroupMemberDecrease func(eventName, robotId, fromWxId, fromName, rawMsg string) int64

// 系统消息事件
var OnEventSysMsg func(eventName, robotId string, msgType int64, rawMsg string) int64

//export LoadingInfo
func LoadingInfo(lApiString C.int) C.int {
	if AppInfo == nil {
		return C.int(0)
	}
	return C.int(Initialize(goInt(lApiString), AppInfo()))
}

//export EventInit
func EventInit(appPath *C.char) C.int {
	if AppInit == nil {
		return C.int(0)
	}
	SetFatal()
	return C.int(AppInit(goString(appPath)))
}

//export EventEnable
func EventEnable() C.int {
	if AppEnable == nil {
		return C.int(0)
	}
	return C.int(AppEnable())
}

//export EventStop
func EventStop(eType C.int) C.int {
	if AppStop == nil {
		return C.int(0)
	}
	return C.int(AppStop())
}

//export EventLogin
func EventLogin(robotId *C.char, robotName *C.char, eType C.int, msg *C.char) C.int {
	if Login == nil {
		return C.int(0)
	}
	return C.int(Login(goString(robotId), goString(robotName), goInt(eType), goString(msg)))
}

//export Menu
func Menu() {
	if AppSetting == nil {
		return
	}
	AppSetting()
}

//export EventGroupMsg
func EventGroupMsg(robotId *C.char, msgType C.int, fromWxId, fromName, finalFromWxId, finalFromName, toWxid, toName, msg, msgId, rawMsg *C.char) C.int {
	if OnEventGroupMsg == nil {
		return C.int(0)
	}
	return C.int(OnEventGroupMsg("EventGroupMsg",
		goString(robotId), goInt(msgType),
		goString(fromWxId), goString(fromName),
		goString(finalFromWxId), goString(finalFromName),
		goString(toWxid), goString(toName), goString(msg), goString(msgId),
		goString(rawMsg),
	))
}

//export EventFriendMsg
func EventFriendMsg(robotId *C.char, msgType C.int, fromWxId, fromName, toWxid, toName, msg, msgId, rawMsg *C.char) C.int {
	if OnEventGroupMsg == nil {
		return C.int(0)
	}
	return C.int(OnEventFriendMsg("EventFriendMsg",
		goString(robotId), goInt(msgType),
		goString(fromWxId), goString(fromName),
		// goString(finalFromWxId), goString(finalFromName),
		goString(toWxid), goString(toName), goString(msg), goString(msgId),
		goString(rawMsg),
	))
}

//export EventSendOutMsg
func EventSendOutMsg(robotId *C.char, msgType, sendOutScene C.int, toWxid, toName, msg, msgId, rawMsg *C.char) C.int {
	if OnEventGroupMsg == nil {

		return C.int(0)
	}
	return C.int(OnEventSendOutMsg("EventSendOutMsg",
		goString(robotId), goInt(msgType),
		goInt(sendOutScene),
		// goString(fromWxId), goString(fromName),
		// goString(finalFromWxId), goString(finalFromName),
		goString(toWxid), goString(toName), goString(msg), goString(msgId),
		goString(rawMsg),
	))
}

//export EventReceivedTransfer
func EventReceivedTransfer(robotId, fromWxId, fromName, toWxid, money, rawMsg *C.char) C.int {
	if OnEventGroupMsg == nil {

		return C.int(0)
	}
	return C.int(OnEventReceivedTransfer("EventReceivedTransfer",
		goString(robotId),
		goString(fromWxId), goString(fromName),
		goString(toWxid),
		goString(money),
		goString(rawMsg),
	))
}

//export EventScanCashMoney
func EventScanCashMoney(robotId, fromWxId, fromName, msg, rawMsg *C.char) C.int {
	if OnEventScanCashMoney == nil {

		return C.int(0)
	}
	return C.int(OnEventScanCashMoney("EventScanCashMoney",
		goString(robotId),
		goString(fromWxId), goString(fromName),
		goString(msg),
		goString(rawMsg),
	))
}

//export EventFriendVerify
func EventFriendVerify(robotId, fromWxId, fromName, toWxid, rawMsg *C.char) C.int {
	if OnEventFriendVerify == nil {

		return C.int(0)
	}
	return C.int(OnEventFriendVerify("EventFriendVerify",
		goString(robotId),
		goString(fromWxId), goString(fromName),
		goString(toWxid),
		goString(rawMsg),
	))
}

//export EventContactsChange
func EventContactsChange(robotId *C.char, msgType C.int, fromWxId, fromName, msg *C.char) C.int {
	if OnEventGroupMsg == nil {

		return C.int(0)
	}
	return C.int(OnEventContactsChange("EventContactsChange",
		goString(robotId), goInt(msgType),
		goString(fromWxId), goString(fromName),
		goString(msg),
	))
}

//export EventGroupMemberAdd
func EventGroupMemberAdd(robotId, fromWxId, fromName, rawMsg *C.char) C.int {
	if OnEventGroupMsg == nil {

		return C.int(0)
	}
	return C.int(OnEventGroupMemberAdd("EventGroupMemberAdd",
		goString(robotId),
		goString(fromWxId), goString(fromName),
		goString(rawMsg),
	))
}

//export EventGroupMemberDecrease
func EventGroupMemberDecrease(robotId, fromWxId, fromName, rawMsg *C.char) C.int {
	if OnEventGroupMsg == nil {

		return C.int(0)
	}
	return C.int(OnEventGroupMemberDecrease("EventGroupMemberDecrease",
		goString(robotId),
		goString(fromWxId), goString(fromName),
		goString(rawMsg),
	))
}

//export EventSysMsg
func EventSysMsg(robotId *C.char, msgType C.int, rawMsg *C.char) C.int {
	if OnEventGroupMsg == nil {
		return C.int(0)
	}
	return C.int(OnEventSysMsg("EventSysMsg",
		goString(robotId), goInt(msgType),
		goString(rawMsg),
	))
}
