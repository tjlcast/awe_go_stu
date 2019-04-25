package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

/**
	tips:
	1、http.get 返回请求结果，其中的body字段是一个可读的服务器响应流。
	2、ioutil.ReadAll函数从response中读取到全部内容。
 */

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", b)
	}
}
