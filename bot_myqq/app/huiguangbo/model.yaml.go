package main

import (
	"sync"
	"time"
)

// 慧转发/广播配置模型
type HZFConf struct {
	sync.RWMutex `yaml:"-"`

	SendGroupEnable     bool   `yaml:"send_group_enable"`
	ServerURL           string `yaml:"openapi_server_url"`
	ServerToken         string `yaml:"openapi_server_token"`  //频道标识
	SenderWsserverToken string `yaml:"sender_wsserver_token"` //发送验证token

	FeedKeyworldList    string `yaml:"feed_keyword_list"`    //订阅关键词列表 格式 key|key2
	FeedKeyworldFilter  string `yaml:"feed_keyword_filter"`  //订阅屏蔽/过滤关键词列表 格式 key|key2
	FeedKeyworldReplace string `yaml:"feed_keyword_replace"` //订阅文案关键词替换，格式 key>>key2|key3
	DedupEnable         bool   `yaml:"dedup_enable"`         //去重/重复数据删除 开关

	SenderSleep time.Duration `yaml:"sender_sleep"`
	GroupList   []GroupInfo   `yaml:"group_list"`
}

type GroupInfo struct {
	Name   string `yaml:"group_name"`
	Id     int64  `yaml:"group_id"`
	IsFeed bool   `yaml:"feed_msg_enable"`
	IsTo   bool   `yaml:"to_msg_enable"`
}

type GroupInfoEx struct {
	GroupInfo
	RobotId int64 `yaml:"robot_id"`
}

var Conf *HZFConf

func init() {
	Conf = new(HZFConf)
}
