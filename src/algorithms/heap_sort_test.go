package algorithms

import "testing"
import "fmt"

func TestBuildHeap(t *testing.T) {
	values := []int{3, 4, 2, 6, 7, 11, 16, 15, 18, 10}
	BuildHeap(values, len(values))
	fmt.Println("大堆构建之后的结果：", values)
}

func TestHeapSort(t *testing.T) {
	values := []int{3, 4, 2, 6, 7, 11, 16, 15, 18, 10}
	ret := HeapSort(values)
	fmt.Println("堆排序之后：", ret)
}
