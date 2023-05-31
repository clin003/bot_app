package main

import (
	"bot_app/bot_myqq/core"
	"bot_app/bot_myqq/onebot"
	"fmt"
	"strings"
	"time"

	"gitee.com/lyhuilin/model/feedmsg"
	"gitee.com/lyhuilin/util"
)

// 将richMsg消息转化为string
func richMsgToSendingMessage(richMsg feedmsg.FeedRichMsgModel) (retMsg string, err error) {
	m := ""
	if richMsg.Msgtype == "rich" {
		richMsgTextContent := richMsg.Text.Content
		if k, f, ok := feedKeyworldCheck(richMsgTextContent); !ok {
			errText := fmt.Sprintf("订阅关键词检查(屏蔽词:%s),文案:%s", f, richMsgTextContent)
			err = fmt.Errorf(errText)
			onebot.DEBUG(errText)
			return
		} else {
			if len(k) > 0 {
				onebot.DEBUG(fmt.Sprintf("订阅关键词检查(订阅词:%s): %s", k, richMsgTextContent))
			}
		}
		if isDedupEnable() && util.MsgSignatureCheckEx(richMsg.MsgID, richMsgTextContent, 50) {
			errText := fmt.Sprintf("去重过滤 (%s)", richMsg.ToString())
			err = fmt.Errorf(errText)
			onebot.DEBUG(errText)
			return
		}
		richMsgTextContent = feedKeyworldReplace(richMsgTextContent)
		if strings.Contains(richMsgTextContent, "@全体成员") {
			richMsgTextContent = strings.ReplaceAll(richMsgTextContent, "@全体成员", "")
			m = "[@all]"
		}
		if len(richMsgTextContent) > 0 {
			if len(m) > 0 {
				m = m + " " + richMsgTextContent
			} else {
				m = richMsgTextContent
			}
			// m = m + " " + richMsgTextContent
		}

		if len(richMsg.Image.PicURL) > 0 &&
			strings.HasPrefix(richMsg.Image.PicURL, "http") &&
			!strings.Contains(richMsg.Image.PicURL, "gchat.qpic.cn") {
			picMsg := fmt.Sprintf("[pic=%s]", richMsg.Image.PicURL)

			if len(m) > 0 {
				// m = fmt.Sprintf("%s\r\n%s", m, picMsg)
				m = fmt.Sprintf("%s  %s", m, picMsg)
			} else {
				m = picMsg
			}
		}
	} else if richMsg.Msgtype == "text" {
		richMsgTextContent := richMsg.Text.Content
		if strings.Contains(richMsgTextContent, "@全体成员") {
			richMsgTextContent = strings.ReplaceAll(richMsgTextContent, "@全体成员", "")
			m = "[@all]"
		}
		if len(richMsgTextContent) > 0 {
			m = m + " " + richMsgTextContent
		}
	}

	if len(m) > 0 {
		retMsg = m
		return retMsg, nil
	}
	if err == nil {
		err = fmt.Errorf("no msg(空消息):%v", richMsg)
	}
	return "", err
}
func sendRichMsgToGroupListQQ(richMsg feedmsg.FeedRichMsgModel) {
	if !isReady() || !isSendGroupEnable() {
		return
	}
	onebot.DEBUG(fmt.Sprintf("收到广播消息，开始处理(%s)", richMsg.ToString()))

	// 处理(格式化)待发布消息
	sendGroupMsg := func(robotID, groupID int64) {
		groupCode := groupID
		msg, err := richMsgToSendingMessage(richMsg)
		if err != nil {
			onebot.ERROR(fmt.Sprintf("消息处理异常(%d): %s", groupCode, richMsg.ToString()), err)
			return
		}
		if len(msg) <= 0 {
			return
		}

		if !groupMsgKeyCrosscheck(groupCode, msg) {
			onebot.DEBUG(fmt.Sprintf("消息忽略(其他机器人已处理)(%d): %s", groupCode, richMsg.ToString()))
			return
		}
		// 广播消息
		sendResult := core.SendMsgEX(robotID, 0, 2, groupCode, 0, msg, 0) //robot.SendGroupMessage(groupCode, msg)
		if sendResult != nil {
			onebot.INFO(fmt.Sprintf("群(%d) 广播模式 已启用,发送消息 (ID: %d InternalId: %d ) ", groupCode, robotID, sendResult))
		} else {
			onebot.INFO(fmt.Sprintf("群(%d) 广播模式 已启用,发送消息 失败 :%s", groupCode, richMsg.ToString()))
		}
	}

	groupList := getAllFeedGroupInfoExListEx()
	isSended := make(map[int64]bool, 0)
	for _, v := range groupList {
		if _, ok := isSended[v.Id]; ok {
			continue
		}
		isSended[v.Id] = true

		vID := v.Id
		rID := v.RobotId
		sendGroupMsg(rID, vID)
		time.Sleep(getSenderSleep())
	}
}
