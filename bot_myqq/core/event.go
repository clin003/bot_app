package core

import "C"

var Info func() string
var Event func(selfID int64, mseeageType int64, subType int64, groupID int64, userID int64, noticeID int64, message string, messageNum int64, messageID int64, rawMessage string, ret int64) int64
var MQSet func() int64
var MQEnd func() int64

//export MQ_Info
func MQ_Info() *C.char {
	return cString(Info())
}

//export MQ_Event
func MQ_Event(QQ *C.char, MsgType, SubMsgType C.int, MsgFrom, TigObjAct, TigObjPas, Msg, MsgNum, MsgId, RawMsg *C.char, PtrNext C.int) C.int {
	return C.int(Event(CStr2GoInt(QQ), goInt(MsgType), goInt(SubMsgType), CStr2GoInt(MsgFrom), CStr2GoInt(TigObjAct),
		CStr2GoInt(TigObjPas), UnescapeEmoji(goString(Msg)), CStr2GoInt(MsgNum), CStr2GoInt(MsgId), goString(RawMsg), goInt(PtrNext)))
}

//export MQ_End
func MQ_End() C.int {
	if MQEnd == nil {
		return C.int(0)
	}
	return C.int(MQEnd())
}

//export MQ_Set
func MQ_Set() {
	MQSet()
}
