package dingtalk

type Message interface {
	ToByte() ([]byte, error)
}

const (
	MessageTypeText       = "text"
	MessageTypeLink       = "link"
	MessageTypeMarkdown   = "markdown"
	MessageTypeActionCard = "actionCard"
	MessageTypeFeedCard   = "feedCard"
)

type At struct {
	AtMobiles []string `json:"atMobiles"`
	AtUserIds []string `json:"atUserIds"`
	IsAtAll   bool     `json:"isAtAll"`
}
