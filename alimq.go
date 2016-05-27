package alimq

var (
	ACCESS_KEY string
	SECRET_KEY string
	URL_PREFIX string
)

var newline string = "\n"

type SendMessage struct {
	Topic      string
	Tag        string
	ProducerId string
	Key        string
	Body       string
	Time       int64
}
