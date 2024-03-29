package core

import "syscall"

var (
	api = syscall.NewLazyDLL("MyQQApi.dll")
	// 获取dll api
	sendMsg               = getApiProc("Api_SendMsg")
	sendMsgEx             = getApiProc("Api_SendMsgEx")
	sendXML               = getApiProc("Api_SendXML")
	sendJson              = getApiProc("Api_SendJSON")
	upVote                = getApiProc("Api_UpVote")
	getCookies            = getApiProc("Api_GetCookies")
	getBlogPsKey          = getApiProc("Api_GetBlogPsKey")
	getZonePsKey          = getApiProc("Api_GetZonePsKey")
	getGroupPsKey         = getApiProc("Api_GetGroupPsKey")
	getClassRoomPsKey     = getApiProc("Api_GetClassRoomPsKey")
	getBkn                = getApiProc("Api_GetBkn")
	getBkn32              = getApiProc("Api_GetBkn32")
	getLongLdw            = getApiProc("Api_GetLongLdw")
	getClientKey          = getApiProc("Api_GetClientkey")
	getLongClientKey      = getApiProc("Api_GetLongClientkey")
	adminInviteGroup      = getApiProc("Api_AdminInviteGroup")
	noAdminInviteGroup    = getApiProc("Api_NoAdminInviteGroup")
	getNick               = getApiProc("Api_GetNick")
	getGroupCard          = getApiProc("Api_GetGroupCard")
	getObjLevel           = getApiProc("Api_GetQQLevel")
	getFriendList         = getApiProc("Api_GetFriendList")
	getFriendListB        = getApiProc("Api_GetFriendList_B")
	getGroupAdmin         = getApiProc("Api_GetAdminList")
	getGroupListA         = getApiProc("Api_GetGroupList_A")
	getGroupList          = getApiProc("Api_GetGroupList")
	getGroupListB         = getApiProc("Api_GetGroupList_B")
	getGroupMemberList    = getApiProc("Api_GetGroupMemberList")
	getGroupMemberListB   = getApiProc("Api_GetGroupMemberList_B")
	getGroupMemberListC   = getApiProc("Api_GetGroupMemberList_C")
	isShutUp              = getApiProc("Api_IsShutUp")
	shutUp                = getApiProc("Api_ShutUP")
	joinGroup             = getApiProc("Api_JoinGroup")
	quitGroup             = getApiProc("Api_QuitGroup")
	uploadPic             = getApiProc("Api_UpLoadPic")
	getPicLink            = getApiProc("Api_GetPicLink")
	outPutLog             = getApiProc("Api_OutPut")
	teaEncry              = getApiProc("Api_TeaEncrypt")
	teaDecry              = getApiProc("Api_TeaDecrypt")
	sessionKey            = getApiProc("Api_SessionKey")
	gn2Gid                = getApiProc("Api_GNTransGID")
	gid2Gn                = getApiProc("Api_GIDTransGN")
	pbGroupNotic          = getApiProc("Api_PBGroupNotic")
	getNotice             = getApiProc("Api_GetNotice")
	shakeWindow           = getApiProc("Api_ShakeWindow")
	handleFriendEvent     = getApiProc("Api_HandleFriendEvent")
	handleGroupEvent      = getApiProc("Api_HandleGroupEvent")
	setAnon               = getApiProc("Api_SetAnon")
	getAnon               = getApiProc("Api_GetAnon")
	setGroupCard          = getApiProc("Api_SetGroupCard")
	quitDisGroup          = getApiProc("Api_QuitDisGroup")
	createDisGroup        = getApiProc("Api_CreateDisGroup")
	kickDisGroupMBR       = getApiProc("Api_KickDisGroupMBR")
	inviteDisGroup        = getApiProc("Api_InviteDisGroup")
	getDisGroupList       = getApiProc("Api_GetDisGroupList")
	getDisGroupMemberList = getApiProc("Api_GetDisGroupMemberList")
	setDisGroupName       = getApiProc("Api_SetDisGroupName")
	kickGroupMBR          = getApiProc("Api_KickGroupMBR")
	getObjVote            = getApiProc("Api_GetObjVote")
	uploadVoice           = getApiProc("Api_UpLoadVoice")
	getVoiceLink          = getApiProc("Api_GetVoiLink")
	getTimeStamp          = getApiProc("Api_GetTimeStamp")
	sendPack              = getApiProc("Api_SendPack")
	getObjInfo            = getApiProc("Api_GetObjInfo")
	getGender             = getApiProc("Api_GetGender")
	getQQAge              = getApiProc("Api_GetQQAge")
	getAge                = getApiProc("Api_GetAge")
	getPerExp             = getApiProc("Api_GetPerExp")
	getSign               = getApiProc("Api_GetSign")
	getEmail              = getApiProc("Api_GetEmail")
	getGroupName          = getApiProc("Api_GetGroupName")
	getVer                = getApiProc("Api_GetVer")
	getQQList             = getApiProc("Api_GetQQList")
	getOnLineList         = getApiProc("Api_GetOnlineQQlist")
	getOffLineList        = getApiProc("Api_GetOffLineList")
	addQQ                 = getApiProc("Api_AddQQ")
	delQQ                 = getApiProc("Api_DelQQ")
	loginQQ               = getApiProc("Api_Login")
	offLineQQ             = getApiProc("Api_Logout")
	ifFriend              = getApiProc("Api_SetStatus")
	setRInf               = getApiProc("Api_SetRInf")
	getRInf               = getApiProc("Api_GetStatus")
	addFriend             = getApiProc("Api_AddFriend")
	delFriend             = getApiProc("Api_DelFriend")
	addBkList             = getApiProc("Api_AddBkList")
	delBkList             = getApiProc("Api_DelBkList")
	setShieldedGroup      = getApiProc("Api_SetShieldedGroup")
	sendVoice             = getApiProc("Api_SendVoice")
	setAdmin              = getApiProc("Api_SetAdmin")
	pbHomeWork            = getApiProc("Api_PBHomeWork")
	isEnble               = getApiProc("Api_IsEnable")
	disabledPlugin        = getApiProc("Api_DisabledPlugin")
	withdrawMsg           = getApiProc("Api_WithdrawMsg")
	beInput               = getApiProc("Api_BeInput")
	getQQAddMode          = getApiProc("Api_GetQQAddMode")
	isOnline              = getApiProc("Api_IsOnline")
	getOnlineState        = getApiProc("Api_GetOnlineState")
	getGroupMemberNum     = getApiProc("Api_GetGroupMemberNum")
	getWpa                = getApiProc("Api_GetWpa")
	getGroupAddMode       = getApiProc("Api_GetGroupAddMode")
	getGroupLv            = getApiProc("Api_GetGroupChatLv")
	setFriendsRemark      = getApiProc("Api_SetFriendsRemark")
	getFriendsRemark      = getApiProc("Api_GetFriendsRemark")
	signIn                = getApiProc("Api_SignIn")
)

func getApiProc(name string) *syscall.LazyProc {
	return api.NewProc(name)
}

//发送消息
//@param QQ 		机器人qq
//@param msgType 	消息类型 0在线临时会话 1好友 2群 3讨论组 4群临时会话 5讨论组临时会话 7好友验证回复会话（0、7只支持Pro版）
//@param objGroup 	发送群信息、讨论组、群或讨论组临时会话信息时填写，如发送对象为好友或信息类型是0时可空
//@param objQQ 		收信对象QQ
//@param msg 		内容
//@param bubId 		气泡 -1为随机气泡
func SendMsg(QQ int64, msgType int64, objGroup, objQQ int64, msg string) {
	_, _, _ = sendMsg.Call(str2ptr(Int2Str(QQ)), int2ptr(msgType), str2ptr(Int2Str(objGroup)),
		str2ptr(Int2Str(objQQ)), str2ptr(EscapeEmoji(msg)))
}

//发送消息
//@param QQ 		机器人qq
//@parm  sendType	发送类型 1 普通发送 2 匿名发送
//@param msgType 	消息类型 0在线临时会话 1好友 2群 3讨论组 4群临时会话 5讨论组临时会话 7好友验证回复会话（0、7只支持Pro版）
//@param objGroup 	发送群信息、讨论组、群或讨论组临时会话信息时填写，如发送对象为好友或信息类型是0时可空
//@param objQQ 		收信对象QQ
//@param msg 		内容
//@param bubId 		气泡 -1为随机气泡
func SendMsgEX(QQ int64, sendType, msgType int64, objGroup, objQQ int64, msg string, bubId int64) error {
	_, _, err := sendMsgEx.Call(str2ptr(Int2Str(QQ)), int2ptr(sendType), int2ptr(msgType), str2ptr(Int2Str(objGroup)),
		str2ptr(Int2Str(objQQ)), str2ptr(msg), int2ptr(bubId))

	return err
}

//发送XML消息
//@param QQ 		机器人qq
//@param sendType 	发送方式 1普通 2匿名（匿名需要群开启）
//@param msgType 	消息类型 0在线临时会话 1好友 2群 3讨论组 4群临时会话 5讨论组临时会话 7好友验证回复会话（0、7只支持Pro版）
//@param objGroup 	群号 发送群信息、讨论组、群或讨论组临时会话信息时填写，如发送对象为好友或信息类型是0时可空
//@param objQQ 		收信对象QQ
//@param objMsg 	XML内容
//@param subType 	结构子类型 0 基本 2 点歌
func SendXML(QQ int64, sendType, msgType int64, objGroup, objQQ int64, objMsg string, subType int64) {
	_, _, _ = sendXML.Call(str2ptr(Int2Str(QQ)), int2ptr(sendType), int2ptr(msgType),
		str2ptr(Int2Str(objGroup)), str2ptr(Int2Str(objQQ)), str2ptr(objMsg), int2ptr(subType))
}

//发送JSON结构消息
func SendJSON(QQ int64, sendType, msgType int64, objGroup, objQQ int64, objJson string) {
	_, _, _ = sendJson.Call(str2ptr(Int2Str(QQ)), int2ptr(sendType), int2ptr(msgType),
		str2ptr(Int2Str(objGroup)), str2ptr(Int2Str(objQQ)), str2ptr(objJson))
}

//点赞
//调用一次点一下，成功返回空，失败返回理由如：每天最多给他点十个赞哦，调用此Api时应注意频率，每人每日可被赞10次，每日每Q最多可赞50人
//@param QQ	机器人qq
//@param objQQ 	被赞QQ
func UpVote(QQ, objQQ int64) string {
	r, _, _ := upVote.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2str(r)
}

//取得机器人网页操作用的Cookies
//@param QQ	机器人qq
func GetCookies(QQ int64) string {
	r, _, _ := getCookies.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//取得腾讯微博页面操作用参数P_skey
func GetBlogPsKey(QQ int64) string {
	r, _, _ := getBlogPsKey.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//取得QQ空间页面操作用参数P_skey
func GetZonePsKey(QQ int64) string {
	r, _, _ := getZonePsKey.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//取得QQ群页面操作用参数P_skey
func GetGroupPsKey(QQ int64) string {
	r, _, _ := getGroupPsKey.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//取得腾讯课堂页面操作用参数P_skey
func GetClassRoomPsKey(QQ int64) string {
	r, _, _ := getClassRoomPsKey.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//取得机器人网页操作用参数Bkn或G_tk
func GetBkn(QQ int64) string {
	r, _, _ := getBkn.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//取得机器人网页操作用参数长Bkn或长G_tk
func GetBkn32(QQ int64) string {
	r, _, _ := getBkn32.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//取得机器人网页操作用参数长Ldw
func GetLongLdw(QQ int64) string {
	r, _, _ := getLongLdw.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//取得机器人网页操作用的Clientkey
func GetClientKey(QQ int64) string {
	r, _, _ := getClientKey.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//取得机器人网页操作用的长Clientkey
func GetLongClientKey(QQ int64) string {
	r, _, _ := getLongClientKey.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//管理员邀请对象入群，频率过快会失败
func AdminInviteGroup(QQ, objQQ, objGroup int64) {
	_, _, _ = adminInviteGroup.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)), str2ptr(Int2Str(objGroup)))
}

//非管理员邀请对象入群，频率过快会失败
func NoAdminInviteGroup(QQ, objQQ, objGroup int64) {
	_, _, _ = noAdminInviteGroup.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)), str2ptr(Int2Str(objGroup)))
}

//取对象昵称
//测试通过
func GetNick(QQ, objQQ int64) string {
	r, _, _ := getNick.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2str(r)
}

//取对象群名片
//测试通过
func GetGroupCard(QQ, objGroup, objQQ int64) string {
	r, _, _ := getGroupCard.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)), str2ptr(Int2Str(objQQ)))
	return ptr2str(r)
}

//取对象QQ等级 成功返回等级  失败返回-1
//测试通过
func GetObjLevel(QQ, objQQ int64) int {
	r, _, _ := getObjLevel.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return int(r)
}

//取得好友列表，返回获取到的原始JSON格式信息，需自行解析
//测试通过
func GetFriendList(QQ int64) string {
	r, _, _ := getFriendList.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//取得好友列表，返回内容#换行符分割
//测试通过
func GetFriendListB(QQ int64) string {
	r, _, _ := getFriendListB.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//取得群管理，返回获取到的原始JSON格式信息，需自行解析
//测试通过
func GetGroupAdmin(QQ, objGroup int64) string {
	r, _, _ := getGroupAdmin.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)))
	return ptr2str(r)
}

//取得群列表，#换行符分割 不受最高获取500群限制（如需获取群名称请对应调用 Api_GetGroupName 理论群名获取不会频繁）
//测试通过
func GetGroupListA(QQ int64) string {
	r, _, _ := getGroupListA.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//取得群列表，返回获取到的原始JSON格式信息，需自行解析
//测试通过
func GetGroupList(QQ int64) string {
	r, _, _ := getGroupList.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//取得群列表，返回获取到的原始JSON格式信息，需自行解析
//测试通过
func GetGroupListB(QQ int64) string {
	r, _, _ := getGroupListB.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//取得群成员列表，返回获取到的原始JSON格式信息，需自行解析（有群员昵称）
//测试通过
func GetGroupMemberList(QQ, objGroup int64) string {
	r, _, _ := getGroupMemberList.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)))
	return ptr2str(r)
}

//取得群成员列表，返回QQ号和身份Json格式信息 失败返回空（无群员昵称）
//测试通过
func GetGroupMemberListB(QQ, objGroup int64) string {
	r, _, _ := getGroupMemberListB.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)))
	return ptr2str(r)
}

//取得群成员列表，返回获取到的原始JSON格式信息，需自行解析（有群员昵称）
//测试通过
func GetGroupMemberListC(QQ, objGroup int64) string {
	r, _, _ := getGroupMemberListC.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)))
	return ptr2str(r)
}

//查询对象或自己是否被禁言  禁言返回真 失败或未禁言返回假
//测试通过
func IsShutUp(QQ, objGroup, objQQ int64) bool {
	r, _, _ := isShutUp.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)), str2ptr(Int2Str(objQQ)))
	return ptr2bool(r)
}

//群内禁言某人
func ShutUP(QQ, objGroup, objQQ int64, shut int64) {
	_, _, _ = shutUp.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)), str2ptr(Int2Str(objQQ)), int2ptr(shut))
}

//申请加群
func JoinGroup(QQ, objGroup int64, reason string) {
	_, _, _ = joinGroup.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)), str2ptr(reason))
}

//退群
func QuitGroup(QQ, objGroup int64) {
	_, _, _ = quitGroup.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)))
}

//上传图片（通过读入字节集方式），可使用网页链接或本地读入，成功返回该图片GUID，失败返回空
func IRUploadPic(QQ int64, uploadType int64, obj int64, picData []byte) string {
	r, _, _ := uploadPic.Call(str2ptr(Int2Str(QQ)), int2ptr(uploadType), str2ptr(Int2Str(obj)), byte2ptr(picData))
	return ptr2str(r)
}

//根据图片GUID取得图片下载连接
func GetPicLink(QQ int64, picType int64, obj int64, guid string) string {
	r, _, _ := getPicLink.Call(str2ptr(Int2Str(QQ)), int2ptr(picType), str2ptr(Int2Str(obj)), str2ptr(guid))
	return ptr2str(r)
}

//向CleverQQ日志窗口发送一条本插件的日志，可用于调试输出需要的内容，或定位插件错误与运行状态
func OutPutLog(msg string) {
	_, _, _ = outPutLog.Call(str2ptr(msg))
}

//腾讯Tea加密
func TeaEncry(msg, key string) string {
	r, _, _ := teaEncry.Call(str2ptr(msg), str2ptr(key))
	return ptr2str(r)
}

//腾讯Tea解密
func TeaDecry(msg, key string) string {
	r, _, _ := teaDecry.Call(str2ptr(msg), str2ptr(key))
	return ptr2str(r)
}

//获取会话SessionKey密钥
func SessionKey(QQ int64) string {
	r, _, _ := sessionKey.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//群号转群ID
func GNTransGID(objGroup int64) string {
	r, _, _ := gn2Gid.Call(str2ptr(Int2Str(objGroup)))
	return ptr2str(r)
}

//群ID转群号
func GIDTransGN(gid int64) string {
	r, _, _ := gid2Gn.Call(str2ptr(Int2Str(gid)))
	return ptr2str(r)
}

//发布群公告（成功返回真，失败返回假），调用此API应保证响应QQ为管理员
func PBGroupNotice(QQ, objGroup int64, noticeTitle, noticeContent string) bool {
	r, _, _ := pbGroupNotic.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)), str2ptr(noticeTitle), str2ptr(noticeContent))
	return ptr2bool(r)
}

//取群公告，返回该群最新公告
func GetNotice(QQ, objGroup int64) string {
	r, _, _ := getNotice.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)))
	return ptr2str(r)
}

//向好友发起窗口抖动，调用此Api腾讯会限制频率
func ShakeWindow(QQ, ObjQQ int64) bool {
	r, _, _ := shakeWindow.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(ObjQQ)))
	return ptr2bool(r)
}

//某人请求添加好友验证处理
//.版本 2
//
//.参数 QQ, 文本型, , 机器人QQ
//.参数 objQQ, 文本型, , 请求添加好友人QQ
//.参数 handleType, 整数型, , 10同意 20拒绝 30忽略
//.参数 addMsg, 文本型, , 拒绝添加好友 附加信息
func HandleFriendEvent(QQ, objQQ int64, handleType int64, addMsg string) {
	_, _, _ = handleFriendEvent.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)), int2ptr(handleType), str2ptr(addMsg))
}

//某人请求入群 被邀请入群 等 群验证处理
//.版本 2
//
//.参数 QQ, 文本型, , 机器人QQ
//.参数 reqType, 整数型, , 213请求入群  214我被邀请加入某群  215某人被邀请加入群
//.参数 objQQ, 文本型, , 申请入群、被邀请人的QQ （当请求类型为214时这里为邀请人QQ）
//.参数 objGroup, 文本型, , 收到请求群号
//.参数 seq, 文本型, , 需要处理事件的seq ，这个参数在收到群事件时，IRC_原始信息会传递
//.参数 handleType, 整数型, , 10同意 20拒绝 30忽略
//.参数 addMsg, 文本型, , 拒绝入群附加信息
func HandleGroupEvent(QQ int64, reqType int64, objQQ, objGroup, seq int64, handleType int64, addMsg string) {
	_, _, _ = handleGroupEvent.Call(str2ptr(Int2Str(QQ)), int2ptr(reqType),
		str2ptr(Int2Str(objQQ)), str2ptr(Int2Str(objGroup)), str2ptr(Int2Str(seq)), int2ptr(handleType), str2ptr(addMsg))
}

//开关群匿名消息发送功能    成功返回true  失败返回false
func SetAnon(QQ, objGroup int64, _switch bool) bool {
	r, _, _ := setAnon.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)), bool2ptr(_switch))
	return ptr2bool(r)
}

//开关群匿名消息发送功能    成功返回true  失败返回false
func GetAnon(QQ, objGroup int64) bool {
	r, _, _ := getAnon.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)))
	return ptr2bool(r)
}

//修改对象群名片 成功返回true 失败返回false
func SetGroupCard(QQ, objGroup, objQQ int64, card string) bool {
	r, _, _ := setGroupCard.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)), str2ptr(Int2Str(objQQ)), str2ptr(card))
	return ptr2bool(r)
}

//退出讨论组
func QuitDisGroup(QQ, disId int64) {
	_, _, _ = quitDisGroup.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(disId)))
}

//创建一个讨论组 成功返回讨论组ID 失败返回空
func CreateDisGroup(QQ int64, disName string) string {
	r, _, _ := createDisGroup.Call(str2ptr(Int2Str(QQ)), str2ptr(disName))
	return ptr2str(r)
}

//将对象移除讨论组 成功返回空 失败返回理由
func KickDisGroupMBR(QQ, disId, objQQ int64) string {
	r, _, _ := kickDisGroupMBR.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(disId)), str2ptr(Int2Str(objQQ)))
	return ptr2str(r)
}

//邀请对象加入讨论组 成功返回空 失败返回理由
func InviteDisGroup(QQ, disId, objQQ int64) string {
	r, _, _ := inviteDisGroup.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(disId)), str2ptr(Int2Str(objQQ)))
	return ptr2str(r)
}

//取出讨论组列表 （返回格式为#换行符分割开的）
func GetDisGroupList(QQ int64) string {
	r, _, _ := getDisGroupList.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//取出讨论组成员列表 （返回格式为 换行符分割开的）
func GetDisGroupMemberList(QQ, disId int64) string {
	r, _, _ := getDisGroupMemberList.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(disId)))
	return ptr2str(r)
}

//修改讨论组名称
func SetDisGroupName(QQ, disId int64, disName string) {
	_, _, _ = setDisGroupName.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(disId)), str2ptr(disName))
}

//将对象移除群
func KickGroupMBR(QQ, objGroup, objQQ int64, noMore bool) {
	_, _, _ = kickGroupMBR.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)), str2ptr(Int2Str(objQQ)), bool2ptr(noMore))
}

//获取对象当前赞数量 失败返回-1
func GetObjVote(QQ, objQQ int64) int32 {
	r, _, _ := getObjVote.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return int32(r)
}

//上传QQ语音，成功返回语音GUID  失败返回空
func UploadVoice(QQ int64, uploadType int64, objGroup int64, voice []byte) string {
	r, _, _ := uploadVoice.Call(str2ptr(Int2Str(QQ)), int2ptr(uploadType), str2ptr(Int2Str(objGroup)), byte2ptr(voice))
	return ptr2str(r)
}

//通过语音GUID获取下载连接 失败返回空
func GetVoiLink(QQ int64, voiceGuid string) string {
	r, _, _ := getVoiceLink.Call(str2ptr(Int2Str(QQ)), str2ptr(voiceGuid))
	return ptr2str(r)
}

//获取当前框架内部时间戳
func IRGetTimeStamp() int32 {
	r, _, _ := getTimeStamp.Call()
	return int32(r)
}

//向腾讯发送原始封包（成功返回腾讯返回的包 失败返回空）
func SendPack(QQ int64, pack string) string {
	r, _, _ := sendPack.Call(str2ptr(Int2Str(QQ)), str2ptr(pack))
	return ptr2str(r)
}

//获取对象资料 此方式为http，调用时应自行注意控制频率（成功返回JSON格式自行解析）
func GetObjInfo(QQ, objQQ int64) string {
	r, _, _ := getObjInfo.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2str(r)
}

//取对象性别 1男 2女  未知或失败返回-1
func GetGender(QQ, objQQ int64) int {
	r, _, _ := getGender.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2int(r)
}

//取Q龄 成功返回Q龄 失败返回-1
func GetQQAge(QQ, objQQ int64) int {
	r, _, _ := getQQAge.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2int(r)
}

//取年龄 成功返回年龄 失败返回-1
func GetAge(QQ, objQQ int64) int {
	r, _, _ := getAge.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2int(r)
}

//取个人说明
func GetPerExp(QQ, objQQ int64) string {
	r, _, _ := getPerExp.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2str(r)
}

//取个人签名
func GetSign(QQ, objQQ int64) string {
	r, _, _ := getSign.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2str(r)
}

//取邮箱，获取对象QQ资料内邮箱栏为邮箱时返回
func GetEmail(QQ, objQQ int64) string {
	r, _, _ := getEmail.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2str(r)
}

//取QQ群名
func GetGroupName(QQ, objGroup int64) string {
	r, _, _ := getGroupName.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)))
	return ptr2str(r)
}

//取框架版本号
func GetVer() string {
	r, _, _ := getVer.Call()
	return ptr2str(r)
}

//取框架所有QQ号 换行符分割
func GetQQList() string {
	r, _, _ := getQQList.Call()
	return ptr2str(r)
}

//取框架在线QQ号 换行符分割
func GetOnLineList() string {
	r, _, _ := getOnLineList.Call()
	return ptr2str(r)
}

//取框架离线QQ号 （Pro版可用）换行符分割
func GetOffLineList() string {
	r, _, _ := getOffLineList.Call()
	return ptr2str(r)
}

//向框架帐号列表增加一个登录QQ 成功失败均返回理由（Pro版可用）
func AddQQ(qq, password int64, autoLogin bool) string {
	r, _, _ := addQQ.Call(str2ptr(Int2Str(qq)), str2ptr(Int2Str(password)), bool2ptr(autoLogin))
	return ptr2str(r)
}

//删除框架帐号列表内指定QQ，不可在执行登录过程中删除QQ否则有几率引起错误（Pro版可用）
func DelQQ(qq int64) string {
	r, _, _ := delQQ.Call(str2ptr(Int2Str(qq)))
	return ptr2str(r)
}

//登录指定QQ，应确保QQ号码在列表中已存在
func LoginQQ(qq int64) {
	_, _, _ = loginQQ.Call(str2ptr(Int2Str(qq)))
}

//令指定QQ下线，应确保QQ号码已在列表中且在线
func OffLineQQ(qq int64) {
	_, _, _ = offLineQQ.Call(str2ptr(Int2Str(qq)))
}

//是否QQ好友 好友返回真 非好友或获取失败返回假
func IfFriend(QQ, objQQ int64) bool {
	r, _, _ := ifFriend.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2bool(r)
}

//修改机器人在线状态 昵称 个性签名 性别
func SetRInf(QQ int64, qqType int64, msg string) {
	_, _, _ = setRInf.Call(str2ptr(Int2Str(QQ)), int2ptr(qqType), str2ptr(msg))
}

//获取机器人状态信息，成功返回：昵称、帐号、在线状态、速度、收信、发信、在线时间，失败返回空
func GetRInf(QQ int64) string {
	r, _, _ := getRInf.Call(str2ptr(Int2Str(QQ)))
	return ptr2str(r)
}

//删除好友 成功返回真 失败返回假
func AddFriend(QQ, objQQ int64, meg string) bool {
	r, _, _ := addFriend.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)), str2ptr(meg))
	return ptr2bool(r)
}

//删除好友 成功返回真 失败返回假
func DelFriend(QQ, objQQ int64) bool {
	r, _, _ := delFriend.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2bool(r)
}

//将好友拉入黑名单  成功返回真 失败返回假
func AddBkList(QQ, objQQ int64) bool {
	r, _, _ := addBkList.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2bool(r)
}

//解除好友黑名单
func DelBkList(QQ, objQQ int64) {
	_, _, _ = delBkList.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
}

//屏蔽或接收某群消息
func SetShieldedGroup(QQ, objGroup int64, _type bool) {
	_, _, _ = setShieldedGroup.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)), bool2ptr(_type))
}

//好友语音上传并发送 （成功返回真 失败返回假）
func SendVoice(QQ, objQQ int64, voice []byte) bool {
	r, _, _ := sendVoice.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)), byte2ptr(voice))
	return ptr2bool(r)
}

//设置或取消群管理员   成功返回空  失败返回理由
func SetAdmin(QQ, objGroup, objQQ int64, op bool) string {
	r, _, _ := setAdmin.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)), str2ptr(Int2Str(objQQ)), bool2ptr(op))
	return ptr2str(r)
}

//QQ群作业发布
func PBHomeWork(QQ, objQQ int64, workName, workTitle, workContent string) {
	_, _, _ = pbHomeWork.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)), str2ptr(workName), str2ptr(workTitle), str2ptr(workContent))
}

//取得插件自身启用状态，启用真 禁用假
func IsEnable() bool {
	r, _, _ := isEnble.Call()
	return ptr2bool(r)
}

//请求禁用插件自身
func DisabledPlugin() {
	_, _, _ = disabledPlugin.Call()
}

//消息撤回（成功返回空，失败返回腾讯给出的理由）（Pro版可用）
func WithdrawMsg(QQ, objGroup, msgNum, msgId int64) string {
	r, _, _ := withdrawMsg.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)), str2ptr(Int2Str(msgNum)), str2ptr(Int2Str(msgId)))
	return ptr2str(r)
}

//置正在输入状态（发送消息后会打断状态）
func BeInput(QQ, objQQ int64) {
	_, _, _ = beInput.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
}

//取对象好友添加验证方式 （00 允许任何人  01 需要身份验证  03 需回答正确问题  04 需回答问题  99 已经是好友） （Pro版可用）
func GetQQAddMode(QQ, objQQ int64) string {
	r, _, _ := getQQAddMode.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2str(r)
}

//查询对象是否在线
func IsOnline(QQ, objQQ int64) bool {
	r, _, _ := isOnline.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2bool(r)
}

//查询对象在线状态   返回 1、在线 2、Q我 3、离开 4、忙碌 5、勿扰 6、隐身或离线（Pro可用）
func GetOnlineState(QQ, objQQ int64) int {
	r, _, _ := getOnlineState.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2int(r)
}

//查询对象群当前人数和上限人数
func GetGroupMemberNum(QQ, objGroup int64) string {
	r, _, _ := getGroupMemberNum.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)))
	return ptr2str(r)
}

//查询对方是否允许在线状态临时会话消息（非讨论组和群临时）（Pro版可用）
func GetWpa(QQ, objQQ int64) bool {
	r, _, _ := getWpa.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2bool(r)
}

//查询对象群验证方式 1允许任何人 2需要验证消息 3不允许任何人加群 4需要正确回答问题 5需要回答问题并由管理员审核 6付费群 -1群号不存在（获取失败返回空）Pro版可用
func GetGroupAddMode(QQ, objGroup int64) string {
	r, _, _ := getGroupAddMode.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)))
	return ptr2str(r)
}

//查询QQ群等级，成功返回等级（失败返回-1）Pro版可用
func GetGroupLv(QQ, objGroup int64) int {
	r, _, _ := getGroupLv.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)))
	return ptr2int(r)
}

//修改好友备注姓名
func SetFriendsRemark(QQ, objQQ int64, remark string) {
	_, _, _ = setFriendsRemark.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)), str2ptr(remark))
}

//取好友备注姓名（成功返回备注，失败或无备注返回空）Pro可用
func GetFriendsRemark(QQ, objQQ int64) string {
	r, _, _ := getFriendsRemark.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objQQ)))
	return ptr2str(r)
}

//QQ群签到（成功返回真 失败返回假）
func SignIn(QQ, objGroup int64, location, want2say string) bool {
	r, _, _ := signIn.Call(str2ptr(Int2Str(QQ)), str2ptr(Int2Str(objGroup)), str2ptr(location), str2ptr(want2say))
	return ptr2bool(r)
}
