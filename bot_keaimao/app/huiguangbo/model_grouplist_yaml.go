package main

import (
	"time"
)

type HGBConf struct {
	ServerURL   string `yaml:"huiguangbo_server_url"`
	ServerToken string `yaml:"huiguangbo_server_token"`

	GroupList   []GroupInfo   `yaml: "group_list"`
	SenderSleep time.Duration `yaml: "sender_sleep"`
	// GuildList         []GuildInfo   `yaml: "guild_list"`
	// FeedGuildList []string `yaml: "feed_guild_list"` //GuildID_ChannelID
	// GuildSenderEnable bool          `yaml: "guild_sender_enable"`

	// GroupListEx []GroupInfoEx `yaml: "group_list_ex"`
}

type GroupInfo struct {
	Name   string `yaml:"group_name"`
	Id     string `yaml:"group_id"`
	IsFeed bool   `yaml:"send_msg_enable"`
}

type GroupInfoEx struct {
	GroupInfo
	RobotId string `yaml:"robot_id"`
}

// type GuildInfo struct {
// 	Name        string        `yaml:"guild_name"`
// 	Id          uint64        `yaml:"guild_id"`
// 	ChannelList []ChannelInfo `yaml: "channel_list"`
// 	// Code int64  `yaml:"guild_code"`
// }
// type ChannelInfo struct {
// 	Name      string `yaml:"channel_name"`
// 	Id        uint64 `yaml:"channel_id"`
// 	EventTime uint32 `yaml:"channel_event_time"`
// 	IsFeed    bool   `yaml:"send_msg_enable"`
// }
