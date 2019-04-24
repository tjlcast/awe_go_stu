package main

import (
	"strings"
	"os"
	"fmt"
)

func main() {
	var seperator = " "
	str := strings.Join(os.Args[1:], seperator)
	fmt.Println(str)
}
