package logrusWebhook

type DingTalkTextMsg struct {
	DingTalkMsg
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}
