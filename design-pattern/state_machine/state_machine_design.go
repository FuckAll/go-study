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
	第三种状态机的实现方式：状态模式（这种方式要求所有的状态都要实现对应的状态类接口）
*/

// 所有状态类的接口
type IMario interface {
	obtainMushRoom()
	obtainCape()
	obtainFireFlower()
	meetMonster()
}

// 状态机
/*
	machine := NewMarioStateMachine()
	machine.obtainMushRoom() // 变成超级马里奥
*/
type MarioStateMachine struct {
	currentState IMario
	Score        int
}

// NewMarioStateMachine 创建状态机
func NewMarioStateMachine() *MarioStateMachine {
	ans := &MarioStateMachine{}
	ans.currentState = &SmallMario{ans}
	return ans
}

// 获取蘑菇
func (m *MarioStateMachine) obtainMushRoom() {
	m.currentState.obtainMushRoom()
}

// 获取斗篷
func (m *MarioStateMachine) obtainCape() {
	m.currentState.obtainCape()
}

// 获取火焰
func (m *MarioStateMachine) obtainFireFlower() {
	m.currentState.obtainFireFlower()
}

// 遇到怪兽
func (m *MarioStateMachine) meetMonster() {
	m.currentState.meetMonster()
}

// 小马里奥
type SmallMario struct {
	*MarioStateMachine
}

// 获取蘑菇
func (m *SmallMario) obtainMushRoom() {
	m.MarioStateMachine.currentState = &SuperMario{m.MarioStateMachine}
	m.MarioStateMachine.Score += 100
}

// 获取斗篷
func (m *SmallMario) obtainCape() {
	m.MarioStateMachine.currentState = &CapMario{m.MarioStateMachine}
	m.MarioStateMachine.Score += 200
}

// 获取火焰
func (m *SmallMario) obtainFireFlower() {
	m.MarioStateMachine.currentState = &FireMario{m.MarioStateMachine}
	m.MarioStateMachine.Score += 300
}

// 遇到怪兽
func (m *SmallMario) meetMonster() {
	// do noting...
}

// 超级马里奥
type SuperMario struct {
	*MarioStateMachine
}

// 获取蘑菇
func (m *SuperMario) obtainMushRoom() {

}

// 获取斗篷
func (m *SuperMario) obtainCape() {

}

// 获取火焰
func (m *SuperMario) obtainFireFlower() {

}

// 遇到怪兽
func (m *SuperMario) meetMonster() {

}

// 斗篷马里奥
type CapMario struct {
	*MarioStateMachine
}

// 获取蘑菇
func (m *CapMario) obtainMushRoom() {

}

// 获取斗篷
func (m *CapMario) obtainCape() {

}

// 获取火焰
func (m *CapMario) obtainFireFlower() {

}

// 遇到怪兽
func (m *CapMario) meetMonster() {

}

type FireMario struct {
	*MarioStateMachine
}

// 获取蘑菇
func (m *FireMario) obtainMushRoom() {

}

// 获取斗篷
func (m *FireMario) obtainCape() {

}

// 获取火焰
func (m *FireMario) obtainFireFlower() {

}

// 遇到怪兽
func (m *FireMario) meetMonster() {

}
