package strategy

import (
	"fmt"
	"sync"
)

/*
	策略模式：一个类的行为或其算法可以在运行时更改
*/
type Strategy interface {
	algorithm()
}

type StrategyA struct {
}

func (s *StrategyA) algorithm() {
	fmt.Println("StrategyA algorithm")
}

type StrategyB struct {
}

func (s StrategyB) algorithm() {
	fmt.Println("StrategyB algorithm")
}

var cachedStrategy = map[string]Strategy{}
var once sync.Once

func StrategyFactory(sType string) Strategy {
	once.Do(func() {
		cachedStrategy["A"] = new(StrategyA)
		cachedStrategy["B"] = new(StrategyB)
	})
	return cachedStrategy[sType]
}
