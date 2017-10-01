package decimal

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

	count := len(a)
	mid := count / 2
	for i := 0; i < mid; i++ {
		a[i], a[count-1] = a[count-1], a[i]
		count--
	}

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
