//+build wireinject

package build

import "github.com/google/wire"

/*
	将Provider变成一个集合，方便实用，如果有变动，只需要改集合本身
*/
func InitEndingA(mName MonsterParam, pName PlayerParam) (EndingA, error) {
	wire.Build(NewEndingA, NewPlayerParam, NewMonsterParam)
	return EndingA{}, nil

}

func InitEndingB(mName MonsterParam, pName PlayerParam) (EndingB, error) {
	wire.Build(NewEndingB, NewPlayerParam, NewMonsterParam)
	return EndingB{}, nil
}

var monsterPlayerSet = wire.NewSet(NewPlayerParam, NewMonsterParam)

func InitProviderSetEndingA(mName MonsterParam, pName PlayerParam) (EndingA, error) {
	wire.Build(NewEndingA, monsterPlayerSet)
	return EndingA{}, nil
}

func InitProviderSetEndingB(mName MonsterParam, pName PlayerParam) (EndingB, error) {
	wire.Build(NewEndingB, monsterPlayerSet)
	return EndingB{}, nil
}
