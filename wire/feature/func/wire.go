//+build wireinject

package build

import "github.com/google/wire"

var kitty = Monster{Name: "kitty"}
var mario = Player{Name: "mario"}
var endingA = EndingA{
	Player:  mario,
	Monster: kitty,
}

/*
	取出实例中某个字段
*/
func InitEndingA(playerParam PlayerParam, monsterParam MonsterParam) (EndingA, func(), error) {
	wire.Build(NewEndingA, NewPlayerParam, NewMonsterParam)

	return EndingA{}, nil, nil
}
