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
	fmt.Println("---------------------exsample 1---------------------------")

	x := []int{1, 2, 3}
	fmt.Printf("original x addr :%p\r\n", x)
	/*
	分析下下面的过程，
	对x进行了复制，这里可以理解复制了一个 3个元素底层数组，长度是3的切片（窗口），
	所以它最后的循环是3次，虽然过程中不断地对切片x进行了append，但是range表达式只会在for语句开始执行时被求值一次，并且被迭代的对象是range表达式结果值的副本而不是原值
	虽然这里x是切片，是引用类型，但是append 并没有改变for range迭代副本切片中的元素
	 */
	for i := range x {
		x = append(x, i)
		fmt.Printf("in forrange append x addr: %p\r\n", x)
	}
	fmt.Println("after append x:", x)
	fmt.Println("---------------------exsample 1---------------------------")

	/*
	分析下下面的过程，
	对x进行了复制，这里可以理解复制了一个 6个元素底层数组，长度是6的切片（窗口），
	所以它最后的循环是6次
	第一种情况
	在第0次的时候,对x切片进行append 去除中间元素，由于没有超过cap的大小，所以修改的底层数组并没有重新分配空间=>和forrange复制的那个切片指向同一个
	=>所以在forrange 遍历的结果里，值是发生改变的

	第二种情况 在第0次的时候,对x切片进行append， 由于超过cap的大小，对底层数组重新分配，然后对新的底层数组进行修改=>并没有影响forrange复制那个切片指向的底层数组
	=>所以在forrange 遍历的结果里，值是没有发生改变的
	 */
	fmt.Println("---------------------exsample 2---------------------------")
	for i, e := range x {
		fmt.Println("range i , e :", i, e)
		if i == 0 {
			//第一种情况 没有引起cap变动
			//fmt.Printf("for range i:%v times, x addr :%p, x's length: %d, x's cap: %d\r\n", i, x, len(x), cap(x))
			//x = append(x[:3], x[4:]...)
			//fmt.Printf("for range i:%v times, x addr after append modify :%p\r\n", i, x)

			//第二种情况 引起cap变动， 并且同时修改了第二个索引的元素值
			fmt.Printf("for range i:%v times, x addr :%p, x's length: %d, x's cap: %d\r\n", i, x, len(x), cap(x))
			x = append(x, 7)
			x[i + 1] = 10 + x[i + 1]
			fmt.Printf("for range i:%v times, x addr after append modify :%p\r\n", i, x)

		}

	}
	fmt.Printf("x : %v, x addr: %p, x's length: %d, x's cap: %d\r\n", x, x, len(x), cap(x))
	fmt.Println("---------------------exsample 2---------------------------")
}

/*
the slice value is reduced to a length of 2 inside the loop, but the loop is operating on its own copy of the slice value.
This allows the loop to iterate using the original length without any problems since the backing array is still in tact.
If the code uses the pointer semantic form of the for range, the program panics.
 */
func ForrangeInOriginalLength() {
	five := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	for _, v := range five {
		five = five[:2]
		fmt.Printf("v[%s]\n", v)
	}
}

/*
The for range took the length of the slice before iterating, but during the loop that length changed.
Now on the third iteration, the loop attempts to access an element that is no longer associated with the slice’s length.
 */
func ForrangeInOriginalLength2() {
	five := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	for i := range five {
		five = five[:2]
		fmt.Printf("v[%s]\n", five[i])
	}
}

type user struct {
	name  string
	likes int
}

func (u *user) notify() {
	fmt.Printf("%s has %d likes\n", u.name, u.likes)
}

func (u *user) addLike() {
	u.likes++
}

func ForrangeStructArray() {
	users := []user{
		{name: "bill"},
		{name: "lisa"},
	}

	for _, u := range users {
		u.addLike()
	}

	for _, u := range users {
		u.notify()
	}

	for i := range users {
		users[i].addLike()
	}

	for i := range users {
		users[i].notify()
	}
}

/*
range表达式只会在for语句开始执行时被求值一次，无论后边会有多少次迭代；
range表达式的求值结果会被复制，也就是说，被迭代的对象是range表达式结果值的副本而不是原值。
 */
func ForrangeDemo() {
	//这个是数组类型  值类型
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		fmt.Println(i, e)
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i + 1] += e
		}
	}
	fmt.Println(numbers2)

}

func ForrangeDemo1() {
	//这个是切片类型 引用
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		fmt.Println(i, e)
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i + 1] += e
		}
	}
	fmt.Println(numbers2)

}