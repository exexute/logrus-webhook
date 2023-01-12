package dingtalk

type Msg struct {
	At      *At    `json:"at"`
	MsgType string `json:"msgtype"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	AtUserIds []string `json:"atUserIds"`
	IsAtAll   bool     `json:"isAtAll"`
}

func NewAt(atAll bool, atMobiles []string, atUserIds ...string) *At {
	return &At{
		AtMobiles: atMobiles,
		AtUserIds: atUserIds,
		IsAtAll:   atAll,
	}
}
