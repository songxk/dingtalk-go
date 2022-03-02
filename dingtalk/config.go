package dingtalk

var conf Conf

type Conf struct {
	RobotMap map[int]*Robot
}

func LoadRobot(robotMap map[int]*Robot) {
	conf.RobotMap = robotMap
}
