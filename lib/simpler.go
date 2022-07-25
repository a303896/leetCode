package lib

import (
	"fmt"
	"strconv"
)

type Simpler interface {
	Get()
	Set()
}

type Simple struct {
	Value int
}

func (obj *Simple) Get() int {
	return obj.Value
}

func (obj *Simple) Set(val int) {
	obj.Value = val
}

type RSimpler interface {
	Simpler
	String()
}

type RSimple struct {
	Simple
}

func (obj *RSimple) String() string {
	return strconv.Itoa(obj.Value)
}

func Type(i interface{}) {
	if t,ok := i.(RSimple); ok {
		fmt.Printf("%T.Value is %v \n", t, t)
		return
	}
	if t,ok := i.(*RSimple); ok {
		fmt.Printf("%T.Value is %v \n", t, t)
		return
	}
	fmt.Printf("unknown \n")
}

func Fi(i interface{}) {
	switch t := i.(type) {
	case Simple:
		fmt.Printf("%T.Value is %v \n", t, t)
	case RSimple:
		fmt.Printf("%T.Value is %v \n", t, t)
	case *Simple:
		fmt.Printf("%T.Value is %v \n", t, t)
	case *RSimple:
		fmt.Printf("%T.Value is %v \n", t, t)
	case nil:
		fmt.Printf("i is nil \n")
	default:
		fmt.Printf("unknown \n")
	}
}

//example
/*
	s := new(lib.Simple)
	rs := new(lib.RSimple)
	lib.Type(s)			//output: unknown
	lib.Type(rs)		//output: *lib.RSimple.Value is 0
	so := lib.Simple{Value: 1}
	rso := lib.RSimple{Simple: lib.Simple{Value: 1}}
	lib.Type(so)		//output: unknown
	lib.Type(rso)		//output: lib.RSimple.Value is {{1}}
	lib.Fi(rs)			//output: *lib.RSimple.Value is 0
	lib.Fi(s)			//output: *lib.Simple.Value is &{0}
	lib.Fi(so)			//output: lib.Simple.Value is {1}
	lib.Fi(rso)			//output: lib.RSimple.Value is {{1}}
 */