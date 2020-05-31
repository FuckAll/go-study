package command

/*
	命令模式：将请求封装为一个对象，这样就可以使用不同的请求参数化其他对象，并且能够支持请求的排队执行、记录日志、撤销等
	例如：一个游戏服务端，接收客户端的请求，并且用命令模式实现。
*/
type Command interface {
	execute() error
}

type GotDiamondCommand struct {
}

// 获取钻石
func (gd *GotDiamondCommand) execute() error {
	return nil
}

// 游戏开始
type GotStartCommand struct {
}

// 获取钻石
func (gd *GotStartCommand) execute() error {
	return nil
}

// 碰撞障碍物
type HitObstacleCommand struct {
}

// 获取钻石
func (gd *HitObstacleCommand) execute() error {
	return nil
}

// 归档数据
type ArchiveCommand struct {
}

// 获取钻石
func (gd *ArchiveCommand) execute() error {
	return nil
}

type Event int

const (
	GotDiamond  Event = 1
	GotStar     Event = 2
	HitObstacle Event = 3
	Archive     Event = 4
)

type Request struct {
	event Event
	data  string
}

func (r *Request) getEvent() Event {
	return r.event
}

func (r *Request) getData() string {
	return r.data
}

type GameApplication struct {
}

func (ga *GameApplication) MainLoop() {
	for {
		// 从epoll或者select中获取数据
		requests := []*Request{{event: 1, data: "data"}, {event: 2, data: "data2"}}
		var commandList []Command
		for _, req := range requests {
			var command Command = nil
			switch req.getEvent() {
			case GotDiamond:
				command = new(GotDiamondCommand)
			case GotStar:
				command = new(GotStartCommand)
			case HitObstacle:
				command = new(HitObstacleCommand)
			case Archive:
				command = new(ArchiveCommand)
			}
			commandList = append(commandList, command)
		}
		for _, cmd := range commandList {
			cmd.execute()
		}
	}
}
