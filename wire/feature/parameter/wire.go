//+build wireinject

package build

import (
	"github.com/google/wire"
)

/*
	name 参数会被传入NewMonster和NewPlayer中，因为参数的匹配是靠类型
*/
func InitMission(name string) Mission {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}
}

func InitMissionParam(mName MonsterParam, pName PlayerParam) Mission {
	wire.Build(NewMonsterParam, NewPlayerParam, NewMission)
	return Mission{}
}
