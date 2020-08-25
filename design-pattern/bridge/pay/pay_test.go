package pay

import (
	"github.com/go-study/src/design-pattern/bridge/pay/channel"
	"github.com/go-study/src/design-pattern/bridge/pay/mode"
	"testing"
)

func TestPay(t *testing.T) {
	var pay channel.Pay
	// 微信
	pay = new(channel.WxPay)
	pay.SetPayMode(new(mode.FingerPrintMode))
	pay.Transfer("weixin_1092033111", "100000109893", 100)

	// 支付宝
	pay = new(channel.ZfbPay)
	pay.SetPayMode(new(mode.CipherMode))
	pay.Transfer("jlu19dlxo111", "100000109894", 100)
}
