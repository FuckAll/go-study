package decimal

import (
	"fmt"
	"strconv"
	"strings"
)

/*
十进制：D(Decimal)
二进制：B(Binary)
八进制：O(Octal)
十六进制：H(Hexadecimal)
*/

/*
	输入一个10进制的数和一个想要转换的进制，输出转换后的结果
*/
func decimalChange(k, n int) {
	var a []int
	// var n, k int

	// fmt.Printf("输入一个十进制的数字：\n")
	// fmt.Scanln(&k)
	// fmt.Printf("输入要转化的进制：\n")
	// fmt.Scanln(&n)

	for k != 0 {
		a = append(a, k%n)
		k = k / n
	}

	for from, to := 0, len(a)-1; from < to; from, to = from+1, to-1 {
		a[from], a[to] = a[to], a[from]
	}

	// count := len(a)
	// mid := count / 2
	// for i := 0; i < mid; i++ {
	// 	a[i], a[count-1] = a[count-1], a[i]
	// 	count--
	// }

	println("结果是：")
	for _, d := range a {
		switch d {
		case 10:
			print("A")
		case 11:
			print("B")
		case 12:
			print("C")
		case 13:
			print("D")
		case 14:
			print("E")
		case 15:
			print("F")
		default:
			print(d)
		}
	}
}

/*
   IP地址转换成整数，整数转换成点分十进制
*/

// IntIP2DottedDecimal convert Integer IP
func IntIP2DottedDecimal(ipInt int64) string {
	return fmt.Sprintf("%d.%d.%d.%d", ipInt>>24&0xFF, ipInt>>16&0xFF, ipInt>>8&0xFF, ipInt&0xFF)
}

// DottedDecimal2IntIP convert Dotted Decimal IP to Interger IP
func DottedDecimal2IntIP(ipStr string) int64 {
	var retInt int64
	ipPst := strings.Split(ipStr, ".")
	for k, v := range ipPst {
		tmp, _ := strconv.Atoi(v)
		retInt = retInt + int64(tmp<<(uint(3-k)*8))
	}
	return retInt
}

/*
	翻转字符串
*/

func reserveString(s string) string {
	if len(s) < 1 {
		return ""
	}

	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}
