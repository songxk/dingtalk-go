package dingtalk

type Robot struct {
	AccessToken  string //access token
	Secret       string //secret
	RefGroupName string //机器人所在群名称
}

const (
	AlertRobotTest = iota //test
	AlertRobotSWK         //孙悟空
)

//该列表可以在client里面调用NewRobotClient传入robotId
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
