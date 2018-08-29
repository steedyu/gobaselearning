package oopsample

import "fmt"

/*
匿名字段
同名字段
通过匿名字段实现继承操作，结构体的名称作为结构体成员
 */
type Person2 struct {
	name string
	age  int
	sex  string
}

type student struct {
	Person2 //匿名字段
	id    int
	name  string
	score int
}

//匿名字段
func anonymousField() {
	var stu student = student{Person2:Person2{name:"zhangsanfeng", age:18, sex:"male"}, id:20, name:"张三丰", score:100}
	////采用就近原则，使用子类信息(Person是父类，student是子类)
	//stu.name = "张三丰"
	////结构体变量.匿名字段.结构体成员
	//stu.Person.name = "zhangsanfeng"
	fmt.Println(stu)
}

//同名字段
func sameField() {
	var stu student
	//采用就近原则，使用子类信息(Person是父类，student是子类)
	stu.name = "张三丰"
	//结构体变量.匿名字段.结构体成员
	stu.Person2.name = "zhangsanfeng"
	fmt.Println(stu)
}


/*
指针匿名字段
 */
type person1 struct {
	name string
	age  int
	sex  string
}

/*
能够能嵌套本结构体指针，结构体不能嵌套本结构体
 */
type student1 struct {
	*person1 //指针匿名字段
	id    int
	score int
}

func pointerAnonymField() {

	var stu student1

	//panic: runtime error: invalid memory address or nil pointer dereference
	//stu.name="郭襄"
	//stu.person1.name = "郭小姐"

	stu.person1 = new(person1)
	stu.name = "郭襄"
	stu.person1.name = "郭小姐"
	fmt.Println(stu.name)
}

/*
多重继承
 1st kind:
TestC ：TestB ：TestA
 */
type TestA struct {
	name string
	id   int
}

type TestB struct {
	TestA
	sex string
	age int
}

type TestC struct {
	TestB
	score int
}

//2nd kind: DemoC : DemoA,DemoB
type DemoA struct {
	name string
	id   int
}

type DemoB struct {
	age int
	sex string
}

type DemoC struct {
	DemoA
	DemoB
	score int
}

func multiInherit() {
	var s TestC
	s.TestB.TestA.name = "xiaoming"
	s.name = "zhangsan"
	s.id = 201
	s.sex = "男"
	s.age = 20
	s.score = 10
	fmt.Println(s)

	var s1 DemoC
	s1.name = "geluowen"
	s1.score = 6
	s1.age = 17
	s1.id = 7
	s1.sex = "male"
	fmt.Println(s1)
}

/*
类型，结构体方法定义
 */
type Int int

func (a Int) add(b Int) Int {
	return a + b
}

type person3 struct {
	name string
	age  int
	sex  string
}

type student3 struct {
	person3
	score int
}

func (p *person3) SayHello() {
	fmt.Println("Hello everyone, i am ", p.name, p.age, p.sex)
}

func typestructMethodDef() {
	var a Int = 10
	value := a.add(20)
	fmt.Println(value)

	var stu student3 = student3{person3{"zhangsan", 23, "male"}, 100}
	//子结构体继承父类结构体 允许使用父类结构体成员 允许使用父类的方法
	stu.SayHello()
}


/*
方法重写
 */
type person5 struct {
	name string
	age  int
	sex  string
}

func (p person5) PrintInfo() {
	fmt.Println("Hello everyone, i am ", p.name, p.age, p.sex)
}

type student5 struct {
	person5
	score int
}

func (s student5) PrintInfo() {
	fmt.Println("Hello everyone, i am ", s.name, s.age, s.sex, s.score)
}

func OverwriteMethod() {
	//方法重写
	var stu student5 = student5{person5{"zhangsan", 23, "male"}, 100}
	//默认使用子类的方法  采用就近原则
	//调用子类方法
	stu.PrintInfo()
	//调用父类方法
	stu.person5.PrintInfo()
}

/*
多态实现
 */
type Personer interface {
	SayHello()
}

type teacher1 struct {
	Person2
}

func (t teacher1) SayHello() {
	fmt.Println("Hello everyone, i am a teacher: ", t.name, t.age, t.sex)
}

type student6 struct {
	Person2
}

func (s student6) SayHello() {
	fmt.Println("Hello everyone, i am a student: ", s.name, s.age, s.sex)
}

//定义一个普通函数，函数的参数为接口类型
//只有一个参数，可以有不同的表现，多态
func SayHello(i Personer) {
	i.SayHello()
}

func polymorphism() {
	//多态实现
	var p Personer
	p = &student6{Person2{"xiaohong", 11, "female"}}
	SayHello(p)
	p = &teacher1{Person2{"panhuimei", 31, "female"}}
	SayHello(p)
}


/*
接口继承和转换
 */

//子集
type Humaner1 interface {
	SayHi()
}

//超集
type Personer1 interface {
	Humaner1        //一组子集的集合
	Sing(string)
}

type student7 struct {
	Person2
}

func (s student7) SayHi() {
	fmt.Println("Hello everyone, i am a student: ", s.name, s.age, s.sex)
}

func (s student7) Sing(name string) {
	fmt.Println("Let me sing a song:", name)
}

func interfaceInheritandConvert() {
	//接口继承和转换
	var h Humaner1
	h = &student7{Person2{"panhuimei", 31, "female"}}
	h.SayHi()

	var p Personer1
	p = &student7{Person2{"panhuimei2", 31, "female"}}
	p.SayHi()
	p.Sing("wuhuanzhige")

	//将超集转成子集  反过来不允许
	h = p
	h.SayHi()
}

func OopBaseDemo() {

	//匿名字段
	anonymousField()

	//同名字段
	sameField()

	//指针匿名字段
	pointerAnonymField()

	//多重继承
	multiInherit()

	//类型，结构体方法定义
	typestructMethodDef()

	//多态实现
	polymorphism()

	//接口继承和转换
	interfaceInheritandConvert()
}

