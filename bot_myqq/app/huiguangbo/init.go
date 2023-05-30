package main

func Init() {
	Init_Conf()
	Init_Conf_Save()
	InitWsServer()
}
func InitWsServer() {
	initWsServerEx(getServerURL(), getServerToken(), sendRichMsgToGroupListQQ)
}
