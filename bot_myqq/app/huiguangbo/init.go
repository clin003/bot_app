package main

import (
	"bot_app/bot_myqq/core"
	"bot_app/bot_myqq/onebot"
	"fmt"
	"math/rand"
	"path/filepath"
	"strings"
	"sync"
	"time"

	// "gitee.com/lyhuilin/open_api/model/feedmsg"
	"gitee.com/lyhuilin/model/feedmsg"
	"gitee.com/lyhuilin/util"

	"gopkg.in/yaml.v2"
)

var hgbConf = HGBConf{}
var once sync.Once
var once2 sync.Once
var isReady bool
var feedGroupList map[int64]bool

func getAppName() string {
	return PluginName
}
func getPathConf() (retText string) {
	// path := viper.GetString("module.huiguangbo.path")
	// path := "./conf/huiguangbo/huiguangbo.yaml"
	path := filepath.Join(util.Getwd(), "plugin", getAppName())
	onebot.INFO(path)
	if !util.IsExist(path) {
		util.MkDir(path)
	}
	path = filepath.Join(path, "config.yaml")

	return path
}
func Init() {
	feedGroupList = make(map[int64]bool, 0)
	path := getPathConf()

	bytes, err := util.ReadFile(path)
	if err != nil {
		// log.Errorf(err, "读取配置文件 %s 出错了", path)
		onebot.ERROR(fmt.Sprintf("读取配置文件 %s 出错了", path), err)
	}
	err = yaml.Unmarshal(bytes, &hgbConf)
	if err != nil {
		// log.Errorf(err, "加载配置文件 %s 出错了", path)
		onebot.ERROR(fmt.Sprintf("加载配置文件 %s 出错了", path), err)
	}
	// wsServerUrl := viper.GetString("module.huiguangbo.server_url")
	// channel := viper.GetString("module.huiguangbo.server_token")

	wsServerUrl := hgbConf.ServerURL
	channel := hgbConf.ServerToken
	// go wsClientStart(wsServerUrl, channel)
	if len(channel) > 0 {
		go initWsServer(wsServerUrl, channel, sendMsg)
	}
}

// 将richMsg消息转化为string
func richMsgToSendingMessage(richMsg feedmsg.FeedRichMsgModel) (retMsg string, err error) {
	m := ""
	if richMsg.Msgtype == "rich" {
		richMsgTextContent := richMsg.Text.Content
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
			m = fmt.Sprintf("%s\r\n%s", m, picMsg)
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
func sendMsg(richMsg feedmsg.FeedRichMsgModel) {
	if !isReady {
		return
	}
	onebot.DEBUG(fmt.Sprintf("收到广播消息，开始处理(%s)", richMsg.ToString()))

	// 处理(格式化)待发布消息
	sendGroupMsg := func(robotID, groupID int64) {
		groupCode := groupID
		msg, err := richMsgToSendingMessage(richMsg)
		if err != nil {
			// log.Errorf(err, "消息处理失败(%d): %s", groupCode, richMsg.ToString())
			onebot.ERROR(fmt.Sprintf("消息处理失败(%d): %s", groupCode, richMsg.ToString()), err)
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

	groupList := getAllFeedGroupInfoExList()
	isSended := make(map[int64]bool, 0)
	for _, v := range groupList {
		if _, ok := isSended[v.Id]; ok {
			continue
		}
		isSended[v.Id] = true

		if hgbConf.SenderSleep <= 100*time.Microsecond {
			vID := v.Id
			rID := v.RobotId
			go sendGroupMsg(rID, vID)

		} else {
			vID := v.Id
			rID := v.RobotId
			sendGroupMsg(rID, vID)

			time.Sleep(hgbConf.SenderSleep)
		}

	}
}
func isFeed(groupID int64) bool {
	if v, ok := feedGroupList[groupID]; !ok {
		return false
	} else {
		return v
	}
}

//乱序
func getAllFeedGroupInfoExList() []GroupInfoEx {
	list := make([]GroupInfoEx, 0)
	onlineList, _ := getOnlineList()
	for _, v := range onlineList {
		groupList, _ := getGroupList(v)
		for _, vv := range groupList {
			var groupInfoEx GroupInfoEx
			groupInfoEx.Id = vv
			groupInfoEx.RobotId = v
			groupInfoEx.Name = core.GetGroupName(v, vv)
			groupInfoEx.IsFeed = isFeed(vv)
			if groupInfoEx.IsFeed {
				list = append(list, groupInfoEx)
			}
		}
	}

	// 打乱数组顺序
	rand.Seed(time.Now().UnixNano())
	// retList := make([]GroupInfoEx, len(list))
	// copy(retList, list)
	rand.Shuffle(len(list), func(i int, j int) {
		list[i], list[j] = list[j], list[i]
	})
	return list
}
func getOnlineList() ([]int64, int) {
	return split(core.GetOnLineList())
}
func getGroupList(robotID int64) ([]int64, int) {
	return split(core.GetGroupListA(robotID))
}

func split(liststr string) ([]int64, int) {
	list := make([]int64, 0)
	listTmp := strings.Split(liststr, "\r\n")
	count := len(listTmp)
	for _, v := range listTmp {
		if len(v) > 0 {
			vd := core.Str2Int(v)
			list = append(list, vd)
		}
	}
	return list, count
}

func InitHGBConf() {
	for {
		if core.IsEnable() {
			break
		}
		time.Sleep(60 * time.Second)
	}
	time.Sleep(60 * time.Second)
	onebot.INFO("开始 加载慧广播配置信息")
	onlineList, _ := getOnlineList()
	for _, v := range onlineList {
		groupList, _ := getGroupList(v)
		for _, vv := range groupList {
			groupName := core.GetGroupName(v, vv)
			groupCode := vv
			if len(groupName) > 0 && groupCode > 0 {
				var groupInfo GroupInfo
				groupInfo.Id = groupCode
				groupInfo.Name = groupName
				isInConf := false
				for _, c := range hgbConf.GroupList {
					if c.Id == groupCode {
						isInConf = true
						break
					}
				}
				if !isInConf {
					hgbConf.GroupList = append(hgbConf.GroupList, groupInfo)
				}
			}
		}
	}

	for _, c := range hgbConf.GroupList {
		if c.IsFeed {
			feedGroupList[c.Id] = c.IsFeed
		}
	}

	if hgbConf.SenderSleep <= 0 {
		hgbConf.SenderSleep = 100 * time.Microsecond
	}

	if len(hgbConf.GroupList) > 0 {
		outBody, err := yaml.Marshal(hgbConf)
		if err != nil {
			onebot.ERROR(fmt.Sprintf("生成配置信息编码出错(yaml.Marshal):%v", err))
		} else {
			path := getPathConf()
			err := util.WriteFile(path, outBody)
			if err != nil {
				onebot.ERROR(fmt.Sprintf("写入配置信息到文件(%s)出错(WriteFile):%v", path, err))
			} else {
				onebot.INFO(fmt.Sprintf("写入配置信息到文件(%s):%s", path, string(outBody)))
			}

		}
	}
	onebot.INFO("完成 加载慧广播配置信息")
	once.Do(func() {
		isReady = true
	})

}
