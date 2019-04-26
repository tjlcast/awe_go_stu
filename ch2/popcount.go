package main

/**
把一个整数减去1，再和原整数做与运算，会把该整数最右边一个1变成0。那么一个整数的二进制表示中有多少个1，就可以进行多少次这样的操作。
 */

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	// 一种dp思想
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	// byte 函数会截取末尾8位
	return int(
		pc[byte(x >> (0 * 8))] +
			pc[byte(x >> (1 * 8))] +
			pc[byte(x >> (2 * 8))] +
			pc[byte(x >> (3 * 8))] +
			pc[byte(x >> (4 * 8))] +
			pc[byte(x >> (5 * 8))] +
			pc[byte(x >> (6 * 8))] +
			pc[byte(x >> (7 * 8))])
}

var pc1 [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
} ()

func main() {
}
