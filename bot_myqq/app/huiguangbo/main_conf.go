package main

import (
	"bot_app/bot_myqq/core"
	"bot_app/bot_myqq/onebot"
	"fmt"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"gitee.com/lyhuilin/util"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func getAppName() string {
	return PluginName
}

func getPathConf() (retText string) {
	path := filepath.Join(util.Getwd(), "plugin", getAppName())
	if !util.IsExist(path) {
		util.MkDir(path)
	}
	return filepath.Join(path, "huiguangbo.yaml")
}

func Init_Conf() {
	path := getPathConf()
	bytes, err := util.ReadFile(path)
	if err != nil {
		onebot.ERROR(fmt.Sprintf("读取配置文件 %s 出错了", path), err)
		return
	}
	Conf.Lock()
	defer Conf.Unlock()
	err = yaml.Unmarshal(bytes, Conf)
	if err != nil {
		onebot.ERROR(fmt.Sprintf("加载配置文件 %s 出错了", path), err)
	}
}

func isFeed(groupID int64) bool {
	return viper.GetBool(fmt.Sprintf("isfeed_%d", groupID))
}
func isTo(groupID int64) bool {
	return viper.GetBool(fmt.Sprintf("isto_%d", groupID))
}
func isBot(userID int64) bool {
	return viper.GetBool(fmt.Sprintf("isbot_%d", userID))
}

func getServerToken() string {
	return viper.GetString("openapi_server_token")
}
func getServerURL() string {
	return viper.GetString("openapi_server_url")
}
func getSenderWsserverToken() string {
	return viper.GetString("sender_wsserver_token")
}
func isSendGroupEnable() bool {
	return viper.GetBool("send_group_enable")
}
func getFeedKeyworldList() string {
	return viper.GetString("feed_keyword_list")
}
func getFeedKeyworldFilter() string {
	return viper.GetString("feed_keyword_filter")
}
func getFeedKeyworldReplace() string {
	return viper.GetString("feed_keyword_replace")
}
func isDedupEnable() bool {
	return viper.GetBool("dedup_enable")
}
func getSenderSleep() time.Duration {
	return viper.GetDuration("sender_sleep")
}

// // 乱序
//
//	func getAllFeedGroupInfoExList() []GroupInfoEx {
//		list := make([]GroupInfoEx, 0)
//		onlineList, _ := getOnlineList()
//		for _, v := range onlineList {
//			groupList, _ := getGroupList(v)
//			for _, vv := range groupList {
//				var groupInfoEx GroupInfoEx
//				groupInfoEx.Id = vv
//				groupInfoEx.RobotId = v
//				groupInfoEx.Name = core.GetGroupName(v, vv)
//				// groupInfoEx.IsFeed = isFeed(vv)
//				groupInfoEx.IsTo = isTo(vv)
//				if groupInfoEx.IsTo {
//					list = append(list, groupInfoEx)
//				}
//			}
//		}
//			// 打乱数组顺序
//			rand.Seed(time.Now().UnixNano())
//			rand.Shuffle(len(list), func(i int, j int) {
//				list[i], list[j] = list[j], list[i]
//			})
//			return list
//		}
func getAllFeedGroupInfoExListEx() map[string]GroupInfoEx {
	list := make(map[string]GroupInfoEx, 0)
	onlineList, _ := getOnlineList()
	for _, v := range onlineList {
		groupList, _ := getGroupList(v)
		for _, vv := range groupList {
			var groupInfoEx GroupInfoEx
			groupInfoEx.Id = vv
			groupInfoEx.RobotId = v
			// groupInfoEx.Name = core.GetGroupName(v, vv)
			// groupInfoEx.IsFeed = isFeed(vv)
			groupInfoEx.IsTo = isTo(vv)
			if groupInfoEx.IsTo {
				kk := fmt.Sprintf("%d_%d", v, vv)
				list[kk] = groupInfoEx
			}
		}
	}
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

var onceInitWsServer sync.Once

func Init_Conf_Save() {
	for {
		if core.IsEnable() {
			break
		}
		time.Sleep(60 * time.Second)
	}
	Conf.Lock()
	defer Conf.Unlock()
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
				for _, c := range Conf.GroupList {
					if c.Id == groupCode {
						isInConf = true
						break
					}
				}
				if !isInConf {
					Conf.GroupList = append(Conf.GroupList, groupInfo)
				}
			}
		}
	}

	for _, c := range Conf.GroupList {
		if c.IsFeed {
			viper.Set(fmt.Sprintf("isfeed_%d", c.Id), c.IsFeed)
		}
		if c.IsTo {
			viper.Set(fmt.Sprintf("isto_%d", c.Id), c.IsTo)
		}
	}

	if Conf.SenderSleep <= 0 {
		Conf.SenderSleep = 100 * time.Microsecond
	}
	viper.Set("sender_sleep", Conf.SenderSleep)
	viper.Set("dedup_enable", Conf.DedupEnable)
	viper.Set("feed_keyword_list", Conf.FeedKeyworldList)
	viper.Set("feed_keyword_filter", Conf.FeedKeyworldFilter)
	viper.Set("feed_keyword_replace", Conf.FeedKeyworldReplace)
	viper.Set("send_group_enable", Conf.SendGroupEnable)
	viper.Set("openapi_server_url", Conf.ServerURL)
	viper.Set("openapi_server_token", Conf.ServerToken)
	viper.Set("sender_wsserver_token", Conf.SenderWsserverToken)

	if len(Conf.GroupList) > 0 {
		outBody, err := yaml.Marshal(Conf)
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
	viper.Set("me_is_ready", true)
}
func isReady() bool {
	return viper.GetBool("me_is_ready")
}
