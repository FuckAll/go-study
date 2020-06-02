package facade

import "fmt"

/*
	外观模式：通过创建一个统一的类，用来包装子系统中一个或多个复杂的类，客户端可以通过调用外观类的方法来调用内部子系统中所有方法
	例如：智能家居系统、起床之后要开灯，开空调，开电视，打开窗帘等。睡觉的时候要关灯，关电视，关空调，拉窗帘等
    如果没有智能家居控制器，那么要手动操作所有的过程，有了智能家居控制器，就只需要控制智能家居控制器即可。
*/

// 灯
type SubSystemLight struct {
}

func (s *SubSystemLight) on() {
	fmt.Println("开灯")
}
func (s *SubSystemLight) off() {
	fmt.Println("关灯")
}

// 电视
type SubSystemTelevision struct {
}

func (s *SubSystemTelevision) on() {
	fmt.Println("开电视")
}
func (s *SubSystemTelevision) off() {
	fmt.Println("关电视")
}

// 空调
type SubSystemAirCondition struct {
}

func (s *SubSystemAirCondition) on() {
	fmt.Println("开空调")
}
func (s *SubSystemAirCondition) off() {
	fmt.Println("关空调")
}

type Facade struct {
	light        *SubSystemLight
	television   *SubSystemTelevision
	airCondition *SubSystemAirCondition
}

func (f *Facade) Facade(light *SubSystemLight, television *SubSystemTelevision, airCondition *SubSystemAirCondition) {
	f.light = light
	f.television = television
	f.airCondition = airCondition
}

func (f *Facade) Sleep() {
	fmt.Println("睡觉")
	f.light.off()
	f.television.off()
	f.airCondition.off()
}

func (f *Facade) WakeUp() {
	fmt.Println("起床")
	f.light.on()
	f.television.on()
	f.airCondition.on()
}
