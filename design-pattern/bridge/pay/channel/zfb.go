package channel

import (
	"fmt"
	"github.com/go-study/src/design-pattern/bridge/pay/mode"
)

type ZfbPay struct {
	payMode mode.IPayMode
}

func (w *ZfbPay) SetPayMode(payMode mode.IPayMode) {
	w.payMode = payMode
}

func (w *ZfbPay) Transfer(uID string, tradeID string, amount float64) string {
	security := w.payMode.Security(uID)
	fmt.Printf("模拟支付宝渠道 付划账开始。uId:%s tradeId:%s amount:%f\n", uID, tradeID, amount)
	security = w.payMode.Security(uID)
	fmt.Printf("模拟支付宝渠道付控校验。uId:%s tradeId:%s security: %t\n", uID, tradeID, security)
	if !security {
		fmt.Printf("模拟支付宝渠道 付划账拦截。uId:%s tradeId:%s amount: %f", uID, tradeID, amount)
		return "0001"
	}
	fmt.Printf("模拟支付宝渠道付划账成功。uId:%s tradeId:%s amount:%f\n", uID, tradeID, amount)
	return "0000"
}
