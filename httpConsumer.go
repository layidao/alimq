package alimq

import (
	"encoding/json"
	"errors"
	// "fmt"
	"strconv"
)

func (this *Messages) getUrl() string {
	return URL_PREFIX + "/message/?topic=" + this.Topic + "&time=" + strconv.FormatInt(this.Time, 10) + "&num=32"
}

// 获取发送消息签名字符串
func (this *Messages) getSignStr() string {
	return this.Topic + newline + this.ConsumerId + newline + strconv.FormatInt(this.Time, 10)
}

// 接收信息
func (this *Messages) List() (*[]Message, error) {

	url := this.getUrl()
	signStr := this.getSignStr()
	sign := HamSha1(signStr, []byte(SECRET_KEY))

	header := make(map[string]string)
	header["AccessKey"] = ACCESS_KEY
	header["Signature"] = sign
	header["ConsumerId"] = this.ConsumerId

	body, status, err := HttpGet(url, header)
	if err != nil {
		return nil, err
	}

	statusMessage := getStatusCodeMessage(status)
	if statusMessage != "" {
		return nil, errors.New(statusMessage)
	}

	list := make([]Message, 32)
	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}
	return &list, nil
}

// 删除地址
func (this *Messages) getDelUrl(msgHandle string) string {
	return URL_PREFIX + "/message/?topic=" + this.Topic + "&time=" + strconv.FormatInt(this.Time, 10) + "&msgHandle=" + msgHandle
}

// 签名字符串
func (this *Messages) getDelSignStr(msgHandle string) string {
	return this.Topic + newline + this.ConsumerId + newline + msgHandle + newline + strconv.FormatInt(this.Time, 10)
}

// 删除消息
func (this *Messages) Delete(msgHandle string) (bool, error) {

	url := this.getDelUrl(msgHandle)
	signStr := this.getDelSignStr(msgHandle)
	sign := HamSha1(signStr, []byte(SECRET_KEY))

	header := make(map[string]string)
	header["AccessKey"] = ACCESS_KEY
	header["Signature"] = sign
	header["ConsumerId"] = this.ConsumerId

	_, status, err := HttpDelete(url, header)
	if err != nil {
		return false, err
	}

	statusMessage := getStatusCodeMessage(status)
	if statusMessage != "" {
		return false, errors.New(statusMessage)
	}

	return true, nil
}
