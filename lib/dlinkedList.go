package lib

import (
	"container/list"
	"fmt"
)

func NewDlinkedList(s []int) {
	l := list.New()
	for _,i := range s {
		l.PushBack(i)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
