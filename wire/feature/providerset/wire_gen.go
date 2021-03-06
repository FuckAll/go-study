// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package build

import (
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitEndingA(mName MonsterParam, pName PlayerParam) (EndingA, error) {
	player, err := NewPlayerParam(pName)
	if err != nil {
		return EndingA{}, err
	}
	monster := NewMonsterParam(mName)
	endingA := NewEndingA(player, monster)
	return endingA, nil
}

func InitEndingB(mName MonsterParam, pName PlayerParam) (EndingB, error) {
	player, err := NewPlayerParam(pName)
	if err != nil {
		return EndingB{}, err
	}
	monster := NewMonsterParam(mName)
	endingB := NewEndingB(player, monster)
	return endingB, nil
}

func InitProviderSetEndingA(mName MonsterParam, pName PlayerParam) (EndingA, error) {
	player, err := NewPlayerParam(pName)
	if err != nil {
		return EndingA{}, err
	}
	monster := NewMonsterParam(mName)
	endingA := NewEndingA(player, monster)
	return endingA, nil
}

func InitProviderSetEndingB(mName MonsterParam, pName PlayerParam) (EndingB, error) {
	player, err := NewPlayerParam(pName)
	if err != nil {
		return EndingB{}, err
	}
	monster := NewMonsterParam(mName)
	endingB := NewEndingB(player, monster)
	return endingB, nil
}

// wire.go:

var monsterPlayerSet = wire.NewSet(NewPlayerParam, NewMonsterParam)
