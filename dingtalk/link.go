package dingtalk

import "encoding/json"

type LinkMessage struct {
	MessageType string `json:"msgtype"`
	Link        struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		PicUrl     string `json:"picUrl"`
		MessageUrl string `json:"messageUrl"`
	} `json:"link"`
}

func NewLinkMessage() *LinkMessage {
	return &LinkMessage{
		MessageType: MessageTypeLink,
	}
}

func (m *LinkMessage) ToByte() ([]byte, error) {
	jsonByte, err := json.Marshal(m)
	return jsonByte, err
}

func (m *LinkMessage) SetText(text string) *LinkMessage {
	m.Link.Text = text
	return m
}

func (m *LinkMessage) SetTitle(title string) *LinkMessage {
	m.Link.Title = title
	return m
}

func (m *LinkMessage) SetPicUrl(url string) *LinkMessage {
	m.Link.PicUrl = url
	return m
}

func (m *LinkMessage) SetMessageUrl(url string) *LinkMessage {
	m.Link.MessageUrl = url
	return m
}
