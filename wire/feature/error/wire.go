//+build wireinject

package build

import (
	"github.com/google/wire"
)

/*
	如果有错误，必须返回错误
*/
func InitMissionParam(mName MonsterParam, pName PlayerParam) (Mission, error) {
	wire.Build(NewMonsterParam, NewPlayerParam, NewMission)
	return Mission{}, nil
}
