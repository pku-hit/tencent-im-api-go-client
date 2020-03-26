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

type TencentQueryState struct {
	IsNeedDetail int      // 是否需要返回详细的登录平台信息。0表示不需要，1表示需要
	To_Account   []string // 需要查询这些 UserID 的登录状态，一次最多查询500个 UserID 的状态
}

type TencentQueryStateResponse struct {
	ActionStatus string                        // 请求处理的结果，“OK” 表示处理成功，“FAIL” 表示失败
	ErrorCode    int                           // 错误码，0表示成功，非0表示失败
	ErrorInfo    string                        // 详细错误信息
	QueryResult  []TencentQueryStateResultItem // 返回的用户在线状态结构化信息
}

type TencentQueryStateResult struct {
	To_Account string // 返回的用户的 UserID
	/**
	 * 返回的用户状态，目前支持的状态有：
	 * Online：客户端登录后和即时通信 IM 后台有长连接
	 * PushOnline：iOS 和 Android 进程被 kill 或因网络问题掉线，进入 PushOnline 状态，此时仍然可以接收消息的离线推送。客户端切到后台，但是进程未被手机操作系统 kill 掉时，此时状态仍是 Online
	 * Offline：客户端主动退出登录或者客户端自上一次登录起7天之内未登录过
	 * 如果用户是多终端登录，则只要有一个终端的状态是 Online ，该字段值就是 Online
	 */
	State  string
	Detail []TencentQueryStateResultItem // 详细的登录平台信息
}

type TencentQueryStateResultItem struct {
	Platform string // 登录的平台类型
	Status   string // 该登录平台的状态
}
