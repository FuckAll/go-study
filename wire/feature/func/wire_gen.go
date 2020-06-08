// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package build

// Injectors from wire.go:

func InitEndingA(playerParam PlayerParam, monsterParam MonsterParam) (EndingA, func(), error) {
	player, cleanup, err := NewPlayerParam(playerParam)
	if err != nil {
		return EndingA{}, nil, err
	}
	monster := NewMonsterParam(monsterParam)
	buildEndingA := NewEndingA(player, monster)
	return buildEndingA, func() {
		cleanup()
	}, nil
}

// wire.go:

var kitty = Monster{Name: "kitty"}

var mario = Player{Name: "mario"}

var endingA = EndingA{
	Player:  mario,
	Monster: kitty,
}
