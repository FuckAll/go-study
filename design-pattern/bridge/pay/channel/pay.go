package channel

import "github.com/go-study/src/design-pattern/bridge/pay/mode"

type Pay interface {
	SetPayMode(payMode mode.IPayMode)
	Transfer(uID string, tradeID string, amount float64) string
}
