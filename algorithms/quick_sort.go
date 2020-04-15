package algorithms

import (
	"fmt"
)

/**
 *@Description:简单快速排序算法的实现（从大到小）
 *@author izgnod
 *@time 2017-10-11 下午22:29
 */
func quickSort(values []int, left, right int) {
	if left < right {
		// 设置基准值
		temp := values[left]

		i, j := left, right
		for {
			// 从右向左，找到第一个比基准值小的数
			for values[j] <= temp && i < j {
				j--
			}

			// 从左向右，找到第一个比基准值大的数
			for values[i] >= temp && i < j {
				i++
			}

			// 哨兵相遇， 退出循环
			if i >= j {
				break
			}

			// 交换左右两侧的值
			values[i], values[j] = values[j], values[i]
		}

		// 将基准点移动到哨兵相遇点
		values[left] = values[i]
		values[i] = temp

		fmt.Println("快排序结果：", values)

		// 递归，左右两侧分别进行排序
		quickSort(values, left, i-1)
		quickSort(values, i+1, right)
	}

}
