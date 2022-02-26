package dingtalk

import "testing"

func TestActionCardMessage_Send(t *testing.T) {
	m := NewActionCardMessage().SetTitle("title").
		SetText("![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png)" +
			"### 乔布斯 20 年前想打造的苹果咖啡厅Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划").
		AddButtons([]*ActionCardButton{
			&ActionCardButton{
				Title:     "XXX",
				ActionUrl: "http://www.qq.com",
			},
			&ActionCardButton{
				Title:     "NBBB",
				ActionUrl: "http://www.baidu.com",
			},
		}).
		AddButton(&ActionCardButton{
			Title:     "CCC",
			ActionUrl: "http://google.com",
		}).
		SetBtnOrientation(ActionCardBtnVert)
	c := NewRobotClient(AlertRobotTest)
	_, _ = c.Send(m)
}
