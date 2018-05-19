package ivsample

import (
	"fmt"
)

/*

If you are coming from a language with no notion of pointers, or where every variable is implicitly a pointer don’t panic,
forming a mental model of how variables and pointers relate takes time and practice. Just remember this rule:
A pointer is a value that points to the memory address of another variable.
 */


/*
Strings are value types, not pointers, which is the, IMO(abbr in my opinion),
  the number one cause of null pointer exceptions in languages like Java and C++.
 */
func StringIsValueType() {

	var str string
	if str == "" {
		fmt.Println("StringIsValueType func:", str)
	}
}


/*
Go does not have reference variables
each variable defined in a Go program occupies a unique memory location.

It is not possible to create a Go program where two variables share the same storage location in memory.
It is possible to create two variables whose contents point to the same storage location,
but that is not the same thing as two variables who share the same storage location.
 */
func NoReferenceVariables() {
	var a, b, c int
	fmt.Println("&a:", &a, "&b:", &b, "&c:", &c) // 0x1040a124 0x1040a128 0x1040a12c

	var a1 int
	var b1, c1 *int = &a1, &a1
	fmt.Println("b1:", b1, "c1:", c1)
	fmt.Println("&b1:", &b1, "&c1:", &c1)

	var a2 int = 0
	var b2, c2 int = a2, a2
	fmt.Println("b2:", b2, "c2:", c2)
	fmt.Println("&b2:", &b2, "&c2:", &c2)
}

/*
Go does not have pass-by-reference semantics because Go does not have reference variables.
 */
func AreMapsandChannelsReferencesPass() {
	/*
	这里的理解，联想到.NET 中List<T>变量的传递 List<T>是引用型  但是在方法内进行new或者null的指向，都不会生效
	引用本身按值传递，因此指向另一个对象或者空不会起作用，但引用对象本身状态则可以被修改

	golang 的方法参数都是值传递，当传递 map channel等引用类型时，仍然是值传递
	 */
	// int 及 *int 举例
	var a int
	fmt.Println("&a:", &a, "a:", a)
	fnInt(a)
	fnIntpointer(&a)
	fmt.Println(&a, a)
	fnIntpointer1(&a)

	//map 举例
	var m map[int]int
	fnMakeMap(m)
	fmt.Println("m == nil", m == nil)
	m = make(map[int]int)
	fmapadditem(m)
	fmt.Println(m)

}

func fnIntpointer(a *int) {
	//&a 参数 a *int 分配的指针地址  a 本身的值 是一个*int的指针值
	fmt.Println("before fnIntpointer parameter a &a:", &a, "a:", a)
	//对于 a *int %p表示的是 a 所承载的指针值
	fmt.Println("a %p result:", fmt.Sprintf("%p", a))
	var b int = 2
	fmt.Println("&b:", b)
	a = &b
	fmt.Println("after fnIntpointer parameter a &a:", &a, "a:", a, "*a:", *a)
}

func fnIntpointer1(a *int) {
	*a = 2
	fmt.Println("after fnIntpointer1 parameter a &a:", &a, "a:", a, "*a:", *a)
}

func fnInt(a int) {
	a = 2
	fmt.Println("fnInt", &a, 2)
}

func fnMakeMap(m map[int]int) {
	m = make(map[int]int)
}

func fmapadditem(m map[int]int) {
	m[1] = 1
}


