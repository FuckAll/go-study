package channel

import (
	"fmt"
	"github.com/go-study/src/design-pattern/bridge/pay/mode"
)

type WxPay struct {
	payMode mode.IPayMode
}

var _ Pay = &WxPay{}

func (w *WxPay) SetPayMode(payMode mode.IPayMode) {
	w.payMode = payMode
}

func (w *WxPay) Transfer(uID string, tradeID string, amount float64) string {
	fmt.Printf("模拟微信渠道 付划账开始。uId:%s tradeId:%s amount:%f\n", uID, tradeID, amount)
	security := w.payMode.Security(uID)
	fmt.Printf("模拟微信渠道付控校验。uId:%s tradeId:%s security: %t\n", uID, tradeID, security)
	if !security {
		fmt.Printf("模拟微信渠道 付划账拦截。uId:%s tradeId:%s amount: %f", uID, tradeID, amount)
		return "0001"
	}
	fmt.Printf("模拟微信渠道 付划账成功。uId:%s tradeId:%s amount:%f\n", uID, tradeID, amount)
	return "0000"
}
