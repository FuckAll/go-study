package adapter

import "fmt"

/*
	适配器模式：让原来不兼容的两个几口系统工作
	分类：对象适配器、类适配器、接口适配器（接口中有多个方法，只需要实现需要的方法即可）
	例如：充电器有三项充电口和二项充电口
*/

type TwoPower struct {
}

func (t *TwoPower) powerByTwo() {
	fmt.Println("提供二项供电")
}

type IThreePower interface {
	PowerByThree()
}

// 对象适配器采用组合的方式
type TwoToThreeAdapter struct {
	twoPower *TwoPower
}

func (t *TwoToThreeAdapter) TwoToThreeAdapter(twoPower *TwoPower) {
	t.twoPower = twoPower
}

func (t *TwoToThreeAdapter) PowerByThree() {
	fmt.Println("对象适配器：借助组合适配器转化为二项充电口")
	t.twoPower.powerByTwo()
}

// 类适配器，采用的是继承的方式
type TwoToThreeAdapterClass struct {
	*TwoPower
}

func (t *TwoToThreeAdapterClass) TwoToThreeAdapterClass(twoPower *TwoPower) {
	t.TwoPower = twoPower
}

func (t *TwoToThreeAdapterClass) PowerByThree() {
	fmt.Println("类适配器：借助组合适配器转化为二项充电口")
	t.powerByTwo()
}
