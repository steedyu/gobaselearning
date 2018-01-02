package ivsample

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func ForrangePointer() {

	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	/*
	在for range的时候 stu的指针是固定某个地址
	经过for range之后，stu的指针指向的student的实例是最后一个循环的值
	 */
	m := make(map[string]*student)
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	fmt.Println("错误用法结果---------------------")
	for name, student := range m {
		fmt.Println(name, *student)
	}

	/*
	期望结果
	 */
	expm := make(map[string]*student)
	expm1 := make(map[string]student)
	for _, stu := range stus {
		expm[stu.Name] = &student{}
		*expm[stu.Name] = stu

		expm1[stu.Name] = stu
	}
	//指针
	fmt.Println("期望结果1---------------------")
	for name, student := range expm {
		fmt.Println(name, *student)
	}
	//非指针
	fmt.Println("期望结果2---------------------")
	for name, student := range expm1 {
		fmt.Println(name, student)
	}
}

func ForrangeAppend() {
	/*
	在Go的for…range循环中，range 一个slice的时候，其实range 也是一个指向	slice内容的引用
	第一个循环体  append对x的容量(cap)进行了改变，所以构造了一个新的切片，然后x指向了一个新的切片
	第二个循环体 append对x的容量(cap)为进行改变，所以并没有改变引用关系
	****GoLang里对slice一定要谨慎使用append操作。
	cap未变化时，slice是对数组的引用，并且append会修改被引用数组的值。
	append操作导致cap变化后，会复制被引用的数组，然后切断引用关系。
	 */
	x := []int{1, 2, 3}
	for i := range x {
		x = append(x, i)
		fmt.Println(x)
	}
	fmt.Println(x)


	for i, e := range x {
		fmt.Println("range i , e", i, e)
		if i == 0 {
			fmt.Println(&x[0])
			//x = []int{1, 2, 3, 4}
			x = append(x[:3], x[4:]...)
			fmt.Println(&x[0])
		}
		fmt.Println(x)
	}
}