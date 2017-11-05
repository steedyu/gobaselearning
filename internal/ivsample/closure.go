package ivsample

func genfuncArr() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		//这里funs没有初始化，也没有报错
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	return funs
}

func genfuncArr2() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		x := i
		funs = append(funs, func() {
			println(&x, x)
		})
	}
	return funs
}

func ClosureCase1() {
	funs := genfuncArr()
	for _, f := range funs {
		f()
	}

	funs2 := genfuncArr2()
	for _, f := range funs2 {
		f()
	}
}

func genfuncArr3(x int) (func(), func()) {
	return func() {
		println(x)
		x += 10
	}, func() {
		println(x)
	}
}

func ClosureCase2() {
	a, b := genfuncArr3(100)
	a()
	b()
}