package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

/**
	注意：
	1、ReadFile 函数返回一个字节切片(byte slice)，必须把它转换为 string
	2、bufio.Scanner 、 ioutil.ReadFile 和 ioutil.WriteFile 都使用 *os.File 的 Read 和 Write 方法

 */

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprint(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for key, val := range counts {
		fmt.Print("%s\t%d\n", key, val)
	}
}
