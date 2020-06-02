package strategy

import "testing"

func TestStrategy(t *testing.T) {
	strategyA := StrategyFactory("A")
	strategyA.algorithm()

	strategyB := StrategyFactory("B")
	strategyB.algorithm()
}
