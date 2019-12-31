package tencentInstantMessage

var ShowDebug = true

var TENCENT_SUCCESS = "OK"

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

type TencentResponse struct {
	ActionStatus string // 请求的处理结果，OK 表示处理成功，FAIL 表示失败
	ErrorCode    int    // 错误码，0表示成功，非0表示失败
	ErrorInfo    string // 请求处理失败时的错误信息
}

type TencentCheckAccount struct {
	CheckItem []TencentCheckAccountItem
}

type TencentCheckAccountItem struct {
	UserID string // 请求检查的帐号的 UserID
}

type TencentCheckAccountResponse struct {
	TencentResponse
	ResultItem []TencentCheckAccountResponseItem // 单个帐号的结果对象数组
}

type TencentCheckAccountResponseItem struct {
	UserID        string // 请求检查的帐号的 UserID
	ResultCode    int    // 单个帐号的检查结果：0表示成功，非0表示失败
	ResultInfo    string // 单个帐号检查失败时的错误描述信息
	AccountStatus string // 单个帐号的导入状态：Imported 表示已导入，NotImported 表示未导入
}
