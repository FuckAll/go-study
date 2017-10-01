package simple

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
