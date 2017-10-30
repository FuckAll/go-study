package algorithms

/*
# heap Sort 堆排序
利用“堆”这种数据结构来排序，分为最大堆、最小堆排序。堆是种完全二叉树，特点是父子有序兄弟无序（可能有序可能乱序）。
排序过程即循环构建堆并依次排列堆的根节点的过程。
关键过程为建堆、堆顶与无序区尾部叶子交换
# Runtime:
- Average: O(n lg n)
- Worst: O(n lg n)
- Best: O(n lg n)
*/

// BuildHeap 构造大堆
func BuildHeap(arr []int, length int) {
	if length == 1 {
		return
	}

	//最大非叶节点偏移量（序号为length/2）
	maxBranch := length/2 - 1
	for i := maxBranch; i >= 0; i-- {
		// 右孩子偏移量
		lChild := i*2 + 1
		// 左孩子偏移量
		rChild := lChild + 1

		// 临时最大偏移量
		tmpMax := i

		// 将三个结点构造成堆（父节点最大）
		if arr[lChild] > arr[tmpMax] {
			tmpMax = lChild
		}

		if (rChild <= length-1) && (arr[rChild] > arr[tmpMax]) {
			tmpMax = rChild
		}

		if tmpMax != i {
			arr[tmpMax], arr[i] = arr[i], arr[tmpMax]
		}
	}
}

// HeapSort 堆排序
func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < len(arr); i++ {
		lastLength := length - i
		BuildHeap(arr, lastLength)

		if i < length {
			arr[0], arr[lastLength-1] = arr[lastLength-1], arr[0]
		}
	}
	return arr
}
