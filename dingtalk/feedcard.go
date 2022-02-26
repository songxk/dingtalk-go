package dingtalk

import "encoding/json"

type FeedCardLink struct {
	Title      string `json:"title"`
	MessageUrl string `json:"messageURL"`
	PicUrl     string `json:"picURL"`
}

type FeedCardMessage struct {
	MessageType string `json:"msgtype"`
	FeedCard    struct {
		Links []*FeedCardLink `json:"links"`
	} `json:"feedCard"`
}

func NewFeedCardMessage() *FeedCardMessage {
	return &FeedCardMessage{
		MessageType: MessageTypeFeedCard,
	}
}

func (m *FeedCardMessage) ToByte() ([]byte, error) {
	jsonByte, err := json.Marshal(m)
	return jsonByte, err
}

func (m *FeedCardMessage) AddLink(link *FeedCardLink) *FeedCardMessage {
	m.FeedCard.Links = append(m.FeedCard.Links, link)
	return m
}

func (m *FeedCardMessage) AddLinks(links []*FeedCardLink) *FeedCardMessage {
	m.FeedCard.Links = append(m.FeedCard.Links, links...)
	return m
}
