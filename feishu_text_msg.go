package logrusWebhook

type FeiShuTextMsg struct {
	FeiShuMsg
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}
