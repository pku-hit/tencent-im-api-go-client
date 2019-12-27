package tencentInstantMessage

var ShowDebug = true

type TencentMessage struct {
	SyncOtherMachine int           // 1：把消息同步到 From_Account 在线终端和漫游上（默认）；2：消息不同步至 From_Account；（选填）
	FromAccount      string        `json:"From_Account"` // 消息发送方 Identifier（用于指定发送消息方帐号）（选填）
	ToAccount        string        `json:"To_Account"`   // 消息接收方 Identifier（必填）
	MsgLifeTime      int           // 消息离线保存时长（单位：秒），最长为7天（604800秒）（选填）
	MsgRandom        int32         // 消息随机数，由随机函数产生，用于后台定位问题（必填）
	MsgTimeStamp     int64         // 消息时间戳，UNIX 时间戳（单位：秒）
	MsgBody          []interface{} // 消息内容
}

type TIMMsgElement struct {
	MsgType    string
	MsgContent TIMContent
}

type TIMContent struct {
	Text string
}

type TencentAccount struct {
	Identifier string // 用户名，长度不超过32字节
	Nick       string // 用户昵称
	FaceUrl    string // 用户头像 URL
}
