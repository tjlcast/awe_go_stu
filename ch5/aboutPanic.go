package main

import "fmt"

/**
	一般而言，当panic异常发生时，程序会中断运行，并立即执行在该goroutine中被延迟的函数(defer 机制)

	对于每个goroutine，日志信息中都会有与之相对的，发生panic时的函数调用堆栈跟踪信息。
 */

func printNums(nums []int) {
	for idx, val := range nums {
		fmt.Printf("idx: %d\tval: %d\n", idx, val)
	}
}

func addOneToNums(nums []int) {
	for idx, _ := range nums {
		nums[idx] += 1
	}
}

/**
	在golang中参数为数组时，会进行直接创建一个数值相同的数组副本。
 */
func addOneToNumsPro(nums [5]int) {
	for idx, _ := range nums {
		nums[idx] += 1
	}
}

 func main() {
 	var nums = [5]int{1, 2, 3, 4, 5}

 	fmt.Printf("old nums ----->\n ")
 	printNums(nums[:])

 	fmt.Printf("old nums ----->\n ")
 	addOneToNumsPro(nums)

 	fmt.Printf("old nums ----->\n ")
 	printNums(nums[:])
 }


 func soleTitle() {
 	//type bailout struct{}
	//
 	//defer func() {
 	//	switch p := recover(); p {
	//	case nil: 			// no panic
	//	case bailout{}		// "expected" panic
	//		err = fmt.Errorf("multiple title elements")
	//	default:
	//		panic(p)		// unexpected panic; carry on panicking.
	//	}
	//}
 }
