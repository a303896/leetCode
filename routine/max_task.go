package routine

import (
	"fmt"
	"time"
)

const MAXREQS = 5

var sem = make(chan int, MAXREQS)

func process(r *Request) {
	r.a = -1
	r.b = -1
	time.Sleep(1 * time.Second)
}

func handle(r *Request) {
	fmt.Println("33333")
	sem <- 1
	fmt.Println("11111")
	process(r)
	<- sem
}

func server2(service chan *Request) {
	for {
		req := <-service
		fmt.Println("222222")
		go handle(req)
	}
}

func Start3() {
	service := make(chan *Request)
	go server2(service)
	for i := 0; i < 10; i++ {
		req := new(Request)
		req.a = i
		req.b = i + 10
		req.replyc = make(chan int)
		service <- req
	}
	//time.Sleep(20 * time.Second)
	fmt.Println("done")
}
