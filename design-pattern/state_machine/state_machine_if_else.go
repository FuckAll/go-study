package state_machine

/*
	有限状态机：状态、事件、动作
	以游戏为例：在游戏中，马里奥可以变身为多种形态，比如小马里奥（Small Mario）、超级马里奥（Super Mario）、火焰马里奥（Fire Mario）、
    斗篷马里奥（Cape Mario）等等。在不同的游戏情节下，各个形态会互相转化，并相应的增减积分。
     比如，初始形态是小马里奥，吃了蘑菇之后就会变成超级马里奥，并且增加 100 积分。
	实际上，马里奥形态的转变就是一个状态机。
    其中，马里奥的不同形态就是状态机中的“状态”，游戏情节（比如吃了蘑菇）就是状态机中的“事件”，加减积分就是状态机中的“动作”。
    比如，吃蘑菇这个事件，会触发状态的转移：从小马里奥转移到超级马里奥，以及触发动作的执行（增加 100 积分）。
*/

/*
	第一种状态机的实现方式：用代码堆砌分支逻辑
*/
const (
	SMALL = 0
	SUPER = 1
	FIRE  = 2
	CAPE  = 3
)

type MarioStateMachineIfElse struct {
	Score int
	State int
}

// 获取蘑菇
func (m *MarioStateMachineIfElse) obtainMushRoom() {
	if m.State == SMALL {
		// 小马里奥获取蘑菇之后变成超级马里奥，并且加100分
		m.State = SUPER
		m.Score += 100
	}
}

// 获取斗篷
func (m *MarioStateMachineIfElse) obtainCape() {
	// 小玛丽奥或者超级马里奥获取斗篷，会变成斗篷马里奥，并且加分200
	if m.State == SMALL || m.State == SUPER {
		m.State = CAPE
		m.Score += 200
	}
}

// 获取火焰
func (m *MarioStateMachineIfElse) obtainFireFlower() {
	// 小玛丽奥或者超级马里奥获取火焰，变成火焰马里奥，并且加分300
	if m.State == SMALL || m.State == SUPER {
		m.State = FIRE
		m.Score += 300
	}
}

// 遇到怪兽
func (m *MarioStateMachineIfElse) meetMonster() {
	if m.State == SUPER {
		m.State = SMALL
		m.Score -= 100
	}
	if m.State == CAPE {
		m.State = SMALL
		m.Score -= 200
	}
	if m.State == FIRE {
		m.State = SMALL
		m.Score -= 300
	}
}
