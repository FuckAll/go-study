//+build wireinject

package build

import (
	"github.com/google/wire"
)

/*
	首先这个函数的返回值就是我们需要创建的对象类型，wire只需要知道类型，return后返回什么不重要。
	然后在函数中，我们调用wire.Build()将创建Mission所依赖的类型的构造器传进去。
	例如，需要调用NewMission()创建Mission类型，NewMission()接受两个参数一个Monster类型，一个Player类型。
	Monster类型对象需要调用NewMonster()创建，Player类型对象需要调用NewPlayer()创建。
	所以NewMonster()和NewPlayer()我们也需要传给wire。
	wire.go和wire_gen.go文件头部位置都有一个+build，不过一个后面是wireinject，另一个是!wireinject。
	+build其实是 Go 语言的一个特性。类似 C/C++ 的条件编译，在执行go build时可传入一些选项，
	根据这个选项决定某些文件是否编译。wire工具只会处理有wireinject的文件，
	所以我们的wire.go文件要加上这个。生成的wire_gen.go是给我们来使用的，wire不需要处理，故有!wireinject。
	注意，如果你运行时，出现了InitMission重定义，那么检查一下你的//+build wireinject与package main这两行之间是否有空行，这个空行必须要有！
*/
func InitMission(name string) Mission {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}
}
