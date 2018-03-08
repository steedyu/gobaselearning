package gousualsample

import (
	"fmt"
)

/*
Make zero values useful
 */

type A struct {
	value string
}

func (a *A) Test() string {
	fmt.Println("Test Called")	//这句执行
	return a.value        //执行到此处时，才会抛出panic
}

func (a *A) Test2() string {
	fmt.Println("Test2 Called")	//这句执行
	if a == nil {
		return ""
	}

	return a.value
}

func getA() *A {
	return nil
}

func StructZeroUsefulSample() {
	a := getA()
	a.Test2()
	a.Test()        //这行不会报错，仍然会执行进去
}
