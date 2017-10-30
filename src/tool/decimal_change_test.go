package decimal

import "testing"
import "fmt"

/*
	输入一个10进制的数和一个想要转换的进制，输出转换后的结果
*/
func TestDecimalChange(t *testing.T) {
	decimalChange(100, 6)
}

func TestReserveString(t *testing.T) {
	fmt.Println("翻转字符串之后的结果为：", reserveString("g1od"))
}
