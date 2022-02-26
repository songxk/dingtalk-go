package dingtalk

import (
	"fmt"
	"testing"
)

func TestFeedCardMessage_Send(t *testing.T) {
	m := NewFeedCardMessage()
	m.AddLink(&FeedCardLink{
		Title:      "时代的火车向前开1",
		MessageUrl: "https://www.dingtalk.com/",
		PicUrl:     "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
	}).AddLinks([]*FeedCardLink{
		&FeedCardLink{
			Title:      "时代的火车向前开1",
			MessageUrl: "https://www.dingtalk.com/",
			PicUrl:     "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
		},
		&FeedCardLink{
			Title:      "时代的火车向前开1",
			MessageUrl: "https://www.dingtalk.com/",
			PicUrl:     "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
		},
	})
	c := NewRobotClient(AlertRobotTest)
	r, err := c.Send(m)
	fmt.Println(r, err)

}
