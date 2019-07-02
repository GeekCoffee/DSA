package main

// O(N^2)
//冒泡排序 - 每一轮后有效问题规模都会减少
//时间复杂度估计： N + N-1 + N-2 + N-3 + ... + 2 + 1 = 等差数列
//关键点：改变终点，不改变起点
func BubbleSort(arr []int) {
	//step1：极端例子或者边界处理、异常处理
	if arr == nil || len(arr) <= 1 {
		return
	}

	n := len(arr)

	//step2：正常逻辑处理
	// 2 5 1 8 0 4
	// i
	//   j
	//写法一
	//for i := 0; i < n; i ++ {
	//	for j := 1; j < n-i; j++ {  //改变终点，起点不变
	//		//对比相邻元素的大小，若使用j+1可能会出现越界异常
	//		if arr[j-1] > arr[j]{
	//			swap(&arr[j-1], &arr[j])
	//		}
	//	}
	//}

	//写法二
	//这样写i与i+1就不会有越界情况了
	// N + N-1 + N-2 + N-3 + ... + 2 + 1 = 等差数列 = (N+1)*N/2 ~ O(N^2)
	for end := n-1; end > 0 ; end-- {
		for i := 0; i < end; i++ {
			if arr[i] > arr[i+1]{
				swap(&arr[i], &arr[i+1])
			}
		}
	}
}


//交换两个变量的值
func swap(a,b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}


//插入排序 - 在工程上还有使用，因为比较智能，不像冒泡和选择，流程已经定死了
//当数据几乎有序的情况下，时间复杂度是非常低的，时间性能是非常好的
//关键点：使用j+1和j去逐步完成swap操作，交换到底的情况
func InsertionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	// j 和 j+1比较，然后j和j+1都向前移动
	for i := 1; i < len(arr); i++ {
		for j := i-1; j >= 0 && arr[j+1] < arr[j]; j-- {
			swap(&arr[j+1], &arr[j])
		}
	}

}


//选择排序
//关键点：每找完一个最小值元素，都要与i位置的元素swap，且把minIndex归位到有效问题规模的第一个元素位置上
func SelectionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	n := len(arr)

	//选择有效问题规模的最小值
	for i := 0; i < n-1; i ++ {  //当i指向最后一个元素的时候，有效问题规模已经为1，即也没有无序或者有序这样一种比较规则了
		minIndex := i  //每找完一个最小值，都要把minIndex归位到i所指的位置上
		for j := i; j < n; j++ {  //改变起点，不改变终点
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		//找到最小值，与i所指的位置相换
		swap(&arr[i], &arr[minIndex])
	}


}


//O(N * logN)
//归并排序
func MergeSort(arr []int) {

}

//快速排序-2路
func QuickSort(arr []int) {

}


//O(c) + 空间复杂度O(N)
//计数排序
func BasicSort(arr []int) {

}



func main(){

	//统一测试规范
	TestSort(BubbleSort, 10, 100, 500000)
	TestSort(InsertionSort, 10, 100, 500000)
	TestSort(SelectionSort, 10, 100, 500000)



	//测试排序算法的时间性能
	//arr := GenerateRandomArrayFixed(30000, 1024)
	//fmt.Println(arr)
	//t := TestSortTime(BubbleSort, arr)
	//fmt.Println(arr)
	//fmt.Println(t, " ms")

}


