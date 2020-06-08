package interpreter

import "strings"

/*
	解释器模式：提供了评估语言的语法或表达式的方式，它属于行为型模式。这种模式实现了一个表达式接口，该接口解释一个特定的上下文。这种模式被用在 SQL 解析、符号处理引擎等。
*/

type Expression interface {
	interpret(context string) bool
}

type TerminalExpression struct {
	data string
}

func NewTerminalExpression(data string) *TerminalExpression {
	return &TerminalExpression{data: data}
}

func (t *TerminalExpression) interpret(context string) bool {
	if strings.Contains(context, t.data) {
		return true
	}
	return false
}

type OrExpression struct {
	exp1 Expression
	exp2 Expression
}

func (o *OrExpression) interpret(context string) bool {
	return o.exp1.interpret(context) || o.exp2.interpret(context)
}

func NewOrExpression(exp1 Expression, exp2 Expression) *OrExpression {
	return &OrExpression{exp1: exp1, exp2: exp2}
}

type AndExpression struct {
	exp1 Expression
	exp2 Expression
}

func (a *AndExpression) interpret(context string) bool {
	return a.exp1.interpret(context) && a.exp2.interpret(context)
}

func NewAndExpression(exp1 Expression, exp2 Expression) *AndExpression {
	return &AndExpression{exp1: exp1, exp2: exp2}
}
