package core

// .版本 2

// .程序集 可爱猫应用, , 公开
// .程序集变量 appInfo, 应用信息
// .程序集变量 apiAddr, 整数型
// .程序集变量 authCode, 整数型

// .子程序 _初始化, , , 当基于本类的对象被创建后，此方法会被自动调用

// .子程序 _销毁, , , 当基于本类的对象被销毁前，此方法会被自动调用

// .子程序 初始化, 整数型, 公开
// .参数 info, 应用信息
// .参数 session, 整数型
// .局部变量 json, 类_json
// .局部变量 addr, 整数型
// .局部变量 objName, 文本型
// .局部变量 content, 文本型
// .局部变量 bin, 字节集

// .如果真 (info.ApiVersion ＝ “”)
//     info.ApiVersion ＝ “5.0”
// .如果真结束

// json.置属性 (“name”, info.Name)
// json.置属性 (“desc”, info.Desc)
// json.置属性 (“author”, info.Author)
// json.置属性 (“version”, info.Version)
// json.置属性 (“api_version”, info.ApiVersion)
// json.置属性 (“developer_key”, info.AuthorKey)
// json.置属性 (“menu_title”, info.MenuButtonTitle)
// json.置属性 (“cover_base64”, 编码_BASE64编码 (info.Img))

// content ＝ json.取数据文本 ()

// authCode ＝ Api_Initialize (session, content)

// 返回 (authCode)

// .子程序 发送文本消息, 整数型, 公开, 发送文字消息（支持好友/群）
// .参数 robot_wxid, 文本型, , 用哪个机器人发这条消息
// .参数 to_wxid, 文本型, , 对方的ID（支持好友/群ID）
// .参数 msg, 文本型, , 发送的内容

// 返回 (Api_SendTextMsg (authCode, robot_wxid, to_wxid, msg))

// .子程序 发送文本消息w, 整数型, 公开, 发送字节集形式的文字消息（支持好友/群） * 限测版 *
// .参数 robot_wxid, 文本型, , 用哪个机器人发这条消息
// .参数 to_wxid, 文本型, , 对方的ID（支持好友/群ID）
// .参数 byte, 字节集, , 发送的字节集形式的内容，当前编码为utf8

// 返回 (Api_SendTextMsgByByteHex (authCode, robot_wxid, to_wxid, 字节集_字节集到十六进制 (byte)))

// .子程序 发送图片消息, 整数型, 公开, 发送图片（支持好友/群）
// .参数 robot_wxid, 文本型, , 用哪个机器人发这条消息
// .参数 to_wxid, 文本型, , 对方的ID（支持好友/群ID）
// .参数 path, 文本型, , 图片文件的绝对路径

// 返回 (Api_SendImageMsg (authCode, robot_wxid, to_wxid, path))

// .子程序 发送视频消息, 整数型, 公开, 发送视频消息（支持好友/群）
// .参数 robot_wxid, 文本型, , 用哪个机器人发这条消息
// .参数 to_wxid, 文本型, , 对方的ID（支持好友/群ID）
// .参数 path, 文本型, , 视频存放的绝对路径

// 返回 (Api_SendVideoMsg (authCode, robot_wxid, to_wxid, path))

// .子程序 发送文件消息, 整数型, 公开, 发送文件（支持好友/群）
// .参数 robot_wxid, 文本型, , 用哪个机器人发这条消息
// .参数 to_wxid, 文本型, , 对方的ID（支持好友/群ID）
// .参数 path, 文本型, , 文件的路径

// 返回 (Api_SendFileMsg (authCode, robot_wxid, to_wxid, path))

// .子程序 发送名片消息, 整数型, 公开, 发送名片消息（支持好友/群）
// .参数 robot_wxid, 文本型, , 用哪个机器人发这条消息
// .参数 to_wxid, 文本型, , 对方的ID（支持好友/群ID）
// .参数 friend_wxid, 文本型, , 要发送的好友/公众hao的wxid

// 返回 (Api_SendCardMsg (authCode, robot_wxid, to_wxid, friend_wxid))

// .子程序 发送群消息并艾特, 整数型, 公开, 发送群聊消息，并且艾特某个群成员
// .参数 robot_wxid, 文本型, , 用哪个机器人发这条消息
// .参数 group_wxid, 文本型, , 要发送的群ID
// .参数 member_wxid, 文本型, , 要艾特成员的ID，支持多个，可用英文,号或者|号分割他们
// .参数 member_name, 文本型, , 要艾特成员的自定义昵称内容，可不填，则取默认的昵称
// .参数 msg, 文本型, , 消息内容

// 返回 (Api_SendGroupMsgAndAt (authCode, robot_wxid, group_wxid, member_wxid, member_name, msg))

// .子程序 发送动态表情, 整数型, 公开, 发送动态表情（动图消息）
// .参数 robot_wxid, 文本型, , 用哪个机器人发这条消息
// .参数 to_wxid, 文本型, , 对方的ID（支持好友/群ID）
// .参数 path, 文本型, , 表情的文件的绝对路径

// 返回 (Api_SendEmojiMsg (authCode, robot_wxid, to_wxid, path))

// .子程序 发送分享链接, 整数型, 公开, 分享链接，可以分享音乐链接，电影链接，美食链接
// .参数 robot_wxid, 文本型, , 用哪个机器人发这条消息
// .参数 to_wxid, 文本型, , 对方的ID（支持好友/群ID）
// .参数 title, 文本型, , 链接标题
// .参数 text, 文本型, , 链接内容
// .参数 target_url, 文本型, 可空, 跳转链接
// .参数 pic_url, 文本型, 可空, 图片的链接
// .参数 icon_url, 文本型, 可空, 图标的链接

// 返回 (Api_SendLinkMsg (authCode, robot_wxid, to_wxid, title, text, target_url, pic_url, icon_url))

// .子程序 发送小程序消息, 整数型, 公开, 发送一个小程序消息
// .参数 robot_wxid, 文本型
// .参数 to_wxid, 文本型, , 对方的ID（支持好友/群ID）
// .参数 xmlContent, 文本型, , 小程序消息的xml内容

// 返回 (Api_SendMiniAppMsg (authCode, robot_wxid, to_wxid, xmlContent))

// .子程序 发送音乐分享, 整数型, 公开, 发送一个音乐消息，分享给你中意的那个人，type：0 随机模式 / 1 网易云音乐 / 2 酷狗音乐 / 3 QQ音乐 / 因乐库原因当前版本都为网易云音乐
// .参数 robot_wxid, 文本型
// .参数 to_wxid, 文本型, , 对方的ID（支持好友/群ID）
// .参数 name, 文本型
// .参数 type, 整数型, , 0 随机模式 / 1 网易云音乐 / 2 酷狗音乐

// 返回 (Api_SendMusicMsg (authCode, robot_wxid, to_wxid, name, type))

// .子程序 转发消息, 整数型, 公开, 对一个消息进行转发操作，稳定效率，可转发任意形式的消息，推荐使用
// .参数 robot_wxid, 文本型
// .参数 to_wxid, 文本型, , 对方的ID（支持好友/群ID）
// .参数 msg_id, 文本型, , 原来的消息id

// 返回 (Api_ForwardMsg (authCode, robot_wxid, to_wxid, msg_id))

// .子程序 取登录账号昵称, 文本型, 公开, 取已登录账号的指定ID昵称
// .参数 robot_wxid, 文本型

// 返回 (Api_GetRobotName (authCode, robot_wxid))

// .子程序 取登录账号头像, 文本型, 公开, 取已登录账号的指定ID头像 * 限测版 *
// .参数 robot_wxid, 文本型

// 返回 (Api_GetRobotHeadimgurl (authCode, robot_wxid))

// .子程序 取登录账号列表, 文本型, 公开, 取已登录的所有机器人账号信息

// 返回 (Api_GetLoggedAccountList (authCode))

// .子程序 取好友列表, 整数型, 公开, 第二个参数不填则取全部账号的数据，第三个参数为真时，会重新加载（切勿频繁使用）
// .参数 friendArr, 好友, 参考 可空 数组, 要接收的好友列表
// .参数 robot_wxid, 文本型, 可空, 如不填，则取的是所有登录账号的好友列表
// .参数 is_refresh, 逻辑型, 可空, 为真将重新加载（注意切记不要频繁加载这里），不然将取缓存，默认为假
// .局部变量 content, 文本型
// .局部变量 json, 类_json
// .局部变量 n, 整数型
// .局部变量 i, 整数型
// .局部变量 friend, 好友

// content ＝ Api_GetFriendList (authCode, robot_wxid, is_refresh)

// .如果真 (json.解析 (content))

//     清除数组 (friendArr)

//     n ＝ json.成员数 ()

//     .计次循环首 (n, i)

//         friend.wxid ＝ json.取通用属性 (“[” ＋ 到文本 (i － 1) ＋ “].wxid”)
//         friend.nickname ＝ json.取通用属性 (“[” ＋ 到文本 (i － 1) ＋ “].nickname”)
//         friend.note ＝ json.取通用属性 (“[” ＋ 到文本 (i － 1) ＋ “].note”)
//         friend.rotob_wxid ＝ json.取通用属性 (“[” ＋ 到文本 (i － 1) ＋ “].robot_wxid”)

//         加入成员 (friendArr, friend)

//     .计次循环尾 ()
// .如果真结束

// 返回 (n)

// .子程序 取群聊列表, 整数型, 公开, 第二个参数不填则取全部账号的数据，第三个参数为真时，会重新加载（切勿频繁使用）
// .参数 groupArr, 群聊, 参考 可空 数组, 要接收的群聊列表
// .参数 robot_wxid, 文本型, 可空, 取哪个账号的列表，不填则取全部
// .参数 is_refresh, 逻辑型, 可空, 为真将重新加载（注意切记不要频繁加载这里），不然将取缓存，默认为假
// .局部变量 content, 文本型
// .局部变量 json, 类_json
// .局部变量 n, 整数型
// .局部变量 i, 整数型
// .局部变量 group, 群聊

// content ＝ Api_GetGroupList (authCode, robot_wxid, is_refresh)

// .如果真 (json.解析 (content))

//     清除数组 (groupArr)

//     n ＝ json.成员数 ()

//     .计次循环首 (n, i)

//         group.wxid ＝ json.取通用属性 (“[” ＋ 到文本 (i － 1) ＋ “].wxid”)
//         group.nickname ＝ json.取通用属性 (“[” ＋ 到文本 (i － 1) ＋ “].nickname”)
//         group.robot_wxid ＝ json.取通用属性 (“[” ＋ 到文本 (i － 1) ＋ “].robot_wxid”)

//         加入成员 (groupArr, group)

//     .计次循环尾 ()
// .如果真结束

// 返回 (n)

// .子程序 取群成员详细, 整数型, 公开, 取某个群的群成员资料，当最后一个参数为真时，会重新加载（切勿频繁使用）
// .参数 robot_wxid, 文本型, , 已登机器人账号ID
// .参数 group_wxid, 文本型, , 群ID
// .参数 member_wxid, 文本型, , 群成员ID
// .参数 member, 好友, 参考, 要接收的群成员信息
// .参数 is_refresh, 逻辑型, 可空, 为真将重新加载（注意切记不要频繁加载这里），不然将取缓存，默认为假
// .局部变量 content, 文本型
// .局部变量 json, 类_json
// .局部变量 n, 整数型
// .局部变量 i, 整数型

// content ＝ Api_GetGroupMemberDetailInfo (authCode, robot_wxid, group_wxid, member_wxid, is_refresh)

// .如果真 (json.解析 (content))

//     member.wxid ＝ json.取通用属性 (“wxid”)
//     member.nickname ＝ json.取通用属性 (“nickname”)
//     member.wx_num ＝ json.取通用属性 (“wx_num”)
//     member.headimgurl ＝ json.取通用属性 (“headimgurl”)
//     member.sex ＝ json.取属性数值 (“sex”)
//     member.city ＝ json.取通用属性 (“city”)
//     member.rotob_wxid ＝ json.取通用属性 (“robot_wxid”)
// .如果真结束

// 返回 (0)

// .子程序 取群成员列表, 整数型, 公开, 取某个群的群成员列表，当最后一个参数为真时，会重新加载（切勿频繁使用）
// .参数 robot_wxid, 文本型, , 已登账号ID
// .参数 group_wxid, 文本型, , 群ID
// .参数 memberArr, 好友, 参考 数组, 要接收的群成员列表
// .参数 is_refresh, 逻辑型, 可空, 为真将重新加载列表（注意切记不要频繁加载这里），不然将取缓存，默认为假
// .局部变量 content, 文本型
// .局部变量 json, 类_json
// .局部变量 n, 整数型
// .局部变量 i, 整数型
// .局部变量 member, 好友

// content ＝ Api_GetGroupMemberList (authCode, robot_wxid, group_wxid, is_refresh)

// .如果真 (json.解析 (content))

//     清除数组 (memberArr)

//     n ＝ json.成员数 ()

//     .计次循环首 (n, i)

//         member.wxid ＝ json.取通用属性 (“[” ＋ 到文本 (i － 1) ＋ “].wxid”)
//         member.nickname ＝ json.取通用属性 (“[” ＋ 到文本 (i － 1) ＋ “].nickname”)
//         member.note ＝ json.取通用属性 (“[” ＋ 到文本 (i － 1) ＋ “].note”)
//         member.rotob_wxid ＝ json.取通用属性 (“[” ＋ 到文本 (i － 1) ＋ “].robot_wxid”)

//         加入成员 (memberArr, member)

//     .计次循环尾 ()
// .如果真结束

// 返回 (n)

// .子程序 取好友列表_json, 文本型, 公开, 第二个参数不填则取全部账号的数据，第三个参数为真时，会重新加载（切勿频繁使用）
// .参数 robot_wxid, 文本型, 可空, 如不填，则取的是所有登录账号的好友列表
// .参数 is_refresh, 逻辑型, 可空, 为真将重新加载（注意切记不要频繁加载这里），不然将取缓存，默认为假
// .参数 out_rawdata, 逻辑型, 可空, 是否输出原始数据，为真则输出原始流（16进制），可用于正常显示emoji等特殊字符
// .局部变量 content, 文本型
// .局部变量 json, 类_json
// .局部变量 n, 整数型
// .局部变量 i, 整数型
// .局部变量 friend, 好友

// content ＝ Api_GetFriendList (authCode, robot_wxid, is_refresh)

// 返回 (content)

// .子程序 取群聊列表_json, 文本型, 公开, 第二个参数不填则取全部账号的数据，第三个参数为真时，会重新加载（切勿频繁使用）
// .参数 robot_wxid, 文本型, 可空, 取哪个账号的列表，不填则取全部
// .参数 is_refresh, 逻辑型, 可空, 为真将重新加载（注意切记不要频繁加载这里），不然将取缓存，默认为假
// .局部变量 content, 文本型
// .局部变量 json, 类_json
// .局部变量 n, 整数型
// .局部变量 i, 整数型
// .局部变量 group, 群聊

// content ＝ Api_GetGroupList (authCode, robot_wxid, is_refresh)

// 返回 (content)

// .子程序 取群成员详细_json, 文本型, 公开, 取某个群的群成员资料，当最后一个参数为真时，会重新加载（切勿频繁使用）
// .参数 robot_wxid, 文本型, , 已登机器人账号ID
// .参数 group_wxid, 文本型, , 群ID
// .参数 member_wxid, 文本型, , 群成员ID
// .参数 is_refresh, 逻辑型, 可空, 为真将重新加载（注意切记不要频繁加载这里），不然将取缓存，默认为假
// .局部变量 content, 文本型
// .局部变量 json, 类_json
// .局部变量 n, 整数型
// .局部变量 i, 整数型

// content ＝ Api_GetGroupMemberDetailInfo (authCode, robot_wxid, group_wxid, member_wxid, is_refresh)

// 返回 (content)

// .子程序 取群成员列表_json, 文本型, 公开, 取某个群的群成员列表，当最后一个参数为真时，会重新加载（切勿频繁使用）
// .参数 robot_wxid, 文本型, , 已登账号ID
// .参数 group_wxid, 文本型, , 群ID
// .参数 is_refresh, 逻辑型, 可空, 为真将重新加载列表（注意切记不要频繁加载这里），不然将取缓存，默认为假
// .局部变量 content, 文本型
// .局部变量 json, 类_json
// .局部变量 n, 整数型
// .局部变量 i, 整数型
// .局部变量 member, 好友

// content ＝ Api_GetGroupMemberList (authCode, robot_wxid, group_wxid, is_refresh)

// 返回 (content)

// .子程序 取联系人头像, 文本型, 公开, 取出好友/群聊/公众号头像
// .参数 robot_wxid, 文本型, , 已登账号ID
// .参数 to_wxid, 文本型, , 好友/群聊/公众号
// .局部变量 content, 文本型
// .局部变量 n, 整数型
// .局部变量 i, 整数型
// .局部变量 member, 好友

// content ＝ Api_GetContactHeadimgurl (authCode, robot_wxid, to_wxid)

// 返回 (content)

// .子程序 接收好友转账, 整数型, 公开, 接收好友转账
// .参数 robot_wxid, 文本型, , 哪个机器人收到的好友转账，就填那个机器人的ID
// .参数 from_wxid, 文本型, , 好友的ID（给你转账的那个人的ID）
// .参数 json_msg, 文本型, , 请传入转账事件里的原消息

// 返回 (Api_AcceptTransfer (authCode, robot_wxid, from_wxid, json_msg))

// .子程序 同意群聊邀请, 整数型, 公开, 同意群聊邀请（适用于需要邀请链接的群聊，小群会直接通过） * 限测版 *
// .参数 robot_wxid, 文本型, , 哪个机器人收到的群聊邀请，就填那个机器人的ID号
// .参数 json_msg, 文本型, , 请传入事件的原消息

// 返回 (Api_AgreeGroupInvite (authCode, robot_wxid, json_msg))

// .子程序 同意好友请求, 整数型, 公开, 同意新的好友请求
// .参数 robot_wxid, 文本型, , 哪个机器人收到的好友验证，就填哪个机器人的那个ID
// .参数 json_msg, 文本型, , 请传入好友验证事件的原消息

// 返回 (Api_AgreeFriendVerify (authCode, robot_wxid, json_msg))

// .子程序 修改好友备注, 整数型, 公开, 修改好友的备注 * 限测版 *
// .参数 robot_wxid, 文本型, , 要操作的机器人ID
// .参数 friend_wxid, 文本型, , 要备注的好友ID
// .参数 note, 文本型, , 新的备注

// 返回 (Api_ModifyFriendNote (authCode, robot_wxid, friend_wxid, note))

// .子程序 删除好友, 整数型, 公开, 删除一个好友 * 限测版 *
// .参数 robot_wxid, 文本型, , 要操作的机器人ID
// .参数 friend_wxid, 文本型, , 要删除的好友ID

// 返回 (Api_DeleteFriend (authCode, robot_wxid, friend_wxid))

// .子程序 取插件信息, 应用信息, 公开

// 返回 (appInfo)

// .子程序 取应用目录, 文本型, 公开, 格式为：D:\可爱猫\app\demo.cat.dll\

// 返回 (Api_GetAppDirectory (authCode))

// .子程序 开启错误捕获, , 公开

// Api_SetFatal (authCode)

// .子程序 添加日志, , 公开, 向框架append一条新的日志，一般只需要填第一个参数，第二个为自动分隔开显示
// .参数 msg1, 文本型
// .参数 msg2, 文本型, 可空
// .局部变量 code, 整数型
// .局部变量 content, 文本型

// content ＝ msg1

// .如果真 (是否为空 (msg2) ＝ 假)
//     content ＝ msg1 ＋ “ | ” ＋ msg2
// .如果真结束

// Api_AppendLogs (authCode, content)

// .子程序 重载插件, 整数型, 公开, 会停止当前插件的运行，请在执行当前代码的下一句，立即返回()，因为执行后，当前插件内部已经销毁，包括线程，之后，会重新加载一遍app目录下的dll，可用于热更新插件（无需重启机器人）

// 返回 (Api_ReloadPlug (authCode))

// .子程序 踢出群成员, 整数型, 公开, 踢出指定群成员（仅机器人为群主或管理员时有效）
// .参数 robot_wxid, 文本型, , 要操作的机器人ID
// .参数 group_wxid, 文本型, , 群ID
// .参数 member_wxid, 文本型, , 群成员ID

// 返回 (Api_RemoveGroupMember (authCode, robot_wxid, group_wxid, member_wxid))

// .子程序 修改群名称, 整数型, 公开, 修改一个群的群名称，切勿频繁使用
// .参数 robot_wxid, 文本型, , 要操作的机器人ID
// .参数 group_wxid, 文本型, , 群ID
// .参数 group_name, 文本型, , 新的群名称

// 返回 (Api_ModifyGroupName (authCode, robot_wxid, group_wxid, group_name))

// .子程序 修改群公告, 整数型, 公开, 发布群聊的新公告，仅机器人为群主或管理员时有效
// .参数 robot_wxid, 文本型, , 要操作的机器人ID
// .参数 group_wxid, 文本型, , 群ID
// .参数 content, 文本型, , 新公告内容

// 返回 (Api_ModifyGroupNotice (authCode, robot_wxid, group_wxid, content))

// .子程序 建立新群, 文本型, 公开, 建立新群，请填写多个好友ID的数组，人数必须大于3个人才可以，成功返回新群的群号 * 限测版 *
// .参数 robot_wxid, 文本型, , 要操作的机器人ID
// .参数 friendArr, 文本型, 数组, 要建立新群的好友数组，至少要两个人以上
// .局部变量 str, 文本型
// .局部变量 n, 整数型
// .局部变量 i, 整数型

// n ＝ 取数组成员数 (friendArr)

// .如果真 (n ＜ 2)
//     返回 (“”)
// .如果真结束

// .计次循环首 (n, i)

//     str ＝ str ＋ friendArr [i] ＋ 选择 (n ＝ i, “”, “|”)

// .计次循环尾 ()

// 返回 (Api_BuildingGroupPlus (authCode, robot_wxid, str))

// .子程序 退出群聊, 整数型, 公开, 退出指定群聊 * 限测版 *
// .参数 robot_wxid, 文本型, , 要操作的机器人ID
// .参数 group_wxid, 文本型, , 群ID

// 返回 (Api_QuitGroup (authCode, robot_wxid, group_wxid))

// .子程序 邀请加入群聊, 整数型, 公开, 邀请好友加入指定群聊，低于40人自动拉取，超过40人自动发送邀请链接，切勿频繁操作 * 限测版 *
// .参数 robot_wxid, 文本型, , 要操作的机器人ID
// .参数 group_wxid, 文本型, , 群ID
// .参数 friend_wxid, 文本型, , 要邀请的好友ID

// 返回 (Api_InviteInGroup (authCode, robot_wxid, group_wxid, friend_wxid))

// .子程序 置顶联系人, 整数型, 公开, 对一个常用的联系人进行置顶操作
// .参数 robot_wxid, 文本型, , 要操作的机器人ID
// .参数 friend_wxid, 文本型, , 需要置顶的人wxid，可以是好友或者群类型

// 返回 (Api_OnTop (authCode, robot_wxid, friend_wxid))

// .子程序 取消置顶联系人, 整数型, 公开, 对一个常用的联系人进行取消置顶操作
// .参数 robot_wxid, 文本型, , 要操作的机器人ID
// .参数 friend_wxid, 文本型, , 需要取消置顶的人wxid，可以是好友或者群类型

// 返回 (Api_OffTop (authCode, robot_wxid, friend_wxid))

// .子程序 取框架版本号, 文本型, 公开

// 返回 (Api_GetFrameVersion (authCode))
