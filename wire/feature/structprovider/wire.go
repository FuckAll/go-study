//+build wireinject

package build

import "github.com/google/wire"

/*
	结构构造器：wire.Struct 直接生成结构体，并且初始化结构体
*/
func InitEndingA(mName MonsterParam, pName PlayerParam) (EndingA, error) {
	wire.Build(wire.Struct(new(EndingA), "Player", "Monster"), NewPlayerParam, NewMonsterParam)
	return EndingA{}, nil
}

func InitEndingB(mName MonsterParam, pName PlayerParam) (EndingB, error) {
	wire.Build(wire.Struct(new(EndingB), "*"), NewPlayerParam, NewMonsterParam)
	return EndingB{}, nil
}

var monsterPlayerSet = wire.NewSet(NewPlayerParam, NewMonsterParam)
var endingASet = wire.NewSet(wire.Struct(new(EndingA), "*"), monsterPlayerSet)

func InitProviderSetEndingA(mName MonsterParam, pName PlayerParam) (*EndingA, error) {
	wire.Build(endingASet)
	return &EndingA{}, nil
}

//
//func InitProviderSetEndingB(mName MonsterParam, pName PlayerParam) (EndingB, error) {
//	wire.Build(NewEndingB, monsterPlayerSet)
//	return EndingB{}, nil
//}
