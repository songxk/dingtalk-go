package dingtalk

import "testing"

func TestLinkMessage_Send(t *testing.T) {
	m := NewLinkMessage().SetTitle("title").SetText("text").SetMessageUrl("http://www.baidu.com").SetPicUrl("https://inews.gtimg.com/newsapp_bt/0/12171811596_909/0")
	NewRobotClient(AlertRobotTest).Send(m)
}
