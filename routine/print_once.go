package routine

import (
	"fmt"
	"sync"
	"time"
)

var set = make(map[int]bool)
//当一个协程调用了 Lock() 方法时，其他协程被阻塞了，直到Unlock()调用将锁释放。因此被包裹部分的代码就能够避免冲突，实现互斥。
var m sync.Mutex

func PrintOnce() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			if _,exist := set[i]; !exist {
				fmt.Println(i)
			}
			set[i] = true
			m.Unlock()
		}(100)
	}
	time.Sleep(time.Second)
}
