package algorithms

import (
	"fmt"
)

// public static void selectSort(int[] arr) {
// 	if(arr == null || arr.length == 0)
// 		return ;
// 	int minIndex = 0;
// 	for(int i=0; i<arr.length-1; i++) { //只需要比较n-1次
// 		minIndex = i;
// 		for(int j=i+1; j<arr.length; j++) { //从i+1开始比较，因为minIndex默认为i了，i就没必要比了。
// 			if(arr[j] < arr[minIndex]) {
// 				minIndex = j;
// 			}
// 		}

// 		if(minIndex != i) { //如果minIndex不为i，说明找到了更小的值，交换之。
// 			swap(arr, i, minIndex);
// 		}
// 	}

// }

/**
 *@Description:简单选择排序算法的实现
 *@author izgnod
 *@time 2017-10-11 下午22:29
 */

func selectSort(arr []int) {
	if len(arr) < 1 {
		return
	}

	fmt.Println("排序前：", arr)

	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}

		fmt.Println("第", i, "次排序：", arr)
	}

}
