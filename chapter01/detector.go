package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"time"
)

//对数器，作用是快速验证算法的正确性
//对数器包括数组对数器、二叉树对数器等
//以后还会添加计算算法流程的时间工具


//生成一个长度随机，且值也随机的数组
func GenerateRandomArray(maxSize, maxValue int) []int {
	rand.Seed(int64(maxSize) + time.Now().Unix())  //设置随机种子
	arr := make([]int, rand.Intn(maxSize))  //随机make一个长度为n的数组
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue)
	}
	return arr
}


//生成一个固定长度是随机数组
func GenerateRandomArrayFixed(size, maxValue int) []int {
	rand.Seed(int64(size) + time.Now().Unix())  //设置随机种子
	arr := make([]int, rand.Intn(size))  //随机make一个长度为n的数组
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue)
	}
	return arr
}

//copy整型数组
func CopyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	arr2 := make([]int, len(arr))  //长度一样
	for i := range arr { //数据一样
		arr2[i] = arr[i]
	}
	return arr2
}


//判断两个数组是否相等
//不仅数组长度要相等，而且数据顺序要相等
func IsEqual(arr1, arr2 []int) bool {
	//判空操作
	if (arr1 == nil && arr2 != nil) || (arr1 != nil && arr2 == nil) {
		return false
	}
	if arr1 == nil && arr2 == nil {
		return false
	}
	if len(arr1) != len(arr2) {
		return false
	}
	//判断两个数组的每个元素是否相等
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i]{
			return false
		}
	}
	//否则返回true
	return true
}


//一个绝对正确的算法
func RightSort(arr []int){
	sort.Ints(arr)  //golang内置排序算法，对整型数组的排序
	//当然还有其他的自定义排序
}


func TestSortTime(f func([]int), arr []int) int64 {
	start := time.Now().UnixNano()
	f(arr)  //传入哪个函数就使用哪个函数
	end := time.Now().UnixNano()
	t := (end - start) / 1000000
	return t
}


func TestSort(f func([]int), maxSize int, maxValue int, opsTime int) {

	succeed := true //是否排序成功标志

	for i := 0; i < opsTime; i ++ {
		arr1 := GenerateRandomArray(maxSize, maxValue)
		arr2 := CopyArray(arr1)
		f(arr1) //传入哪个函数就使用哪个函数
		RightSort(arr2)
		if !IsEqual(arr1, arr2){
			succeed = false
			break
		}
	}

	funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	if succeed {
		fmt.Println("Succeed ,sort method: ", funcName)
	}else{
		fmt.Println("test failure!!!! , sort method: ", funcName)
	}

	//打印效果
	arr := GenerateRandomArray(maxSize, maxValue)
	fmt.Println(arr)
	f(arr)
	fmt.Println(arr)
}
