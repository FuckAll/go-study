//+build wireinject

package build

import "github.com/google/wire"

var kitty = Monster{Name: "kitty"}

/*
	接口绑定：
*/
func InitAppear(param PlayerParam, monsterParam MonsterParam) (ConcreteAppear, error) {
	wire.Build(wire.InterfaceValue(new(IAppear), IAppear(nil)), NewConcreteAppear)
	return ConcreteAppear{}, nil
}

func InitNewEndingA(param PlayerParam, monsterParam MonsterParam) (EndingA, error) {
	wire.Build(NewEndingA, NewPlayerParam, NewMonsterParam)
	return EndingA{}, nil
}
