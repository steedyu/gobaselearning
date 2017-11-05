package ivsample

import (
	"fmt"
	"reflect"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func EmbeddedStructMethod() {
	t := Teacher{}
	//继承：如果 struct 中的一个匿名段实现了一个 method，那么包含这个匿名段的 struct 也能调用该 method。
	//这里并不存在 类似.NET override
	t.ShowA()

	//重写：如果 struct 中的一个匿名段实现了一个 method，包含这个匿名段的 struct 是可以重写匿名字段的方法的。
	t.ShowB()
}

func CompareStruct() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}
	/*
	进行结构体比较时候，只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关。
	invalid operation: sn1 != sn3 (mismatched types struct { age int; name string } and struct { name string; age int })
	  */
	//sn3:= struct {
	//	name string
	//	age  int
	//}{age:11,name:"qq"}
	//if sn1 != sn3 {
	//	fmt.Println("sn1 != sn3")
	//}


	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11,
		m: map[string]string{"a": "1"},
	}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11,
		m: map[string]string{"a": "1"},
	}

	/*
	结构体属性中有不可以比较的类型，如map,slice。 如果该结构属性都是可以比较的，那么就可以使用“==”进行比较操作。
	.\struct.go:74: invalid operation: sm1 == sm2 (struct containing map[string]string cannot be compared)
	 */
	//if sm1 == sm2 {
	//	fmt.Println("sm1 == sm2")
	//}

	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
	}else {
		fmt.Println("sm1 != sm2")
	}
}