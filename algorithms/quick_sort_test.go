package algorithms

import "testing"
import "fmt"

func TestQuickSort(t *testing.T) {
	// values := []int{3, 4, 2, 6, 7, 11, 16, 15, 18, 10}

	values := []int{3, 4, 2, 6, 3, 2}
	quickSort(values, 0, len(values)-1)
	fmt.Println("快速排序结果：", values)
}
