package dingtalk

import "encoding/json"

type MarkdownMessage struct {
	MessageType string `json:"msgtype"`
	Markdown    struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At At `json:"at"`
}

func NewMarkdownMessage() *MarkdownMessage {
	return &MarkdownMessage{
		MessageType: MessageTypeMarkdown,
	}
}

func (m *MarkdownMessage) ToByte() ([]byte, error) {
	jsonByte, err := json.Marshal(m)
	return jsonByte, err
}

func (m *MarkdownMessage) SetTitle(title string) *MarkdownMessage {
	m.Markdown.Title = title
	return m
}

func (m *MarkdownMessage) SetText(text string) *MarkdownMessage {
	m.Markdown.Text = text
	return m
}

func (m *MarkdownMessage) SetAt(at At) *MarkdownMessage {
	m.At = at
	return m
}

func (m *MarkdownMessage) SetAtAll(atAll bool) *MarkdownMessage {
	m.At.IsAtAll = atAll
	return m
}

func (m *MarkdownMessage) SetAtMobiles(mobiles []string) *MarkdownMessage {
	m.At.AtMobiles = mobiles
	return m
}
