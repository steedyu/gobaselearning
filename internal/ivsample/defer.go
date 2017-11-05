package ivsample

import "fmt"

func DeferCallOrder() {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()
	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	panic("触发异常")
}

func DeferParamterMethod() {
	a := 1
	b := 2
	// 作为defer的calc方法参数的calc方法会被先执行，从而得出参数
	// 而作为defer的calc方法会在defer 会被排在defer最后一个执行
	defer calc("1", a, calc("10", a, b))
	a = 0
	// 同上
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func DeferParamterMethod2() {
	a := 1
	b := 2
	defer func(c int) {
		calc("1", a, c)
	}(calc("10", a, b))
	a = 0
	defer func() {
		calc("2", a, calc("20", a, b))
	}()
	b = 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}


func deferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func deferFunc1_1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return
}

var gt int
func deferFunc2(i int) int {
	t := i
	gt = i
	defer func() {
		t += 3
		gt += 3
		fmt.Println("after body of defer in deferFunc2:", "t=", t, "gt=", gt)
	}()
	return gt
}

func deferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func deferFunc3_1(i int) (t int) {
	if i > 2 {
		return 0
	}

	defer func() {
		t += i
	}()

	return 2
}

func deferFunc3_2(i int) (t int) {
	defer func() {
	 	t += t * 5 + 1
	}()

	if i > 2 {
		return 0
	}

	defer func() {
		t += i
	}()

	return 2
}

/*
defer和函数返回值
需要明确一点是defer需要在函数结束前执行。
函数返回值名字会在函数起始处被初始化为对应类型的零值并且作用域为整个函数
deferFunc1有函数返回值t作用域为整个函数，在return之前defer会被执行，所以t会被修改，返回4;
deferFunc2函数中t的作用域为函数，返回1;
(其实这里使用了一个全局变量gt的情况下，返回仍为1，这里的运行结果的理解上更像是，当你没有定义返回值名字时，return 语句执行的时候，返回值就确定了；
而如果使用返回值名字的时候，在整个方法真正结束的时候，这个值才会被最终确定。同时结果也与是否在return语句处有这个和无这个变量没有关系
在defFunc2中,先执行到return语句，再执行defer语句
)
deferFunc3返回3
从retrun结果3理解，在return 2这句执行的时候，会将2赋值到t中，而在defer执行的时候，t += i => 2 + 1 =3 最后返回等于3
deferFunc3_1
defer 的位置写在方法中 retrun 后面的位置，如果在运行到defer之前已经return，defer将不会被执行；如果在运行到defer之前没有return,defer将会被执行
 */
func DeferInMethodWithReturnValue() {
	fmt.Println(deferFunc1(1))
	fmt.Println(deferFunc1_1(1))
	fmt.Println(deferFunc2(1))
	fmt.Println("after deferFunc2 gt=",gt)
	fmt.Println(deferFunc3(1))
	fmt.Println(deferFunc3_1(3))
	fmt.Println(deferFunc3_1(1))
	fmt.Println(deferFunc3_2(3))
	fmt.Println(deferFunc3_2(1))
}