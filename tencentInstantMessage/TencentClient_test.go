package tencentInstantMessage

import (
	"fmt"
	"testing"
)

const (
	sdkAppid   = "1400290348"
	secretKey  = "715981f67f0ec3754c250af90faccd97b65941c002adf31ea037d6f62d91232c"
	identifier = "administrator"
)

func TestSendTxtMsg(t *testing.T) {
	timClient := NewTencentInstantMessageClient(sdkAppid, secretKey, identifier, 180*86400)
	var msg []string
	msg = append(msg, "Tom & Jack")
	timClient.SendTxtMsg("mary", "tom", msg)
}

func TestImportAccount(t *testing.T) {
	timClient := NewTencentInstantMessageClient(sdkAppid, secretKey, identifier, 180*86400)
	timClient.ImportAccount("tom", "tom", "https://wx.qlogo.cn/mmopen/vi_32/AbK5wc86LMiba8Z79FqiaJiabWgTF3HNjprCbTNSicsbWmLZt4rZiboWbwn0IZh7fbWWKFsT8ufFRH7jrvKdjOTdcNw/132")
}

func TestCheckAccount(t *testing.T) {
	timClient := NewTencentInstantMessageClient(sdkAppid, secretKey, identifier, 180*86400)
	timClient.CheckAccount([]string{"tom"})
}

func TestGetUserSigWithUser(t *testing.T) {
	timClient := NewTencentInstantMessageClient(sdkAppid, secretKey, identifier, 180*86400)
	userId := "1217029546898300928"
	userSig := timClient.GetUserSigWithUser(userId)
	fmt.Println(userSig)
}

func TestQueryState(t *testing.T) {
	timClient := NewTencentInstantMessageClient(sdkAppid, secretKey, identifier, 180*86400)
	accountIds := []string{"1218067429370359808", "1218069177208139776"}
	response, _ := timClient.QueryState(1, accountIds)
	fmt.Println(response)
}
