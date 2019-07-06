package sort

import (
	"math/rand"
	"time"
)



//master公式 => O(N * logN)
//归并排序 - 关键：外部两数组对比排序
func MergeSortRecursive(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	preProcessMerge(arr, 0, len(arr)-1)  //从0到len-1的范围做分治，都是闭区间[0,len(arr)-1]
}

//merge之前的预处理
func preProcessMerge(arr []int, L int, R int) {
	if L == R {
		return  //递归基，什么都不做
	}

	//避免数太大，整型变量溢出，所以做一次偏移量操作，整除向下取整
	//CPU中位移操作需要的时钟周期数比除法器需要用的时钟周期数要少，即时间快
	mid := L + ((R-L) >> 1)
	preProcessMerge(arr, L, mid)  //先排arr的左半部分
	preProcessMerge(arr, mid+1, R)  //再排序arr的右半部分
	merge(arr, L, mid, R)  // 合并两个子数组的元素到临时子数组上，本质上是对arr的局部排序
}

//归并操作
func merge(arr []int, L int,mid int, R int) {
	helpArr := make([]int, R-L+1)  //每次开辟(R-L+1)空间的数组，子规模不同，则开辟的空间也不同
	p1, p2 := L, mid+1  //p1指向arr的L到mid区间的子数组，p2指向arr的mid+1到R区间的子数组
	i := 0
	for ; p1 <= mid && p2 <= R ; {  //相当于while循环
		if arr[p1] < arr[p2] {
			helpArr[i] = arr[p1]
			p1 ++
			i ++
		}else { //arr[p1] >= arr[p2]
			helpArr[i] = arr[p2]
			p2 ++
			i ++
		}
	}

	for ; p1 <= mid ; { //p1 <= mid,说明p2已经越界
		helpArr[i] = arr[p1]
		p1 ++
		i ++
	}

	for ; p2 <= R; { //p2 <= R，说明p1已经越界
		helpArr[i] = arr[p2]
		p2 ++
		i ++
	}

	//拷贝回原数组，让原数组局部有序到全局有序
	//注意arr是全局变量，而helpArr是局部变量，所以要改变的次数是len(helpArr)
	//所以要进行L位置的偏移
	for i := 0; i < len(helpArr); i++{
		arr[L + i] = helpArr[i]
	}

}

// 3路快排可以加速2路快排，加速点在partition的过程中添加了=等于key的区域
// 每一次遍历就可以排序得==等于这一片区域，在这个地方加速了排序过程
// 经典2路快排，每一次划分区域只能排序得一个数，就是中间那个key，然后左右两边的子数组继续递归排序
// 目前效率最好的是3路随机排序，当然也需要看数据的特点
// 关键：partition划分区域，less、more、cur区域
func QuickSortOfRandom3Part(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	quickSort(arr, 0 , len(arr)-1)  //区间选择，以下标为准
}

func quickSort(arr []int, L int, R int) {
	if L < R {
		rand.Seed(int64(R-L+1) + time.Now().Unix())  //设置随机种子
		randIndex := L + rand.Intn(R-L+1)
		swapArrElem(arr, randIndex, R)  //数组最后一个元素，与随机得到的randIndex位置上的元素交换，减少退化的可能性
		p := partition(arr,L, R)  // p只有两个元素的数组 []int{less ,more}元素
		quickSort(arr, L, p[0]-1)  //递归排序less区域
		quickSort(arr, p[1]+1, R)  //递归排序more区块
	}

}

//分治的过程
func partition(arr []int, L int, R int) []int {
	less := L - 1  //less区域
	more := R   //more区域 ，默认使用数组最后一个数做key，所以最后一个数不参与遍历排序
	cur := L  //equal区域

	for ; cur < more; { //当cur等于more时，循环结束
		if arr[cur] < arr[R] {
			less ++
			swap(&arr[less], &arr[cur])  //cur一定是要走的，可以自己画图分析一下
			cur ++
		}else if arr[cur] > arr[R] {
			more --
			swap(&arr[more], &arr[cur])
		}else { // arr[cur] == num
			cur ++
		}
	}

	swapArrElem(arr, more, R)  //把R位置上的元素归位到equal区域

	//返回一组数组，数组存储的不是具体数据，而是存储equal区域的头和尾元素的下标
	//闭区间
	return []int{less+1, more}
}


//堆排序 - O(N * logN)
//建堆过程 - O(N) = log1 + log2 + log3 + ... + logN-1 约等于 log (N^N) = N
//排好后，从大到小
func MinHeapSort(arr []int){
	if arr == nil || len(arr) < 2{
		return
	}

	for i := 0 ; i < len(arr); i ++{  //建立大根堆
		MinHeapInsertion(arr, i)
	}

	heapSize := len(arr)
	heapSize --
	swap(&arr[0], &arr[heapSize])
	for ; heapSize > 0;  {
		MinHeapify(arr, 0, heapSize)  //虽然末尾元素会暂时在最后，但是不影响整体的下溢流程
		heapSize --
		swap(&arr[0], &arr[heapSize])
	}


}


// 向完全二叉树中插入节点，本质上是在数组中不断向父节点进行swap操作
// 上溢操作
func MinHeapInsertion(arr []int, i int) {
	for ; arr[i] < arr[(i-1)/2]; {  //当父节点大于子节点的时候，父节点下溢，要保证大值数都沉在完全二叉树底部
		if arr[i] < arr[(i-1)/2] {
			swap(&arr[i], &arr[(i-1)/2])
		}
		i = (i-1)/2  //继续占据父节点的位置
	}
}


//下溢操作
func MinHeapify(arr[] int, index int, heapSize int) {
	leftChild := index * 2 + 1
	for ; leftChild < heapSize; { //当叶子节点超出heapSize范围，循环结束
		smaller := leftChild   //设置smaller默认是leftChild
		if leftChild+1 < heapSize && arr[leftChild] > arr[leftChild+1]{ //若有右孩子且比左孩子大
			smaller = leftChild+1 //右孩子节点比较小
		}

		//当父节点比孩子节点大的时候，swap
		if arr[index] > arr[smaller] {

		}else{ // 当父节点比孩子节点小的时候，不需要再进行下溢了，把smaller指针指向父节点
			smaller = index
		}

		if smaller == index { //若达到这一步，说明父节点已经比子节点小了
			break
		}

		swap(&arr[smaller], &arr[index])
		index = smaller //继续下溢，与交换节点的索引相互换,index可能在左也可能在右
		leftChild = index*2 + 1
	}

}



func MaxHeapSort(arr []int){
	if arr == nil || len(arr) < 2{
		return
	}

	for i := 0 ; i < len(arr); i ++{  //建立大根堆
		MaxHeapInsertion(arr, i)
	}

	heapSize := len(arr)
	heapSize --
	swap(&arr[0], &arr[heapSize])
	for ; heapSize > 0;  {
		MaxHeapify(arr, 0, heapSize)  //虽然末尾元素会暂时在最后，但是不影响整体的下溢流程
		heapSize --
		swap(&arr[0], &arr[heapSize])
	}


}


// 向完全二叉树中插入节点，本质上是在数组中不断向父节点进行swap操作
// 上溢操作
func MaxHeapInsertion(arr []int, i int) {
	for ; arr[i] > arr[(i-1)/2]; {  //若子节点比父节点还大的情况下，继续上溢
		if arr[i] > arr[(i-1)/2] {
			swap(&arr[i], &arr[(i-1)/2])
		}
		i = (i-1)/2  //继续占据父节点的位置
	}
}


//下溢操作
func MaxHeapify(arr[] int, index int, heapSize int) {
	leftChild := index * 2 + 1
	for ; leftChild < heapSize; { //当叶子节点超出heapSize范围，循环结束
		largest := leftChild   //设置largest默认为leftChild
		if leftChild+1 < heapSize && arr[leftChild] < arr[leftChild+1]{ //若有右孩子且比左孩子大
			largest = leftChild+1 //右孩子节点比较大
		}

		//确定largest后，与父节点比较，若父节点较大，把index赋值给largest，即让largest指向index的指向
		if arr[largest] > arr[index] {

		}else{
			largest = index
		}

		if largest == index { //若达到这一步，说明父节点已经比子节点大了
			break
		}

		swap(&arr[largest], &arr[index])
		index = largest //继续下溢，与交换节点的索引相互换,index可能在左也可能在右
		leftChild = index*2 + 1
	}

}




