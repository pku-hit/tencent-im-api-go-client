package tencentInstantMessage

import (
	"testing"
)

const (
	sdkAppid   = "1400290348"
	secretKey  = "715981f67f0ec3754c250af90faccd97b65941c002adf31ea037d6f62d91232c"
	identifier = "administrator"
)

func TestSendTxtMsg(t *testing.T) {
	timClient := NewTencentInstantMessageClient(sdkAppid, secretKey, identifier)
	var msg []string
	msg = append(msg, "Tom & Jack")
	timClient.SendTxtMsg("黄国超", "黄国超", msg)
}

func TestImportAccount(t *testing.T) {
	timClient := NewTencentInstantMessageClient(sdkAppid, secretKey, identifier)
	timClient.ImportAccount("黄国超", "黄国超", "https://wx.qlogo.cn/mmopen/vi_32/AbK5wc86LMiba8Z79FqiaJiabWgTF3HNjprCbTNSicsbWmLZt4rZiboWbwn0IZh7fbWWKFsT8ufFRH7jrvKdjOTdcNw/132")
}
