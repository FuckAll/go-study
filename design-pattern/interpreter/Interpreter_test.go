package interpreter

import (
	"fmt"
	"testing"
)

/*
	解释器模式：提供了评估语言的语法或表达式的方式，它属于行为型模式。这种模式实现了一个表达式接口，该接口解释一个特定的上下文。这种模式被用在 SQL 解析、符号处理引擎等。
*/

func TestInterpreter(t *testing.T) {
	zhangSan := NewTerminalExpression("张三")
	liSi := NewTerminalExpression("李四")
	isMale := NewOrExpression(zhangSan, liSi)

	nana := NewTerminalExpression("娜娜")
	married := NewTerminalExpression("结婚")
	isMarriedWoman := NewAndExpression(nana, married)

	fmt.Printf("张三是一名男性？:%t\n", isMale.interpret("张三"))
	fmt.Printf("娜娜是一名结婚女性？:%t\n", isMarriedWoman.interpret("娜娜 结婚"))

}
