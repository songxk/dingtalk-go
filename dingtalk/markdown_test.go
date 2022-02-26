package dingtalk

import "testing"

func TestMarkdownMessage_Send(t *testing.T) {
	m := NewMarkdownMessage().SetTitle("title").SetText("#### 杭州天气 @188xxxxxxxx \n > 9度，西北风1级，空气良89，相对温度73%\n > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n > ###### 10点20分发布 [天气](https://www.dingtalk.com) \n").
		SetAtMobiles([]string{"188xxxxxxxx"})
	c := NewRobotClient(AlertRobotTest)
	_, _ = c.Send(m)
}
