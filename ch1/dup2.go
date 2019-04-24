package main

import (
	"os"
	"fmt"
	"bufio"
)

/**
	注意：
	1、os.Stdin 和 file.Open 都是 os.File 对象
	2、make 将创建一个对象引用
 */

func main() {
	counts := make(map[string]int)
	args := os.Args[1:]

	if len(args) == 0 {
		// 从标准输入中读取
		countLines(os.Stdin, counts)
	} else {
		// 从文件中读取
		for _, arg := range args {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprint(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}