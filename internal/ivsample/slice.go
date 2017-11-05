package ivsample

import "fmt"

func Sliceappend() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s) //[0 0 0 0 0 1 2 3]
}

func SliceNew() {
	/*
	.\slice.go:13: first argument to append must be slice; have *[]int
	编译不通过
	new 初始化生成了一个指向[]int这个类型的指针，所以append不能操作它
	 */
	list := new([]int)
	//list = append(list, 1)
	*list = append(*list, 1)
	fmt.Println(list)
}

func Sliceappend2() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	//.\slice.go:26: cannot use s2 (type []int) as type int in append
	//s1 = append(s1, s2)
	s1 = append(s1, s2...)
	fmt.Println(s1)
}
