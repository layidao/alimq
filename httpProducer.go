package alimq

import (
	// "crypto/hmac"
	// "crypto/sha1"
	"encoding/json"
	"errors"
	// "fmt"
	"strconv"
)

// 获取发送消息地址
func (this *SendMessage) getPostUrl() string {
	return URL_PREFIX + "/message/?topic=" + this.Topic + "&time=" + strconv.FormatInt(this.Time, 10) + "&tag=" + this.Tag + "&key=" + this.Key
}

// 获取发送消息签名字符串
func (this *SendMessage) getSignStr() string {
	return this.Topic + newline + this.ProducerId + newline + Md5(this.Body) + newline + strconv.FormatInt(this.Time, 10)
}

// 发送信息
func (this *SendMessage) Send() (string, error) {

	url := this.getPostUrl()
	signStr := this.getSignStr()

	sign := HamSha1(signStr, []byte(SECRET_KEY))

	header := make(map[string]string)
	header["AccessKey"] = ACCESS_KEY
	header["Signature"] = sign
	header["ProducerId"] = this.ProducerId

	body, status, err := httpPost(url, header, []byte(this.Body))
	if err != nil {
		return "", err
	}

	statusMessage := getStatusCodeMessage(status)
	if statusMessage != "" {
		return "", errors.New(statusMessage)
	}

	var rs interface{}
	err = json.Unmarshal(body, &rs)
	if err != nil {
		return "", err
	}

	result := rs.(map[string]interface{})

	sendStatus := result["sendStatus"].(string)
	if sendStatus != "SEND_OK" {
		return "", errors.New(sendStatus)
	}

	return result["msgId"].(string), nil
}
