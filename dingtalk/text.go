package dingtalk

import "encoding/json"

type TextMessage struct {
	MessageType string `json:"msgtype"`
	Text        struct {
		Content string `json:"content"`
	} `json:"text"`
	At At `json:"at"`
}

func NewTextMessage() *TextMessage {
	return &TextMessage{
		MessageType: MessageTypeText,
	}
}

func (m *TextMessage) ToByte() ([]byte, error) {
	jsonByte, err := json.Marshal(m)
	return jsonByte, err
}

func (m *TextMessage) SetContent(content string) *TextMessage {
	m.Text.Content = content
	return m
}

func (m *TextMessage) SetAt(at At) *TextMessage {
	m.At = at
	return m
}

func (m *TextMessage) SetAtAll(atAll bool) *TextMessage {
	m.At.IsAtAll = atAll
	return m
}

func (m *TextMessage) SetAtMobiles(mobiles []string) *TextMessage {
	m.At.AtMobiles = mobiles
	return m
}
