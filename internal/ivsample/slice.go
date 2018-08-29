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

func SliceDiffSliceForSameArr() {

	sl1 := []int{1, 2, 3, 4, 5, 6}
	sl2 := sl1[0:2]

	fmt.Println(len(sl1), cap(sl1), fmt.Sprintf("%p", sl1))
	fmt.Println(len(sl2), cap(sl2), fmt.Sprintf("%p", sl2))
}

func SliceFuncParameter() {

	appfunc := func(slice []int) {
		fmt.Println(len(slice), cap(slice), slice, fmt.Sprintf("%p", slice))
		slice = append(slice, 0, 1, 2)
		fmt.Println(len(slice), cap(slice), slice, fmt.Sprintf("%p", slice))
	}

	sl1 := []int{1, 2, 3, 4, 5, 6}
	sl2 := sl1[0:2]
	appfunc(sl2)
	fmt.Println(sl2)
	fmt.Println(sl1)

}

func SliceArrandTheirPoint() {

	var arr [5]int = [5]int{0, 1, 2, 3, 4}
	var p *[5]int = &arr
	fmt.Printf("%p, %p \n", p, &arr)

	var slice []int = []int{0, 1, 2, 3, 4}
	var sp *[]int = &slice
	//切片名本身是一个地址
	fmt.Printf("%p, %p \n", sp, slice)
	fmt.Printf("%p, %p \n", *sp, slice)
}


