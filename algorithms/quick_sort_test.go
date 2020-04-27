package algorithms

import "testing"
import "fmt"

func TestQuickSort(t *testing.T) {
	// values := []int{3, 4, 2, 6, 7, 11, 16, 15, 18, 10}

	values := []int{3, 4, 2, 6, 3, 2}
	quickSort(values, 0, len(values)-1)
	fmt.Println("快速排序结果：", values)
}

func TestQuickSort02(t *testing.T) {
	// values := []int{3, 4, 2, 6, 7, 11, 16, 15, 18, 10}

	values := []int{3, 4, 2, 6, 3, 2}
	quickSort02(values, 0, len(values)-1)
	fmt.Println("快速排序结果：", values)
}

func quickSort02(values []int, left, right int) {
	if left < right {
		guard := values[left]
		i, j := left, right
		for {
			// 从右向左找到，比哨兵大的值
			for values[j] <= guard && i < j {
				j--
			}
			// 从左向右， 比哨兵小的值
			for values[i] >= guard && i < j {
				i++
			}
			// 退出条件(i是中心点)
			if i >= j {
				break
			}
			// 交换数据
			values[i], values[j] = values[j], values[i]
		}

		// 交换哨兵位置
		values[left] = values[i]
		values[i] = guard

		// 递归
		quickSort(values, left, i-1)
		quickSort(values, i+1, right)
	}
}
