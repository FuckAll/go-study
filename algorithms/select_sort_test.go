package algorithms

import "testing"
import "fmt"

func TestSelectSort(t *testing.T) {
	values := []int{3, 4, 1, 6, 7, 11, 16, 15, 18, 10, 2}
	selectSort(values)
	fmt.Println("选择排序结果：", values)
}
