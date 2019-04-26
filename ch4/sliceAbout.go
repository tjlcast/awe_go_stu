package ch4


/**
	数字的初始化	--->	a := [...]int{0, 1, 2, 3, 4, 5}
	切片的初始化	--->	s := []int{0, 1, 2, 3, 4, 5} 自动创建一个匿名的数值
 */

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

