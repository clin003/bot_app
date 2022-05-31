package core

import "syscall"

var (
	api = syscall.NewLazyDLL("coupler.dll")
	// 获取dll api
	initialize      = getApiProc("Api_Initialize")
	setFatal        = getApiProc("Api_SetFatal")
	getAppDirectory = getApiProc("Api_GetAppDirectory")
	getFrameVersion = getApiProc("Api_GetFrameVersion")
	outPutLog       = getApiProc("Api_AppendLogs")

	sendTextMsg  = getApiProc("Api_SendTextMsg")
	sendImageMsg = getApiProc("Api_SendImageMsg")
	sendVideoMsg = getApiProc("Api_SendVideoMsg")
	sendFileMsg  = getApiProc("Api_SendFileMsg")

	sendCardMsg       = getApiProc("Api_SendCardMsg")
	sendGroupMsgAndAt = getApiProc("Api_SendGroupMsgAndAt")
	sendEmojiMsg      = getApiProc("Api_SendEmojiMsg")
	sendLinkMsg       = getApiProc("Api_SendLinkMsg")
	sendMiniAppMsg    = getApiProc("Api_SendMiniAppMsg")

	sendMusicMsg = getApiProc("Api_SendMusicMsg")
	forwardMsg   = getApiProc("Api_ForwardMsg")

	getRobotName         = getApiProc("Api_GetRobotName")
	getLoggedAccountList = getApiProc("Api_GetLoggedAccountList")

	getFriendList            = getApiProc("Api_GetFriendList")
	getGroupList             = getApiProc("Api_GetGroupList")
	getGroupMemberDetailInfo = getApiProc("Api_GetGroupMemberDetailInfo")
	getGroupMemberList       = getApiProc("Api_GetGroupMemberList")

	getContactHeadimgurl = getApiProc("Api_GetContactHeadimgurl")
	acceptTransfer       = getApiProc("Api_AcceptTransfer")
	agreeGroupInvite     = getApiProc("Api_AgreeGroupInvite")
	agreeFriendVerify    = getApiProc("Api_AgreeFriendVerify")
	removeGroupMember    = getApiProc("Api_RemoveGroupMember")
	modifyGroupName      = getApiProc("Api_ModifyGroupName")
	modifyGroupNotice    = getApiProc("Api_ModifyGroupNotice")

	buildingGroupPlus = getApiProc("Api_BuildingGroupPlus")
	quitGroup         = getApiProc("Api_QuitGroup")
	inviteInGroup     = getApiProc("Api_InviteInGroup")
)
var authCode int

func getApiProc(name string) *syscall.LazyProc {
	return api.NewProc(name)
}

func Initialize(session int64, appJsonText string) int {
	r, _, _ := initialize.Call(int2ptr(session), str2ptr(appJsonText))
	authCode = ptr2int(r) //int64(ptr2int(r))

	return authCode
	// 	r, _, _ := getOnlineState.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	// 	return ptr2int(r)
}
func SetFatal() {
	_, _, _ = setFatal.Call(uintptr(authCode))
	return
}
func GetAppDirectory() string {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = getAppDirectory.Call(
		uintptr(authCode),
	)
	return ptr2str(r)
}
func GetFrameVersion() string {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = getFrameVersion.Call(
		uintptr(authCode),
	)
	return ptr2str(r)
}
func OutPutLog(msg string) {
	_, _, _ = outPutLog.Call(uintptr(authCode), str2ptr(msg))
}

func SendTextMsg(robotId, toId, msg string) {
	if authCode <= 0 {
		return
	}
	_, _, _ = sendTextMsg.Call(uintptr(authCode), str2ptr(robotId), str2ptr(toId), str2ptr(msg))
	// return ptr2int(r) //int64(ptr2int(r))
}
func SendImageMsg(robotId, toId, filePath string) {
	if authCode <= 0 {
		return
	}
	_, _, _ = sendImageMsg.Call(uintptr(authCode), str2ptr(robotId), str2ptr(toId), str2ptr(filePath))
	// return ptr2int(r) //int64(ptr2int(r))
}
func SendVideoMsg(robotId, toId, filePath string) {
	if authCode <= 0 {
		return
	}
	_, _, _ = sendVideoMsg.Call(uintptr(authCode), str2ptr(robotId), str2ptr(toId), str2ptr(filePath))
	// return ptr2int(r) //int64(ptr2int(r))
}
func SendFileMsg(robotId, toId, filePath string) {
	if authCode <= 0 {
		return
	}
	_, _, _ = sendFileMsg.Call(uintptr(authCode), str2ptr(robotId), str2ptr(toId), str2ptr(filePath))
	// return ptr2int(r) //int64(ptr2int(r))
}
func SendCardMsg(robotId, toId, friendId string) {
	if authCode <= 0 {
		return
	}
	_, _, _ = sendCardMsg.Call(uintptr(authCode), str2ptr(robotId), str2ptr(toId), str2ptr(friendId))
	// return ptr2int(r) //int64(ptr2int(r))
}
func SendGroupMsgAndAt(robotId, toGId, memberIds, memberNames, msg string) {
	if authCode <= 0 {
		return
	}
	_, _, _ = sendGroupMsgAndAt.Call(
		uintptr(authCode), str2ptr(robotId), str2ptr(toGId), str2ptr(memberIds), str2ptr(memberNames), str2ptr(msg),
	)
	// return ptr2int(r) //int64(ptr2int(r))
}
func SendEmojiMsg(robotId, toId, filePath string) {
	if authCode <= 0 {
		return
	}
	_, _, _ = sendEmojiMsg.Call(
		uintptr(authCode), str2ptr(robotId), str2ptr(toId), str2ptr(filePath),
	)
	// return ptr2int(r) //int64(ptr2int(r))
}
func SendLinkMsg(robotId, toId, title, text, targetUrl, picUrl, iconUrl string) {
	if authCode <= 0 {
		return
	}
	_, _, _ = sendLinkMsg.Call(
		uintptr(authCode), str2ptr(robotId),
		str2ptr(toId), str2ptr(title), str2ptr(text), str2ptr(targetUrl), str2ptr(picUrl), str2ptr(iconUrl),
	)
	// return ptr2int(r) //int64(ptr2int(r))
}
func SendMiniAppMsg(robotId, toId, xmlContent string) {
	if authCode <= 0 {
		return
	}
	_, _, _ = sendMiniAppMsg.Call(
		uintptr(authCode), str2ptr(robotId),
		str2ptr(toId), str2ptr(xmlContent),
	)
	// return ptr2int(r) //int64(ptr2int(r))
}
func SendMusicMsg(robotId, toId, name string, mType int64) {
	if authCode <= 0 {
		return
	}
	_, _, _ = sendMusicMsg.Call(
		uintptr(authCode), str2ptr(robotId),
		str2ptr(toId), str2ptr(name), int2ptr(mType),
	)
	// return ptr2int(r) //int64(ptr2int(r))
}
func ForwardMsg(robotId, toId, msgId string) {
	if authCode <= 0 {
		return
	}
	_, _, _ = forwardMsg.Call(
		uintptr(authCode), str2ptr(robotId),
		str2ptr(toId), str2ptr(msgId),
	)
	// return ptr2int(r) //int64(ptr2int(r))
}
func GetRobotName(robotId string) string {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = getRobotName.Call(
		uintptr(authCode), str2ptr(robotId),
		// str2ptr(toId), str2ptr(msgId),
	)
	return ptr2str(r) //int64(ptr2int(r))
}
func GetLoggedAccountList(robotId string) string {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = getLoggedAccountList.Call(
		uintptr(authCode),
		// str2ptr(robotId),
		// str2ptr(toId), str2ptr(msgId),
	)
	return ptr2str(r) //int64(ptr2int(r))
}
func GetFriendList(robotId string, isRefresh bool) string {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = getFriendList.Call(
		uintptr(authCode),
		str2ptr(robotId),
		bool2ptr(isRefresh),
		// str2ptr(toId), str2ptr(msgId),
	)
	return ptr2str(r) //int64(ptr2int(r))
}
func GetGroupList(robotId string, isRefresh bool) string {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = getGroupList.Call(
		uintptr(authCode),
		str2ptr(robotId),
		bool2ptr(isRefresh),
		// str2ptr(toId), str2ptr(msgId),
	)
	return ptr2str(r) //int64(ptr2int(r))
}
func GetGroupMemberDetailInfo(robotId, groupId, memberId string, isRefresh bool) string {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = getGroupMemberDetailInfo.Call(
		uintptr(authCode),
		str2ptr(robotId),
		str2ptr(groupId),
		str2ptr(memberId),
		bool2ptr(isRefresh),
		// str2ptr(toId), str2ptr(msgId),
	)
	return ptr2str(r) //int64(ptr2int(r))
}
func GetGroupMemberList(robotId, groupId string, isRefresh bool) string {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = getGroupMemberList.Call(
		uintptr(authCode),
		str2ptr(robotId),
		str2ptr(groupId),
		bool2ptr(isRefresh),
	)
	return ptr2str(r)
}
func GetContactHeadimgurl(robotId, toId string) string {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = getContactHeadimgurl.Call(
		uintptr(authCode),
		str2ptr(robotId),
		str2ptr(toId),
	)
	return ptr2str(r)
}
func AcceptTransfer(robotId, fromId, msgJson string) int {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = acceptTransfer.Call(
		uintptr(authCode),
		str2ptr(robotId),
		str2ptr(fromId),
		str2ptr(msgJson),
	)
	return ptr2int(r)
}
func AgreeGroupInvite(robotId, msgJson string) int {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = agreeGroupInvite.Call(
		uintptr(authCode),
		str2ptr(robotId),
		str2ptr(msgJson),
	)
	return ptr2int(r)
}
func AgreeFriendVerify(robotId, msgJson string) int {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = agreeFriendVerify.Call(
		uintptr(authCode),
		str2ptr(robotId),
		str2ptr(msgJson),
	)
	return ptr2int(r)
}
func RemoveGroupMember(robotId, groupId, memberId string) int {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = removeGroupMember.Call(
		uintptr(authCode),
		str2ptr(robotId),
		str2ptr(groupId),
		str2ptr(memberId),
	)
	return ptr2int(r)
}
func ModifyGroupName(robotId, groupId, groupName string) int {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = modifyGroupName.Call(
		uintptr(authCode),
		str2ptr(robotId),
		str2ptr(groupId),
		str2ptr(groupName),
	)
	return ptr2int(r)
}
func ModifyGroupNotice(robotId, groupId, content string) int {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = modifyGroupNotice.Call(
		uintptr(authCode),
		str2ptr(robotId),
		str2ptr(groupId),
		str2ptr(content),
	)
	return ptr2int(r)
}
func BuildingGroupPlus(robotId, memberIds string) int {
	// memberIds=id1|id2
	if authCode <= 0 {
		return ""
	}
	r, _, _ = buildingGroupPlus.Call(
		uintptr(authCode),
		str2ptr(robotId),
		str2ptr(memberIds),
	)
	return ptr2int(r)
}
func QuitGroup(robotId, groupId string) int {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = quitGroup.Call(
		uintptr(authCode),
		str2ptr(robotId),
		str2ptr(groupId),
	)
	return ptr2int(r)
}
func InviteInGroup(robotId, groupId, friendId string) int {
	if authCode <= 0 {
		return ""
	}
	r, _, _ = inviteInGroup.Call(
		uintptr(authCode),
		str2ptr(robotId),
		str2ptr(groupId),
		str2ptr(friendId),
	)
	return ptr2int(r)
}
