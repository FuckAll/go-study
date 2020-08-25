package mode

import "fmt"

type FingerPrintMode struct {
}

func (c *FingerPrintMode) Security(uID string) bool {
	fmt.Println("指纹支付， 控校验环境安全")
	return true
}
