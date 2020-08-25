package mode

import "fmt"

type CipherMode struct {
}

func (c *CipherMode) Security(uID string) bool {
	fmt.Println("密码支付， 控校验环境安全")
	return true
}
