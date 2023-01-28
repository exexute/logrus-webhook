package logrusWebhook

type FeiShuMsg struct {
	MsgType   string `json:"msg_type"`
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
}
