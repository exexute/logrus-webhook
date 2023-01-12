package logrusWebhook

import "encoding/json"

type DingTalkMarkdownMsg struct {
	DingTalkMsg
	Markdown *DingTalkMarkdown `json:"markdown"`
}

type DingTalkMarkdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func NewDingTalkMarkdown(title, text string) *DingTalkMarkdown {
	return &DingTalkMarkdown{
		Title: title,
		Text:  text,
	}
}

func (m *DingTalkMarkdown) String() string {
	data, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(data)
}

func UnmarshalDingTalkMarkdown(content string) (*DingTalkMarkdown, error) {
	m := &DingTalkMarkdown{}
	err := json.Unmarshal([]byte(content), m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
