package memento

import "strings"

/*
	备忘录模式：在不违背封装原则的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态，一遍之后恢复为先前的状态。
	实现如下例子：输入文本，用户输入:list则展示当前的文本，:undo则撤销上一次输入的文本。
	> hello
    > :list
    hello
    >world
    >:list
    helloworld
    >:undo
    >:list
	hello
*/

type InputText struct {
	text strings.Builder
}

func (i *InputText) CreateSnapshot() *Snapshot {
	return &Snapshot{text: i.text}
}

func (i *InputText) Append(input string) {
	i.text.WriteString(input)
}

func (i *InputText) RestoreSnapshot(snapshot *Snapshot) {
	i.text = snapshot.text
}

func (i *InputText) GetText() string {
	return i.text.String()
}

type Snapshot struct {
	text strings.Builder
}

func (s *Snapshot) GetText() string {
	return s.text.String()
}

type SnapshotHolder struct {
	// 每次存储上一个版本
	snapshots []*Snapshot
}

func (sh *SnapshotHolder) PushSnapshots(snapshot *Snapshot) {
	sh.snapshots = append(sh.snapshots, snapshot)
}

func (sh *SnapshotHolder) PopSnapshots() *Snapshot {
	length := len(sh.snapshots)
	if length <= 0 {
		return nil
	}
	ans := sh.snapshots[length-1]
	sh.snapshots = sh.snapshots[:length-1]
	return ans
}
