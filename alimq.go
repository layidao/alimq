package alimq

var newline string = "\n"

var (
	ACCESS_KEY string
	SECRET_KEY string
	URL_PREFIX string
)

type SendMessage struct {
	Topic      string
	Tag        string
	ProducerId string
	Key        string
	Body       string
	Time       int64
}

type Messages struct {
	Topic      string
	Tag        string
	ConsumerId string
	// Key        string
	// Body       string
	Time int64
}

type Message struct {
	Body           string `json:"body"`
	BornTime       string `json:"bornTime"`
	Key            string `json:"key"`
	MsgHandle      string `json:"msgHandle"`
	MsgId          string `json:"msgId"`
	ReconsumeTimes int    `json:"reconsumeTimes"`
	Tag            string `json:"tag"`
}

// 生产者状态码
func getStatusCodeMessage(statusCode int) string {
	switch statusCode {
	case 200:
		return ""
	case 201:
		return ""
	case 204:
		return ""
	case 400:
		return "请求失败"
	case 403:
		return "鉴权失败"
	case 408:
		return "请求超时"
	default:
		return "未知错误"
	}
}
