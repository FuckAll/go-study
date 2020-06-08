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
func InitPlayer() Player {
	wire.Build(wire.Value(endingA), wire.FieldsOf(new(EndingA), "Player"))
	return Player{}
}
