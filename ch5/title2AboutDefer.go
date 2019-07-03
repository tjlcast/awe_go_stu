package ch5

import "os"

/**
	当defer语 句被执行时，跟在defer后面的函数会被延迟执行。
	直到包含该defer语句的函数执行完毕时，defer后的函数才会被执行，
	不论包含defer语句的函数是通过return正常结束，还是由于panic导致的异常结束。

	多条defer语句，它们的执行顺序与声明顺序相反。

	defer语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放锁。
 */

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	//return ReadAll(f)
}

