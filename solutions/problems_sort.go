package solutions

import "fmt"

//一些排序算法相关的问题


//数组小和问题
//时间复杂度O(N^2) ，空间复杂度O(1)
//一个一个找前面的数比当前的数小
func SmallSumArray(arr []int) int {
	sum := 0
	for i := 1; i < len(arr); i++{
		for j := 0; j < i; j++{
			if arr[j] < arr[i] {
				sum += arr[j]  //只有小的，才相加
			}
		}
	}
	return sum
}


//使用归并思想
//一批一批找比当前大的数 ， 时间复杂度降低到了O(N*logN)级别
//func SmallSumNumMerge(arr []int) int {
//	if arr == nil || len(arr) < 2 {
//		return 0
//	}
//	return mergeSort(arr,0,len(arr)-1)
//}
//
//func mergeSort(arr []int, L int, R int) int {
//	if L == R {
//		return 0
//	}
//	mid := L + ((R-L) >> 1)
//	return  mergeSort(arr, L, mid) + mergeSort(arr, mid + 1, R) + merge(arr, L, mid, R)
//}
//
//func merge(arr []int, L int, mid int, R int) int {
//	helpArr := make([]int, R-L+1)  //申请临时辅助数组
//	p1 := L
//	p2 := mid+1
//	i := 0
//	res := 0
//	for ; p1 <= mid && p2 <= R; {
//		if arr[p1] <= arr[p2] {
//			res += (R-L+1) * arr[p1]
//		}else {
//			res += 0
//		}
//
//		if arr[p1] <= arr[p2] {
//			helpArr[i] = arr[p1]
//			p1 ++
//			i ++
//		} else if arr[p1] > arr[p2] {
//			helpArr[i] = arr[p2]
//			p2 ++
//			i ++
//		}
//	}
//
//	for ; p1 <= mid; { //p2已经越界
//		helpArr[i] = arr[p1]
//		p1 ++
//		i ++
//	}
//
//	for ; p2 <= R; { //p2已经越界
//		helpArr[i] = arr[p2]
//		p2 ++
//		i ++
//	}
//
//	//把临时数组的数copy到arr中，让arr局部有序
//	for i := 0; i < len(helpArr); i ++ {
//		arr[L+i] = helpArr[i]
//	}
//
//	return res
//}


//逆序对问题
func PrintAllReverseNum(arr []int) {
	for i := 0; i < len(arr); i ++ {
		for j := 0; j < i; j ++ {
			if arr[j] > arr[i] {
				fmt.Print(arr[j], " ", arr[i], "      ")
			}
		}
	}
}
