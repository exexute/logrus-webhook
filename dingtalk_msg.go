package logrusWebhook

type DingTalkMsg struct {
	At      *DingTalkAt `json:"at"`
	MsgType string      `json:"msgtype"`
}

type DingTalkAt struct {
	AtMobiles []string `json:"atMobiles"`
	AtUserIds []string `json:"atUserIds"`
	IsAtAll   bool     `json:"isAtAll"`
}

func NewDingTalkAt(atAll bool, atMobiles []string, atUserIds ...string) *DingTalkAt {
	return &DingTalkAt{
		AtMobiles: atMobiles,
		AtUserIds: atUserIds,
		IsAtAll:   atAll,
	}
}
