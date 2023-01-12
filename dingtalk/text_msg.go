package dingtalk

type TextMsg struct {
	Msg
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}
