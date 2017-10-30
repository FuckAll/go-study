package algorithms

import (
	"fmt"
	"unsafe"
)

/*
	中国象棋有个规则,"将帅不得见面".现在棋盘上只有一个将,一个帅,并且帅这边还有一个仕, 请写一个程序,输出棋盘上所有可能存在的合理残局情况.注:需要当场写程序并且调试通过.
*/

/*
将，帅，仕只可以走田字格
0  1  2
3  4  5
6  7  8
*/

// Chess chess
func Chess() {
	// i->将 j->帅 k->仕
	// var i,j,k int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k += 2 {
				// 帅和仕不能在一个位置
				if i == j {
					continue
				}

				// 将和帅只要不在一个直线上
				if (i % 3) != (j % 3) {
					fmt.Printf("将：%d, 帅：%d， 仕：%d \n", i, j, k)
				}

				// 将和帅在一条直线上，并且仕在帅上面
				if (j%3 == k%3) && (k < j) {
					fmt.Printf("将：%d, 帅：%d， 仕：%d \n", i, j, k)
				}
			}
		}
	}
}

const INT_SIZE int = int(unsafe.Sizeof(0))

// CheckBigEndian 检查是大端还是小端
func CheckBigEndian() {
	var i int = 0x1
	bs := (*[INT_SIZE]byte)(unsafe.Pointer(&i))
	if bs[0] == 0 {
		fmt.Println("system edian is little endian")
	} else {
		fmt.Println("system edian is big endian")
	}

	fmt.Println("This is bs:", bs)
}

/*
定一个用字符串表示的十进制大整数,求它除以d(d<10^9)的余数r.
*/

// BigIntDiv 除数
func BigIntDiv(BigInt string, d int) int {
	r := 0
	for _, s := range BigInt {
		r = (r*10 + int(s-'0')) % d
	}
	return r
}

/**
 * 在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
 * 请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
 * <p/>
 * 规律：首先选取数组中右上角的数字。如果该数字等于要查找的数字，查找过程结束：
 * 如果该数字大于要查找的数字，剔除这个数字所在的列：如果该数字小于要查找的数字，剔除这个数字所在的行。
 * 也就是说如果要查找的数字不在数组的右上角，则每－次都在数组的查找范围中剔除）行或者一列，这样每一步都可以缩小
 * 查找的范围，直到找到要查找的数字，或者查找范围为空。
 *
 * @param matrix 待查找的数组
 * @param number 要查找的数
 * @return 查找结果，true找到，false没有找到
 */

func findNumber(a [][]int, b int) bool {
	rows := len(a)
	cols := len(a[0])

	row, col := 0, cols-1

	for row >= 0 && row < rows && col >= 0 && col < cols {
		if b == a[row][col] {
			return true
		}
		if b > a[row][col] {
			row++
		}
		if b < a[row][col] {
			col--
		}
	}
	return false
}

/**
 * 请实现一个函数，把字符串中的每个空格替换成"%20"，例如“We are happy.“，则输出”We%20are%20happy.“。
 *
 * @param string     要转换的字符数组
 * @param usedLength 已经字符数组中已经使用的长度
 * @return 转换后使用的字符长度，-1表示处理失败
 */

func replaceBlack(s string) string {
	runes, temps := []rune(s), []rune{}
	if len(runes) < 1 {
		return ""
	}

	for _, v := range runes {
		if v == ' ' {
			temps = append(temps, []rune("%20")...)
		} else {
			temps = append(temps, v)
		}

	}

	return string(temps)

}

/*

	斐波那契数列(用空间换取时间)
*/
const MAX = 40

var fibs = [MAX]int{0, 1}

func fibonacci(n int) (ret int) {

	if n >= 1 {
		if fibs[n] != 0 {
			ret = fibs[n]
		} else {
			ret = fibonacci(n-1) + fibonacci(n-2)
		}
		fibs[n] = ret
	} else {
		ret = 0
	}
	return
}

/**
 * 请实现一个函数， 输入一个整数，输出该数二进制表示中1的个数。
 * 例如把9表示成二进制是1001 ，有2位是1. 因此如果输入9，该出2。
 *
 * @param n 待的数字
 * @return 数字中二进制表表的1的数目
 */
// 方法一
func numberOfOne(n int) (ret int) {
	if n == 0 {
		ret = 0
	}
	for i := 0; i < 31; i++ {
		if (n & 1) == 1 {
			ret++
		}
		n = n >> 1
	}
	return
}

// 方法二
func numberOfOne2(n int) (ret int) {
	if n == 0 {
		ret = 0
	}

	for n != 0 {
		if n%2 != 0 {
			ret++
		}
		n = n / 2
	}

	return
}

/**
 * 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，
 * 使得所有奇数位于数组的前半部分，所有偶数位予数组的后半部分。
 *
 * @param arr 输入的数组
 */
func reorderOdd(values []int) []int {
	if len(values) < 1 {
		return nil
	}

	i, j := 0, len(values)-1
	for {
		for values[i]%2 == 1 && i < j {
			i++
		}

		for values[j]%2 == 0 && i < j {
			j--
		}

		if i >= j {
			break
		}
		values[i], values[j] = values[j], values[i]
	}

	return values
}
