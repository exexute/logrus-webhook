package logrusWebhook

import "encoding/json"

type DingTalkLink struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

type DingTalkLinkMsg struct {
	DingTalkMsg
	Link *DingTalkLink `json:"link"`
}

func NewDingTalkLink(title, text, picUrl, messageUrl string) *DingTalkLink {
	return &DingTalkLink{
		Title:      title,
		Text:       text,
		PicUrl:     picUrl,
		MessageUrl: messageUrl,
	}
}

func (l *DingTalkLink) String() string {
	data, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(data)
}

func UnmarshalDingTalkLink(content string) (*DingTalkLink, error) {
	link := &DingTalkLink{}
	err := json.Unmarshal([]byte(content), link)
	if err != nil {
		return nil, err
	}
	return link, nil
}
