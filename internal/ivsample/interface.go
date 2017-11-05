package ivsample

import (
	"fmt"
	"reflect"
	"unsafe"
)

//////////////////////////////////
type Person interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func UseStructImplementInterface() {
	/*build failed
	  cannot use Stduent literal (type Stduent) as type Person in assignment:
	  Stduent does not implement Person (Speak method has pointer receiver)
	*/
	//var peo Person = Student{}
	var peo Person = new(Student)
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

//////////////////////////////////

//////////////////////////////////
func live() Person {
	var stu *Student = nil
	return stu
}

func UseStructImplementInterface3() {
	p := live()

	if p == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}

	//这个考点是很多人忽略的interface内部结构.go中的接口分为两种一种是空的接口类似这样,
	fmt.Println("live() case result:", p)
	var in interface{}
	if in == nil {
		fmt.Println("var in interface{} case result:", in)
	}
}

//////////////////////////////////

//////////////////////////////////
type YourType interface {
	SayHi(str string)
}
type HeType interface {
	SayHi(str string)
	SayBye(str string)
}

type youknow struct {
}
func (t youknow) SayHi(str string) {
	fmt.Println("hi " + str)
}

type dknow struct {
	youknow
}
func (t *dknow) SayBye(str string) {
	fmt.Println("bye bye " + str + ".see you tomorrow.")
}

/*
 dknow组合的youknow实现了SayHi方法，指向了Struct类型，自己实现的SayBye，又指向了Ptr，这种情况怎么算？
 答案是，SayHi指向的非指针对象，golang会自动生成一个指向指针的SayHi方法，那么dknow不就有指向指针的SayHi域SayBye两个方法了嘛，所以dknow就实现了HeType接口了。

 golang 会根据struct类型生成一个Ptr的方法，反之不会
 */
func UseStructImplementInterface2() {
	var myouknow YourType = youknow{}
	myouknow.SayHi("demo")
	var mdknow HeType = new(dknow)
	mdknow.SayBye("demo")
}

//////////////////////////////////


/*
接口类型的变量底层是作为两个成员来实现，一个是type，一个是data。type用于存储变量的动态类型，data用于存储变量的具体数据。

接下来说说interface类型的值和nil的比较问题。这是个比较经典的问题，也算是golang的一个坑。
 */

func InterfaceNilIssue() {
	//x转成interface之后 type类型是*int而不是nil
	var x *int = nil
	foo(x)

	//x1转成interface之后 type类型是int data部分是0
	var x1 int
	foo(x1)


	/*
	变量x2是interface类型，它的底层结构必然是(type, data)。
	由于nil是untyped(无类型)，而又将nil赋值给了变量val，所以val实际上存储的是(nil, nil)。
	因此很容易就知道val和nil的相等比较是为true的。
	 */
	var x2 interface{}
	foo(x2)

	var x3 interface{} = (*interface{})(nil)
	foo(x3)
}

func foo(x interface{}) {
	typeret := reflect.TypeOf(x)

	if x == nil {
		fmt.Println(typeret, "empty interface")
		return
	}
	fmt.Println(typeret, "non-empty interface")

}

func InterfaceNilIssue1() {
	var val interface{} = int64(58)
	fmt.Println(reflect.TypeOf(val))
	val = 50
	fmt.Println(reflect.TypeOf(val))
}

//////////////////////////////////

type data struct{}

func (this *data) Error() string { return "" }

func test() error {
	var p *data = nil
	return p
}

func InterfaceNilIssue2() {
	var e error = test()

	/*
	error是一个接口类型，test方法中返回的指针p虽然数据是nil，
	但是由于它被返回成包装的error类型，也即它是有类型的。所以它的底层结构应该是(*data, nil)，很明显它是非nil的
	 */
	d := (*struct {
		itab uintptr
		data uintptr
	})(unsafe.Pointer(&e))
	fmt.Println(d)

	if e == nil {
		fmt.Println("e is nil")
	} else {
		fmt.Println("e is not nil")
	}
}

func bad() bool {
	return false
}

func test2() error {
	var p *data = nil
	if bad() {
		return p
	}
	return nil
}

func InterfaceNilIssue3() {
	var e error = test2()

	/*
	error是一个接口类型，test方法中返回的指针p虽然数据是nil，
	但是由于它被返回成包装的error类型，也即它是有类型的。所以它的底层结构应该是(*data, nil)，很明显它是非nil的
	 */
	d := (*struct {
		itab uintptr
		data uintptr
	})(unsafe.Pointer(&e))
	fmt.Println(d)

	if e == nil {
		fmt.Println("e is nil")
	} else {
		fmt.Println("e is not nil")
	}
}





