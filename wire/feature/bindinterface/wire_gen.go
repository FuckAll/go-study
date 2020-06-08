// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package build

// Injectors from wire.go:

func InitAppear(param PlayerParam, monsterParam MonsterParam) (ConcreteAppear, error) {
	iAppear := _wireIAppearValue
	concreteAppear := NewConcreteAppear(iAppear)
	return concreteAppear, nil
}

var (
	_wireIAppearValue = IAppear(nil)
)

func InitNewEndingA(param PlayerParam, monsterParam MonsterParam) (EndingA, error) {
	player, err := NewPlayerParam(param)
	if err != nil {
		return EndingA{}, err
	}
	monster := NewMonsterParam(monsterParam)
	endingA := NewEndingA(player, monster)
	return endingA, nil
}

// wire.go:

var kitty = Monster{Name: "kitty"}