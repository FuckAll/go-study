package memento_test

import (
	"github.com/go-study/src/design-pattern/memento"
	"testing"
)

/*
	备忘录模式：在不违背封装原则的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态，一遍之后恢复为先前的状态。
	实现如下例子：输入文本，用户输入:list则展示当前的文本，:undo则撤销上一次输入的文本。
	> hello
    >:list
    hello
    >world
    >:list
    helloworld
    >:undo
    >:list
	hello
*/

func TestSnapshot(t *testing.T) {
	input := new(memento.InputText)
	snapshotHolder := new(memento.SnapshotHolder)

	// > hello
	snapshotHolder.PushSnapshots(input.CreateSnapshot())
	input.Append("hello")

	// >:list
	if input.GetText() != "hello" {
		t.Errorf(":list want hello, but get %s", input.GetText())
	}
	// >world
	snapshotHolder.PushSnapshots(input.CreateSnapshot())
	input.Append("world")

	// >:list
	if input.GetText() != "helloworld" {
		t.Errorf(":list want helloworld, but get %s", input.GetText())
	}

	// >:undo
	target := snapshotHolder.PopSnapshots()
	input.RestoreSnapshot(target)

	// >:list
	if input.GetText() != "hello" {
		t.Errorf(":list want hello, but get %s", input.GetText())
	}

}
