package onebot

type XEvent struct {
	ID          int64  `db:"id"`           // ID 消息ID 唯一对应
	SelfID      int64  `db:"self_id"`      // SelfID 机器人QQ, 多Q版用于判定哪个QQ接收到该消息
	MessageType int64  `db:"message_type"` // MessageType 消息类型, 接收到消息类型，该类型可在常量表中查询具体定义，此处仅列举： -1 未定义事件 0,在线状态临时会话 1,好友信息 2,群信息 3,讨论组信息 4,群临时会话 5,讨论组临时会话 6,财付通转账 7,好友验证回复会话
	SubType     int64  `db:"sub_type"`     // SubType 消息子类型, 此参数在不同消息类型下，有不同的定义，暂定：接收财付通转账时 1为好友 4为群临时会话 5为讨论组临时会话    有人请求入群时，不良成员这里为1
	GroupID     int64  `db:"group_id"`     // GroupID 消息来源, 此消息的来源，如：群号、讨论组ID、临时会话QQ、好友QQ等
	UserID      int64  `db:"user_id"`      // UserID 触发对象_主动, 主动发送这条消息的QQ，踢人时为踢人管理员QQ
	NoticeID    int64  `db:"notice_id"`    // NoticeID 触发对象_被动, 被动触发的QQ，如某人被踢出群，则此参数为被踢出人QQ
	Message     string `db:"message"`      // Message 消息内容, 此参数有多重含义，常见为：对方发送的消息内容，但当消息类型为 某人申请入群，则为入群申请理由
	MessageNum  int64  `db:"message_num"`  // MessageNum 消息序号, 此参数暂定用于消息回复，消息撤回
	MessageID   int64  `db:"message_id"`   // MessageID 消息ID, 此参数暂定用于消息回复，消息撤回
	RawMessage  string `db:"raw_message"`  // RawMessage 原始信息, UDP收到的原始信息，特殊情况下会返回JSON结构（入群事件时，这里为该事件seq）
	Time        int64  `db:"time"`         // Time 消息时间戳, 接受到消息的时间戳
	Ret         int64  `db:"ret"`          // Ret 回传文本指针, 此参数用于插件加载拒绝理由
}
