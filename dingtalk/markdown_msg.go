package dingtalk

import "encoding/json"

type MarkdownMsg struct {
	Msg
	Markdown *Markdown `json:"markdown"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func NewMarkdown(title, text string) *Markdown {
	return &Markdown{
		Title: title,
		Text:  text,
	}
}

func (m *Markdown) String() string {
	data, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(data)
}

func UnmarshalMarkdown(content string) (*Markdown, error) {
	m := &Markdown{}
	err := json.Unmarshal([]byte(content), m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
