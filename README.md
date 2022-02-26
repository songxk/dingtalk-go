# dingtalk-go
dingtalk oapi



## 使用方式
```
go get github.com/songxk/dingtalk-go
```

### 项目中引用
```
package main

import (
	"fmt"
	"github.com/songxk/dingtalk-go/dingtalk"
)

func main() {
	m := dingtalk.NewTextMessage()
	m.SetContent("Hello")
	c := dingtalk.NewRobotClient(dingtalk.AlertRobotTest)
	_, _ = c.Send(m)
}

```

### 消息发送
```
//text
m := NewTextMessage()
m.SetContent("Hello")
c := NewRobotClient(AlertRobotTest)
_, _ = c.Send(m)

//link
m := NewLinkMessage()
    .SetTitle("title")
    .SetText("text")
    .SetMessageUrl("http://www.baidu.com")
    .SetPicUrl("https://inews.gtimg.com/newsapp_bt/0/12171811596_909/0")
	
NewRobotClient(AlertRobotTest).Send(m)

//markdown
m := NewMarkdownMessage().SetTitle("title").SetText("#### 杭州天气 @188xxxxxxxx \n > 9度，西北风1级，空气良89，相对温度73%\n > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n > ###### 10点20分发布 [天气](https://www.dingtalk.com) \n").
    SetAtMobiles([]string{"188xxxxxxxx"})
c := NewRobotClient(AlertRobotTest)
_, _ = c.Send(m)

//actioncard
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

//feedcard
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


```

### 支持配置多个机器人
```
在robot.go文件中配置robotMap

var robotMap = map[int]*Robot{
	AlertRobotTest: {
		AccessToken:  "accesstoken1",
		Secret:       "secret1",
		RefGroupName: "groupName1",
	},
	AlertRobotSWK: {
		AccessToken:  "accesstoken2",
		Secret:       "secret2",
		RefGroupName: "groupName2",
	},
}

client := NewRobotClient(dingtalk.AlertRobotSWK)

```

