package ivsample

import "fmt"

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
