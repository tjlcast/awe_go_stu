package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {				// ctrl + d to stop. 
		word := input.Text()
		//fmt.Println("word:" + word)
		counts[word]++
	}

	for line, n := range counts {
		if n >= 1 {
			fmt.Printf("%s\t%d\n", line, n)
		}
	}
}
