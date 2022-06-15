package main

import (
	"bot_app/bot_myqq/core"
	"bot_app/bot_myqq/onebot"
	"encoding/json"
)

func init() {
	core.Info = Info
	core.Event = Event
	core.MQSet = MQSet
	core.MQEnd = MQEnd
}

var (
	//插件名称
	PluginName = "huiguangbo"
	//插件版本
	PluginVer = "0.0.15"
	//插件作者
	PluginAuthor = "白菜林"
	//插件说明
	PluginDesc = "慧林淘友软件交流2群:153690156"
	// 以上4者可自行修改
	//插件Skey
	PluginSkey = "SDG5D4Ys89h7DJ849d"
	//插件SDK
	PluginSDK = "S1"
	//以上两个变量请勿修改
)

type PluginInfo struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Skey        string `json:"skey"`
	Sdk         string `json:"sdk"`
}

func Info() string {
	/*pluginInfo := fmt.Sprintf("插件名称{%s}\n插件版本{%s}\n插件作者{%s}\n插件说明{%s}\n插件skey{%s}\n插件sdk{%s}",
	PluginName, PluginVer, PluginAuthor, PluginDesc, PluginSkey, PluginSDK)*/
	pluginInfo := PluginInfo{
		Name:        PluginName,
		Author:      PluginAuthor,
		Description: PluginDesc,
		Version:     PluginVer,
		Skey:        PluginSkey,
		Sdk:         PluginSDK,
	}
	bytes, _ := json.Marshal(pluginInfo)
	return string(bytes)
}

// SelfID 机器人QQ, 多Q版用于判定哪个QQ接收到该消息
// MessageType 消息类型, 接收到消息类型，该类型可在常量表中查询具体定义，此处仅列举： -1 未定义事件 0,在线状态临时会话 1,好友信息 2,群信息 3,讨论组信息 4,群临时会话 5,讨论组临时会话 6,财付通转账 7,好友验证回复会话
// SubType 消息子类型, 此参数在不同消息类型下，有不同的定义，暂定：接收财付通转账时 1为好友 4为群临时会话 5为讨论组临时会话    有人请求入群时，不良成员这里为1
// GroupID 消息来源, 此消息的来源，如：群号、讨论组ID、临时会话QQ、好友QQ等
// UserID 触发对象_主动, 主动发送这条消息的QQ，踢人时为踢人管理员QQ
// NoticeID 触发对象_被动, 被动触发的QQ，如某人被踢出群，则此参数为被踢出人QQ
// Message 消息内容, 此参数有多重含义，常见为：对方发送的消息内容，但当消息类型为 某人申请入群，则为入群申请理由
// MessageNum 消息序号, 此参数暂定用于消息回复，消息撤回
// MessageID 消息ID, 此参数暂定用于消息回复，消息撤回
// RawMessage 原始信息, UDP收到的原始信息，特殊情况下会返回JSON结构（入群事件时，这里为该事件seq）
// Ret 回传文本指针, 此参数用于插件加载拒绝理由
func Event(selfID int64, mseeageType int64, subType int64, groupID int64, userID int64, noticeID int64, message string, messageNum int64, messageID int64, seq string, ret int64) int64 {
	xe := onebot.XEvent{
		ID:          0,
		SelfID:      selfID,
		MessageType: mseeageType,
		SubType:     subType,
		GroupID:     groupID,
		UserID:      userID,
		NoticeID:    noticeID,
		Message:     message,
		MessageNum:  messageNum,
		MessageID:   messageID,
		RawMessage:  seq,
		Ret:         ret,
	}

	switch mseeageType {
	/*case core.MT_P_ENABLE:
	go onebot.ProtectRun(func() { onStart() }, "onStart()")*/
	case core.MT_P_ENABLE, core.MT_P_LOAD:
		go onebot.ProtectRun(func() {
			once2.Do(func() {
				Init()
				InitHGBConf()
			})
		}, "Init()")
	case core.MT_QQ_LOGIN, core.MT_QQ_ADD, core.MT_QQ_OFFLINEACT, core.MT_QQ_OFFLINEPAS, core.MT_QQ_DROPLINE:
		go onebot.ProtectRun(func() { onLogin(xe) }, "onPrivateMessage()")
		// go onebot.ProtectRun(func() {
		// 	InitHGBConf()
		// }, "Init()")

	// 消息事件
	// 0：临时会话 1：好友会话 4：群临时会话 57：好友验证会话
	case core.MT_ONLINETEMP, core.MT_FRIEND, core.MT_GROUPTP, core.MT_DISGROUPTP, core.MT_FRIENDVERIFY:
		go onebot.ProtectRun(func() { onPrivateMessage(xe) }, "onPrivateMessage()")
	// 2：群聊信息
	case core.MT_GROUP, core.MT_DISGROUP:
		go onebot.ProtectRun(func() { onGroupMessage(xe) }, "onGroupMessage()")
	// 通知事件
	// 群文件接收
	case core.MT_G_GROUPFILE:
		go onebot.ProtectRun(func() { noticeFileUpload(xe) }, "noticeFileUpload()")
	// 管理员变动 210为有人升为管理 211为有人被取消管理
	case core.MT_G_SBBECOMEADMIN:
		go onebot.ProtectRun(func() { noticeAdminChange(xe, "set") }, "noticeAdminChange()")
	case core.MT_G_SBOUTGOINGADMIN:
		go onebot.ProtectRun(func() { noticeAdminChange(xe, "unset") }, "noticeAdminChange()")
	// 群成员减少 201为主动退群 202为被踢
	case core.MT_G_SBQUITGROUP:
		go onebot.ProtectRun(func() { noticeGroupMenberDecrease(xe, "leave") }, "OnGroupMenberDecrease()")
	case core.MT_G_ADMINkICKSB:
		go onebot.ProtectRun(func() { noticeGroupMenberDecrease(xe, "kick") }, "noticeGroupMenberDecrease()")
	// 群成员增加
	case core.MT_G_SBAPPROVALGROUP:
		go onebot.ProtectRun(func() { noticeGroupMenberIncrease(xe) }, "noticeGroupMenberIncrease()")
	// 群禁言 203为禁言 204为解禁
	case core.MT_G_SBISSHUTUP:
		go onebot.ProtectRun(func() { noticeGroupBan(xe, "ban") }, "noticeGroupBan()")
	case core.MT_G_SBREMOVESHUTUP:
		go onebot.ProtectRun(func() { noticeGroupBan(xe, "lift_ban") }, "noticeGroupBan()")
	// new
	// 好友添加 100 为单向
	case core.MT_F_SINGLE, core.MT_F_AGREEADDME:
		go onebot.ProtectRun(func() { noticeFriendAdd(xe) }, "noticeFriendAdd()")
	// 群消息撤回 subType 2
	// 好友消息撤回 subType 1
	case 9:
		if xe.SubType == 2 {
			go onebot.ProtectRun(func() { noticGroupMsgDelete(xe) }, "noticGroupMsgDelete()")
		} else {
			go onebot.ProtectRun(func() { noticFriendMsgDelete(xe) }, "noticFriendMsgDelete()")
		}
	// 请求事件
	// 加好友请求
	case core.MT_F_SBADDME:
		go onebot.ProtectRun(func() { requestFriendAdd(xe) }, "requestFriendAdd()")
	// 加群请求／邀请 213为请求 214为被邀
	case core.MT_G_SBWANTTOJOINGROUP:
		go onebot.ProtectRun(func() { requestGroupAdd(xe, "add") }, "requestGroupAdd()")
	case core.MT_G_SBINVITATIONMEJOINGROUP:
		go onebot.ProtectRun(func() { requestGroupAdd(xe, "invite") }, "requestGroupAdd()")
	default:
		//
	}
	return 0
}

func MQEnd() int64 {
	return 0
}

func MQSet() int64 {
	go onStart()
	return 0
}
