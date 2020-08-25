package mode

import "fmt"

type FaceMode struct {
}

func (c *FaceMode) Security(uID string) bool {
	fmt.Println("人脸支付， 控校验环境安全")
	return true
}
