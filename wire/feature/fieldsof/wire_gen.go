// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package build

// Injectors from wire.go:

func InitPlayer() Player {
	buildEndingA := _wireEndingAValue
	player := buildEndingA.Player
	return player
}

var (
	_wireEndingAValue = endingA
)

// wire.go:

var kitty = Monster{Name: "kitty"}

var mario = Player{Name: "mario"}

var endingA = EndingA{
	Player:  mario,
	Monster: kitty,
}
