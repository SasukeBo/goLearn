/*
默认选择

当select中的其它分支都没有准备好时，default分支就会执行。

为了在尝试发送或者接受时不发生阻塞，可使用default分支：

select {
case i := <-c:
	// 使用 i
default:
	// 从 c 中接收会阻塞时执行
}
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
