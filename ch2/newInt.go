package main

/**
编译器会自动选择在栈上还是在堆上分配局部变量的存储空间，
这个选择并不是由用var还是new声明变量的方式决定的。
 */

func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var a int
	return &a
}

var global *int

func fu() {
	var x int	// 这里的x变量必须在堆上分配(有全局指针指向这个内部变量，改变量逃逸)
	x = 1
	global = &x
}

func g() {
	y := new(int)
	*y = 1
}