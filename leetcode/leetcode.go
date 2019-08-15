package main

import (
	"fmt"
)

//暴力法：逐个遍历一次, O(N^2)的时间性能
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++{
		for j := len(nums)-1; j > i; j -- {
			if nums[i] + nums[j] == target{
				return []int{i,j}
			}
		}
	}

	return []int{}
}


//煎饼排序
//思路：find max -> 从max反转 -> 从end反转 -> end--&&end!=start ->回到find max
func pancakeSorting(A []int) []int {
	kArr := make([]int,0)

	if A == nil || len(A) == 0{
		return kArr
	}

	if _Comparator(A) {
		return kArr
	}


	start, end := 0, len(A)-1
	maxIndex := -1
	for ;start < end ; {
		maxIndex = _FindMax(A,start, end)  // 1
		_ReverseArr(A, start, maxIndex) //2
		kArr = append(kArr, maxIndex+1)
		_ReverseArr(A, start, end) //3
		kArr = append(kArr, end+1)
		end -- //4
	}
	fmt.Println(A)
	fmt.Println(kArr)
	return kArr
}

//有序返回true
func _Comparator(A []int) bool {
	for i := len(A)-1; i > 0; i--{
		if A[i] < A[i-1]{
			return false
		}
	}
	return true
}

//指定范围内寻找最大值
func _FindMax(A []int, start, end int) int {
	if start >= end {
		return -1
	}

	maxIndex := start  //0
	for i := start; i <= end; i++{  //0
		if A[maxIndex] < A[i] {  //[0] < A[1]
			maxIndex = i
		}
	}

	return maxIndex
}

//指定范围内反转数组
func _ReverseArr(A []int, start, end int) {
	if start >= end{
		return
	}

	for; start < end; {
		tmp := A[end]
		A[end] = A[start]
		A[start] = tmp
		start++
		end--
	}
}

//树Node
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


//先局部反转，最后整体反转，得到最终效果
func invertTree(root *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}

	leftSubTree := invertTree(root.Left)
	rightSubTree := invertTree(root.Right)

	//整体反转，就是把左右子树调换一下
	root.Left = rightSubTree //反转后的右子树
	root.Right = leftSubTree //反转后的左子树
	return root
}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	PreOrder(root.Left)  //递归遍历左子树
	PreOrder(root.Right) //递归遍历右子树
}

func LevelOrder(node *TreeNode) {
	
}


func getParentIndex(i int) int {
	return (i-1)/2
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	a, b := 0, 0
	n1, n2 := 0, 0 //分别记录l1和l2的长度， 若999+999 = 顶多长度是max(n1,n2)+1，所以从变量中取个十百位也顶多是n1+n2+1次

	for i := 1 ; l1 != nil ; i *= 10 {
		n1 ++
		a = a + l1.Val*i
		l1 = l1.Next
	}

	for i := 1 ; l2 != nil ; i *= 10 {
		n2 ++
		b = b + l2.Val*i
		l2 = l2.Next
	}

	sum := a + b

	//求出分离次数
	parseCount := 0
	if n1 > n2 {
		parseCount = n1
	}else{
		parseCount = n2
	}

	sumArr := make([]int, 0)
	//分离个十百千位到一个数组中
	j := 0
	for i := 1;   ; i = i * 10{
		if sum/i == 0 {
			break
		}
		if j >= parseCount {
			if sum/i%10 != 0 { //长度可能大于n1或者n2
				sumArr = append(sumArr, sum/i%10)
				parseCount ++
			}
			break
		}
		sumArr = append(sumArr, sum/i%10)
		j++
	}



	//然后把数组中的元素插入到新的链表中

	var cur, head, next *ListNode
	for i := 1; i < len(sumArr); i++ {
		cur = &ListNode{sumArr[i-1], nil}
		if i == 1 {
			head = cur
		}
		next = cur
		next = &ListNode{sumArr[i], nil}
		cur.Next = next
	}

	for; head != nil ; {
		fmt.Print(head.Val, " ")
		head = head.Next
	}

	return head
}






func main(){
	//arr := []int{2,7,11,15}
	//fmt.Println(twoSum(arr, 9))

	//arr := []int{1,2,3,4,5,6}
	//_ReverseArr(arr,0, len(arr)-2)
	//fmt.Println(arr)
	//fmt.Println(_FindMax(arr,0, len(arr)-1))

	//pancakeSorting([]int{3,2,4,1})

	//root := &TreeNode{4,nil,nil}
	//root.Left = &TreeNode{2,nil,nil}
	//root.Left.Left = &TreeNode{1,nil,nil}
	//root.Left.Right = &TreeNode{3,nil,nil}
	//root.Right = &TreeNode{7,nil,nil}
	//root.Right.Left = &TreeNode{6,nil,nil}
	//root.Right.Right = &TreeNode{9,nil,nil}
	//
	//PreOrder(root)
	//fmt.Println()

	head1 := &ListNode{9,nil}
	head1.Next = &ListNode{9,nil}
	head1.Next.Next = &ListNode{9,nil}

	head2 := &ListNode{9,nil}
	head2.Next = &ListNode{9,nil}
	head2.Next.Next = &ListNode{9,nil}

	addTwoNumbers(head1, head2)

}
