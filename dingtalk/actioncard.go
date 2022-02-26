package dingtalk

import "encoding/json"

const (
	ActionCardBtnVert = 0 //竖直
	ActionCardBtnHorz = 1 //水平
)

type ActionCardButton struct {
	Title     string `json:"title"`
	ActionUrl string `json:"actionURL"`
}

type ActionCardMessage struct {
	MessageType string `json:"msgtype"`
	ActionCard  struct {
		Title          string              `json:"title"`
		Text           string              `json:"text"` //markdown格式的消息
		BtnOrientation int                 `json:"btnOrientation"`
		SingleTitle    string              `json:"singleTitle"` //设置此项和singleURL后，btns无效
		SingleURL      string              `json:"singleURL"`
		Buttons        []*ActionCardButton `json:"btns"`
	} `json:"actionCard"`
}

func NewActionCardMessage() *ActionCardMessage {
	return &ActionCardMessage{
		MessageType: MessageTypeActionCard,
	}
}

func (m *ActionCardMessage) ToByte() ([]byte, error) {
	jsonByte, err := json.Marshal(m)
	return jsonByte, err
}

func (m *ActionCardMessage) SetTitle(title string) *ActionCardMessage {
	m.ActionCard.Title = title
	return m
}
func (m *ActionCardMessage) SetText(text string) *ActionCardMessage {
	m.ActionCard.Text = text
	return m
}

func (m *ActionCardMessage) SetBtnOrientation(btnOrientation int) *ActionCardMessage {
	m.ActionCard.BtnOrientation = btnOrientation
	return m
}

func (m *ActionCardMessage) SetSingleTitle(singleTitle string) *ActionCardMessage {
	m.ActionCard.SingleTitle = singleTitle
	return m
}

func (m *ActionCardMessage) SetSingleURL(singleURL string) *ActionCardMessage {
	m.ActionCard.SingleURL = singleURL
	return m
}

func (m *ActionCardMessage) AddButtons(buttons []*ActionCardButton) *ActionCardMessage {
	m.ActionCard.Buttons = append(m.ActionCard.Buttons, buttons...)
	return m
}
func (m *ActionCardMessage) AddButton(button *ActionCardButton) *ActionCardMessage {
	m.ActionCard.Buttons = append(m.ActionCard.Buttons, button)
	return m
}
