package builder_test

import (
	"github.com/go-study/src/design-pattern/builder"
	"testing"
)

func TestCharacterBuilder(t *testing.T) {
	director := new(builder.Director)
	director.SetBuilder(new(builder.CharacterBuilder))
	character := director.Create("张三", "迫击炮、机关枪")
	expect := &builder.Character{
		Name: "张三",
		Arms: "迫击炮、机关枪",
	}
	if *character != *expect {
		t.Errorf("want %+v but got %+v", expect, character)
	}
}
