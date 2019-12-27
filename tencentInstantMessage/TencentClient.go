package tencentInstantMessage

import (
	"fmt"
	"github.com/franela/goreq"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

type TencentMessageClient struct {
	SdkAppId   string
	SecretKey  string
	Identifier string
}

func NewTencentInstantMessageClient(sdkAppId string, secretKey string, identifier string) *TencentMessageClient {
	return &TencentMessageClient{
		SdkAppId:   sdkAppId,
		SecretKey:  secretKey,
		Identifier: identifier,
	}
}

func (timClient *TencentMessageClient) buildReq(uri string, method string, body interface{}) (*goreq.Request, error) {
	if len(timClient.SdkAppId) == 0 || len(timClient.SecretKey) == 0 || len(timClient.Identifier) == 0 {
		return nil, fmt.Errorf("invalidate SdkAppId/SecretKey/Identifier")
	}

	req := goreq.Request{
		Method:      method, //"POST",
		Uri:         uri,
		Accept:      "application/json",
		ContentType: "application/json",
		UserAgent:   "Tencent-IM-API-GO-Client",
		Timeout:     30 * time.Second, //30s
	}
	req.Body = body
	req.ShowDebug = ShowDebug

	return &req, nil
}

func (timClient *TencentMessageClient) request(uri string, method string, body interface{}) (*goreq.Response, error) {
	req, err := timClient.buildReq(uri, method, body)
	if nil != err {
		return nil, err
	}

	res, err := req.Do()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (timClient *TencentMessageClient) buildUrl(subUrl string) string {
	appId, _ := strconv.Atoi(timClient.SdkAppId)
	userSig, _ := GenSig(appId, timClient.SecretKey, timClient.Identifier, 1000)
	url := TENCENT_IM_SERVER_URL + fmt.Sprintf(subUrl, timClient.SdkAppId, timClient.Identifier, userSig, rand.Uint32())
	return url
}

/**
 * sendMessage
 */
func (timClient *TencentMessageClient) SendTxtMsg(fromAccount string, toAccount string, messages []string) error {
	url := timClient.buildUrl(SEND_MESSAGE_URL)

	var msgBodys []interface{}
	for _, msg := range messages {
		timElement := TIMMsgElement{}
		timElement.MsgType = "TIMTextElem"
		timElement.MsgContent = TIMContent{
			Text: msg,
		}
		msgBodys = append(msgBodys, timElement)
	}

	random := rand.Int()
	fmt.Println(random)
	timMsg := TencentMessage{
		SyncOtherMachine: 1,
		FromAccount:      fromAccount,
		ToAccount:        toAccount,
		MsgRandom:        rand.Int31(),
		MsgTimeStamp:     time.Now().Unix(),
	}
	timMsg.MsgBody = msgBodys

	resp, err := timClient.request(url, "POST", timMsg)
	if nil != err {
		return err
	}
	defer resp.Body.Close()

	ibytes, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return err
	}

	if ShowDebug {
		fmt.Println("respone:", string(ibytes))
	}

	return nil
}

// import account
func (timClient *TencentMessageClient) ImportAccount(identifier, nick, faceUrl string) error {
	if len([]byte(identifier)) > 32 {
		fmt.Println("param identifier byte length greater than 32")
	}
	url := timClient.buildUrl(ACCOUNT_IMPORT_URL)

	account := TencentAccount{
		Identifier: identifier,
		Nick:       nick,
		FaceUrl:    faceUrl,
	}

	resp, err := timClient.request(url, "POST", account)
	if nil != err {
		fmt.Println(err.Error())
		return err
	}
	defer resp.Body.Close()

	ibytes, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		fmt.Println(err.Error())
		return err
	}

	if ShowDebug {
		fmt.Println("respone:", string(ibytes))
	}

	return nil
}
