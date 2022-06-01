package core

// 登录账号列表
type LoggedAccountList []struct {
	Backgroundimgurl string `json:"backgroundimgurl"`
	Headimgbase64    string `json:"headimgbase64"`
	Headimgurl       string `json:"headimgurl"`
	LoginTime        int64  `json:"login_time"`
	Nickname         string `json:"nickname"`
	Signature        string `json:"signature"`
	Status           int64  `json:"status"`
	WxNum            string `json:"wx_num"`
	WxPid            int64  `json:"wx_pid"`
	WxWindHandle     int64  `json:"wx_wind_handle"`
	Wxid             string `json:"wxid"` //机器人号
}

// 群聊列表
type GroupList []struct {
	RobotWxid string `json:"robot_wxid"`
	Nickname  string `json:"nickname"` //群名
	Wxid      string `json:"wxid"`     //群号
}

// 群号/机器人号列表
type WxidList []struct {
	Nickname string `json:"nickname"`
	Wxid     string `json:"wxid"` //群号
}
