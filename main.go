package Interview

import (
	"./sort"
)

func main(){

	maxSize, maxVal, opsTime := 10, 100, 500000

	////统一测试规范
	////TestSort(BubbleSort, 10, 100, 50000)
	////TestSort(InsertionSort, 10, 100, 50000)
	////TestSort(SelectionSort, 10, 100, 50000)
	//sort.TestSort(sort.MergeSortRecursive, maxSize, maxVal, opsTime)
	//sort.TestSort(sort.QuickSortOfRandom3Part, maxSize, maxVal, opsTime)
	//sort.TestSort(sort.MaxHeapSort, maxSize, maxVal, opsTime)
	sort.TestSort(sort.MinHeapSort, maxSize, maxVal, opsTime)
	//arr := []int{1,3,4,2,5}
	//fmt.Println(solutions.SmallSumArray(arr))

	//solutions.PrintAllReverseNum(arr)



	//fmt.Println(5/2)


	//测试排序算法的时间性能
	//arr := GenerateRandomArrayFixed(30000, 1024)
	//fmt.Println(arr)
	//t := TestSortTime(BubbleSort, arr)
	//fmt.Println(arr)
	//fmt.Println(t, " ms")

	//arr := []int{5,4,1,7,8,0}
	//sort.HeapSort(arr)
	//fmt.Println(arr)

}