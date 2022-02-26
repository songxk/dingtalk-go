package dingtalk

import "testing"

func TestTextMessage_Send(t *testing.T) {
	m := NewTextMessage()
	m.SetContent("Hello")
	c := NewRobotClient(AlertRobotTest)
	_, _ = c.Send(m)
}
