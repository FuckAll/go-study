package adapter_test

import (
	"github.com/go-study/src/design-pattern/adapter"
	"testing"
)

/*
	适配器模式：让原来不兼容的两个几口系统工作
	分类：对象适配器、类适配器、接口适配器（接口中有多个方法，只需要实现需要的方法即可）
	例如：充电器有三项充电口和二项充电口
*/

func TestAdapter(t *testing.T) {
	t2 := new(adapter.TwoToThreeAdapter)
	t2.TwoToThreeAdapter(new(adapter.TwoPower))
	t2.PowerByThree()

	t3 := new(adapter.TwoToThreeAdapterClass)
	t3.TwoToThreeAdapterClass(new(adapter.TwoPower))
	t3.PowerByThree()
}
