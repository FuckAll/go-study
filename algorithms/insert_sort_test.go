package algorithms

import "testing"
import "fmt"

func TestInsertSort(t *testing.T) {
	// values := []int{3, 4, 2, 6, 7, 11, 16, 15, 18, 10}

	values := []int{3, 4, 2, 6, 3, 2}
	InsertionSort(values)
	fmt.Println("插入排序结果：", values)
}
