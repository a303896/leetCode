package routine

import (
	"fmt"
	"time"
)

func TickTok() {
	tick := time.Tick(1e8)			//周期性循环发送时间
	boom := time.After(5e8)			//只发送一次时间
	for {
		select {
		case <- tick:
			fmt.Println("tick.")
		case <- boom:
			fmt.Println("boom!")
			return
		default:
			fmt.Println("     .")
			time.Sleep(5e7)
		}
	}
}
