package main

import (
	"bot_app/bot_keaimao/core"
	"bot_app/bot_keaimao/onebot"

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
var feedGroupList map[string]bool
var wxidList map[string]string

func getAppName() string {
	return PluginName
}
func getAppDirName() string {
	return PluginName + ".cat"
}

// 获取 app 配置文件路径
func getPathConf() (retText string) {
	// path := viper.GetString("module.huiguangbo.path")
	// path := "./conf/huiguangbo/huiguangbo.yaml"
	appConfPath := filepath.Join(util.Getwd(), "app", getAppDirName())
	onebot.INFO(appConfPath)
	if !util.IsExist(appConfPath) {
		util.MkDir(appConfPath)
	}
	appConfFilePath := filepath.Join(appConfPath, "config.yaml")
	return appConfFilePath
}
func Init() {
	feedGroupList = make(map[string]bool, 0)
	wxidList = make(map[string]string, 0)
	confFilePath := getPathConf()

	bytes, err := util.ReadFile(confFilePath)
	if err != nil {
		// log.Errorf(err, "读取配置文件 %s 出错了", path)
		onebot.ERROR(fmt.Sprintf("读取配置文件 %s 出错了", confFilePath), err)
	}
	err = yaml.Unmarshal(bytes, &hgbConf)
	if err != nil {
		// log.Errorf(err, "加载配置文件 %s 出错了", path)
		onebot.ERROR(fmt.Sprintf("加载配置文件 %s 出错了", confFilePath), err)
	}
	// wsServerUrl := viper.GetString("module.huiguangbo.server_url")
	// channel := viper.GetString("module.huiguangbo.server_token")

	wsServerUrl := hgbConf.ServerURL
	channel := hgbConf.ServerToken

	if len(channel) > 0 {
		go initWsServer(wsServerUrl, channel, sendMsg)
	}
	// go wsClientStart(wsServerUrl, channel)
}

// 将richMsg消息转化为string
func richMsgToSendingMessage(richMsg feedmsg.FeedRichMsgModel) (retMsg string, err error) {
	m := ""
	if richMsg.Msgtype == "rich" {
		richMsgTextContent := richMsg.Text.Content
		if strings.Contains(richMsgTextContent, "@全体成员") {
			richMsgTextContent = strings.ReplaceAll(richMsgTextContent, "@全体成员", "")
			m = "[@所有人]"
		}
		if len(richMsgTextContent) > 0 {
			richMsgTextContent = strings.ReplaceAll(richMsgTextContent, "\r", "")
			m = m + " " + richMsgTextContent
		}

		// if len(richMsg.Image.PicURL) > 0 && strings.HasPrefix(richMsg.Image.PicURL, "http") {
		// 	picMsg := fmt.Sprintf("[pic=%s]", richMsg.Image.PicURL)
		// 	m = fmt.Sprintf("%s\r\n%s", m, picMsg)
		// }
	} else if richMsg.Msgtype == "text" {
		richMsgTextContent := richMsg.Text.Content
		if strings.Contains(richMsgTextContent, "@全体成员") {
			richMsgTextContent = strings.ReplaceAll(richMsgTextContent, "@全体成员", "")
			m = "[@所有人]"
		}
		if len(richMsgTextContent) > 0 {
			richMsgTextContent = strings.ReplaceAll(richMsgTextContent, "\r", "")
			if len(m) > 0 {
				m = m + " " + richMsgTextContent
			} else {
				m = richMsgTextContent
			}

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

// 将richMsg消息转化为string
func richMsgToSendingMessageFilepath(richMsg feedmsg.FeedRichMsgModel) (retMsg string, err error) {
	m := ""
	if richMsg.Msgtype == "rich" {
		if len(richMsg.Image.PicURL) > 0 &&
			strings.HasPrefix(richMsg.Image.PicURL, "http") &&
			!strings.Contains(richMsg.Image.PicURL, "gchat.qpic.cn") {
			// picMsg := fmt.Sprintf("[pic=%s]", richMsg.Image.PicURL)
			// m = fmt.Sprintf("%s\r\n%s", m, picMsg)

			appDataDirPath := filepath.Join(util.Getwd(), "data", "image")
			if !util.IsExist(appDataDirPath) {
				util.MkDir(appDataDirPath)
			}

			picURLHash := util.EncryptMd5(richMsg.Image.PicURL)
			picFilepath := filepath.Join(appDataDirPath, picURLHash)
			onebot.INFO(picFilepath)
			if util.IsExist(picFilepath) {
				return picFilepath, nil
			}
			if _, err1 := util.GetUrlWithCookieToStringAndSaveFile(richMsg.Image.PicURL, "", "", "", "", picFilepath); err1 != nil {
				return "", err1
			} else {
				return picFilepath, nil
			}
			// m = picFilepath
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
	sendGroupMsg := func(robotID, groupID string) {
		groupCode := groupID

		if msg, err := richMsgToSendingMessage(richMsg); err != nil {
			// onebot.ERROR(fmt.Sprintf("消息处理失败(%d): %s", groupCode, richMsg.ToString()), err)
		} else if len(msg) > 0 {
			// 广播消息
			core.SendTextMsg(robotID, groupCode, msg)
		}
		if msg, err := richMsgToSendingMessageFilepath(richMsg); err != nil {
			// onebot.ERROR(fmt.Sprintf("消息处理失败(%d): %s", groupCode, richMsg.ToString()), err)
		} else if len(msg) > 0 {
			// 广播消息
			core.SendImageMsg(robotID, groupCode, msg)
		}

	}

	groupList := getAllFeedGroupInfoExList()
	isSended := make(map[string]bool, 0)
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
func isFeed(groupID string) bool {
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
			groupInfoEx.Name = wxidList[vv] //core.GetGroupName(v, vv)
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

func getOnlineList() ([]string, int) {
	return split(core.GetLoggedAccountList())
}
func getGroupList(robotID string) ([]string, int) {
	return split(core.GetGroupList(robotID, false))
}

func split(liststr string) ([]string, int) {
	list := make([]string, 0)
	var listTmp core.WxidList
	err := util.JsonDecode(liststr, &listTmp)
	if err != nil {

	}
	count := len(listTmp)
	for _, v := range listTmp {
		if len(v.Wxid) > 0 {
			list = append(list, v.Wxid)
			wxidList[v.Wxid] = v.Nickname
		}
	}
	return list, count
}

func InitHGBConf() {
	for {
		if list, count := getOnlineList(); len(list) > 0 || count > 0 {
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
			groupName := wxidList[vv] // core.GetGroupName(v, vv)
			groupCode := vv
			if len(groupName) > 0 && len(groupCode) > 0 {
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
