// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package build

// Injectors from wire.go:

func InitMission(name string) Mission {
	player := NewPlayer(name)
	monster := NewMonster()
	mission := NewMission(player, monster)
	return mission
}
