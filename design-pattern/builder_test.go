package design

import "testing"

func TestCharacterBuilder(t *testing.T) {
	director := new(Director)
	director.builder = new(CharacterBuilder)
	character := director.Create("张三", "迫击炮、机关枪")
	expect := &Character{
		Name: "张三",
		Arms: "迫击炮、机关枪",
	}
	if *character != *expect {
		t.Errorf("want %+v but got %+v", expect, character)
	}
}
