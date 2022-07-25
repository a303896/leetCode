package lib

import "fmt"

type Obj interface{}

func MapFunc(f func(o Obj) Obj, list []Obj) []Obj {
	for key,item := range list {
		list[key] = f(item)
	}
	return list
}

func MapFuncVar(f func(o Obj) Obj, list ...Obj) []Obj {
	for key,item := range list {
		list[key] = f(item)
	}
	return list
}

func Mutate(o Obj) Obj {
	switch o.(type) {
	case int:
		o = o.(int) * 2
	case string:
		o = o.(string) + o.(string)
	default:
		fmt.Println("wrong type obj")
	}
	return o
}

//example
/*
	list := []lib.Obj {"He","Ha",54,100,"嘻",50}
	fmt.Println(list)			//output: [He Ha 54 100 嘻 50]
	list = lib.MapFunc(lib.Mutate, list)
	fmt.Println(list)			//output: [HeHe HaHa 108 200 嘻嘻 100]
	list = lib.MapFuncVar(lib.Mutate, "He","Ha",54,100,"嘻",50)
	fmt.Println(list)			//output: [HeHe HaHa 108 200 嘻嘻 100]
 */