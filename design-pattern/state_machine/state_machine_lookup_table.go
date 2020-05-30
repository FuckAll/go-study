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
	第二种状态机的实现方式：查表法（查看目录中查表法图片）
*/

// 事件转移表
var TransitionTable = [][]int{
	{SUPER, CAPE, FIRE, SMALL},
	{SUPER, CAPE, FIRE, SMALL},
	{CAPE, CAPE, CAPE, SMALL},
	{FIRE, FIRE, FIRE, SMALL},
}
var ActionTable = [][]int{
	{+100, +200, +300, +0},
	{+0, +200, +300, -100},
	{+0, +0, +0, -200},
	{+0, +0, +0, -300},
}

type Event int

const (
	GOT_MUSHROOM Event = 0
	GOT_CAP      Event = 1
	GOT_FIRE     Event = 2
	MEET_MONSTER Event = 3
)

type MarioStateMachineLookUpTable struct {
	Score int
	State int
}

// 获取蘑菇
func (m *MarioStateMachineLookUpTable) obtainMushRoom() {
	m.executeEvent(GOT_MUSHROOM)
}

// 获取斗篷
func (m *MarioStateMachineLookUpTable) obtainCape() {
	m.executeEvent(GOT_CAP)
}

// 获取火焰
func (m *MarioStateMachineLookUpTable) obtainFireFlower() {
	m.executeEvent(GOT_FIRE)
}

// 遇到怪兽
func (m *MarioStateMachineLookUpTable) meetMonster() {
	m.executeEvent(MEET_MONSTER)
}

// 执行事件引起的状态转移和时间
func (m *MarioStateMachineLookUpTable) executeEvent(event Event) {
	m.State = TransitionTable[m.State][event]
	m.Score = ActionTable[m.State][event]
}
