//+build wireinject

package build

import "github.com/google/wire"

var kitty = Monster{Name: "kitty"}

/*
	值绑定：用已有的值进行绑定，例如单例
*/
func InitEndingA(mName MonsterParam, pName PlayerParam) (EndingA, error) {
	wire.Build(NewEndingA, NewPlayerParam, wire.Value(kitty))
	return EndingA{}, nil
}

func InitEndingB(mName MonsterParam, pName PlayerParam) (EndingB, error) {
	wire.Build(NewEndingB, NewPlayerParam, wire.Value(kitty))
	return EndingB{}, nil
}
