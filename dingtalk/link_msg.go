package dingtalk

import "encoding/json"

type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

type LinkMsg struct {
	Msg
	Link *Link `json:"link"`
}

func NewLink(title, text, picUrl, messageUrl string) *Link {
	return &Link{
		Title:      title,
		Text:       text,
		PicUrl:     picUrl,
		MessageUrl: messageUrl,
	}
}

func (l *Link) String() string {
	data, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(data)
}

func UnmarshalLink(content string) (*Link, error) {
	link := &Link{}
	err := json.Unmarshal([]byte(content), link)
	if err != nil {
		return nil, err
	}
	return link, nil
}
