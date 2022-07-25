package lib

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age int
}

type Admin struct {
	User
	Title string
}

type Driver struct {
	Name string
	Age int
}

type Car interface {
	String() string
}

type BMW struct {
	Title string
	Color string
	Price int
}

func (car BMW) String() string{
	str := ""
	v := reflect.ValueOf(car)
	for i := 0; i < v.NumField(); i++ {
		str += fmt.Sprintf("%+v", v.Field(i))
	}
	return str
}

func (p Driver) Drive(car Car) {
	v := reflect.ValueOf(p)
	t := reflect.TypeOf(p)
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s:%v", t.Field(i).Name, v.Field(i).Interface())
	}
	fmt.Println("drive", car.String())
}

func (user User) Info() {
	t := reflect.TypeOf(user)
	fmt.Printf("%v \n", t)
	fmt.Println("Type:", t)
	v := reflect.ValueOf(user)
	fmt.Println("Fields:")
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%s %v %v \n", f.Name, f.Type, val)
	}
	for i := 0; i < v.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%s %v \n", m.Name, m.Type)
	}
}

func GetAnonymousField(o interface{}) {
	t := reflect.TypeOf(o)
	//取struct匿名字段
	fmt.Printf("%+v \n", t.Field(0))
	//取struct匿名字段下的字段
	fmt.Printf("%+v \n", t.FieldByIndex([]int{0,1}))
}

/*
对于reflect.Values也有类似的区别。有一些reflect.Values是可取地址的；其它一些则不可以。考虑以下的声明语句：

x := 2                   // value   type    variable?
a := reflect.ValueOf(2)  // 2       int     no
b := reflect.ValueOf(x)  // 2       int     no
c := reflect.ValueOf(&x) // &x      *int    no
d := c.Elem()            // 2       int     yes (x)
其中a对应的变量不可取地址。因为a中的值仅仅是整数2的拷贝副本。b中的值也同样不可取地址。c中的值还是不可取地址，它只是一个指针&x的拷贝。
实际上，所有通过reflect.ValueOf(x)返回的reflect.Value都是不可取地址的。但是对于d，它是c的解引用方式生成的，指向另一个变量，因此是可取地址的。
我们可以通过调用reflect.ValueOf(&x).Elem()，来获取任意变量x对应的可取地址的Value。
*/