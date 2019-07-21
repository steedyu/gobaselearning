package gousualsample

import "fmt"

func SliceDemo() {

	//sliceCopy()
	sliceCutandDelete1()
	//sliceCutAndDelete2()
	//sliceInsert()
	//slicePopandShift()
	//slicePush()
	//sliceFilteringWithoutAllocating()
}

func sliceCopy() {

	//copy 实现的两种方式
	//1
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, len(a))
	fmt.Println("before copy:", &b[0], b)
	copy(b, a)
	fmt.Println("after copy:", &b[0], b)

	//2
	/*
	nil 切片的表示
	sliceHeader{
	    Length:        0,
	    Capacity:      0,
	    ZerothElement: nil,
	}
	关键的地方是元素指针也指向了nil。
	 */
	//空分片可以增大（假设它又非零的容量），而nil分片没有存放元素值得数组，而且从不会为了存放元素而增长。
	b1 := []int(nil)
	fmt.Println(b1, b1 == nil, len(b1), cap(b1))

	var b2 []int
	fmt.Println(b2, b2 == nil, len(b2), cap(b2))

	//创建的分片具有零长度（也许还具有零容量），不过它的指针指的不是nil，因此它不是nil分片。
	var b3 []int = make([]int, 0)
	fmt.Println(b3, b3 == nil, len(b3), cap(b3))

	/*
	nil分片功能上等同于零长度的分片，不同的是它不指向任何东西。它具有零长度，
	但可以给其分配空间，添加元素。就像上面例子里那样，你可以看看通过给nil分片添加而实现分片拷贝的那一段。
	 */
	b1 = append(b1, a...)
	fmt.Println("after copy:", &b1[0], b1)

}

func sliceCutandDelete1() {

	/*
	a[n:m] 表示的是区间闭开区间[n:m)
	 */

	//cut
	a := []int{1, 2, 3, 4, 5}
	fmt.Println("before cut:", len(a), &a[0], &a[1], &a[2], &a[3], &a[4])
	fmt.Println(append(a[:2], a[4:]...))
	fmt.Println(a)
	ar := append(a[:2], a[4:]...)
	fmt.Println("after cut ar:", len(ar), &ar[0], &ar[1], &ar[2])
	fmt.Println("after cut a:", len(a), &a[0], &a[1], &a[2], &a[3], &a[4])
	fmt.Println(a)


	//delete
	//1
	a1 := []int{1, 2, 3, 4, 5}
	fmt.Println("before delete 1:", len(a1), &a1[0], &a1[1], &a1[2], &a1[3], &a1[4])
	a1 = append(a1[:2], a1[2 + 1:]...)
	fmt.Println("after delete:", len(a1), &a1[0], &a1[1], &a1[2], &a1[3])
	fmt.Println(a1)

	//2
	a2 := []int{1, 2, 3, 4, 5}
	fmt.Println("before delete 2:", len(a2), &a2[0], &a2[1], &a2[2], &a2[3], &a2[4])
	a2 = a2[:2 + copy(a2[2:], a2[2 + 1:])]
	fmt.Println("after delete:", len(a2), &a2[0], &a2[1], &a2[2], &a2[3])
	fmt.Println(a2)

	//without preserving order
	a3 := []int{1, 2, 3, 4, 5}
	a3[2] = a3[len(a3) - 1]
	a3 = a3[:len(a3) - 1]
	fmt.Println("after delete:", len(a3), a3)


	/*
	 If the type of the element is a pointer or a struct with pointer fields, which need to be garbage collected,
	 the above implementations of Cut and Delete have a potential memory leak problem: some elements with values are
	 still referenced by slice a and thus can not be collected. The following code can fix this problem:
	 */
}

//其逻辑是运用copy 将数组进行重新填充，然后不需要的设置为nil
func sliceCutAndDelete2() {
	//Cut
	oa := []int{1, 2, 3, 4, 5}
	a := []*int{&oa[0], &oa[1], &oa[2], &oa[3], &oa[4]}
	fmt.Println("before delete:", len(a), a)
	copy(a[2:], a[4:])
	for k, n := len(a) - 4 + 2, len(a); k < n; k++ {
		a[k] = nil // or the zero value of T
	}
	a = a[:len(a) - 4 + 2]
	fmt.Println("after delete:", len(a), a)

	// delete
	oa1 := []int{1, 2, 3, 4, 5}
	a1 := []*int{&oa1[0], &oa1[1], &oa1[2], &oa1[3], &oa1[4]}
	copy(a1[2:], a1[2 + 1:])
	a1[len(a1) - 1] = nil // or the zero value of T
	a1 = a1[:len(a1) - 1]

	oa2 := []int{1, 2, 3, 4, 5}
	a2 := []*int{&oa2[0], &oa2[1], &oa2[2], &oa2[3], &oa2[4]}
	a2[2] = a2[len(a2) - 1]
	a2[len(a2) - 1] = nil
	a2 = a2[:len(a2) - 1]
}


//a = append(a[:i], append(make([]T, j), a[i:]...)...)
//a = append(a, make([]T, j)...)


func sliceInsert() {
	a := []int{1, 2, 3, 4, 5}
	a = append(a[:2], append([]int{6}, a[2:]...)...)
	fmt.Println(a)

	/*
	 The second append creates a new slice with its own underlying storage and copies elements in a[i:] to that slice,
	 and these elements are then copied back to slice a (by the first append). The creation of the new slice (and thus memory garbage)
	 and the second copy can be avoided by using an alternative way
	 */
	a1 := []int{1, 2, 3, 4, 5}
	a1 = append(a1, 0)
	copy(a1[3 + 1:], a1[3:])
	a1[3] = 6
	fmt.Println(a1)
}

func slicePopandShift() {

	var x int
	a := []int{1, 2, 3, 4, 5}
	x, a = a[0], a[1:]
	fmt.Println("after pop/shift:", x, len(a), a)

	a = []int{1, 2, 3, 4, 5}
	x, a = a[len(a) - 1], a[:len(a) - 1]
	fmt.Println("after pop back/shift:", x, len(a), a)
}

func slicePush() {
	a := []int{1, 2, 3, 4, 5}
	a = append(a, 6)
	fmt.Println("slicePush 1", a, &a[0])

	//Push Front/Unshift
	a = append([]int{0}, a...)
	fmt.Println("slicePush 2", a, &a[0])
}

func sliceFilteringWithoutAllocating() {
	a := []int{1, 2, 3, 4, 5}
	b := a[:0]
	for _, x := range a {
		if x > 2 {
			b = append(b, x)
		}
	}
	/*
	这里的结果，感觉是在往b指向的位置插入元素，b中的结束和a中的结束一开始就是不相同
	但是指向的数组是同一块区域
	 */
	fmt.Println("after filtering", len(b), len(a), a, b)
}

func sliceReversing() {

	a := []int{1, 2, 3, 4, 5}
	for i := len(a) / 2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	//The same thing, except with two indices:
	for left, right := 0, len(a) - 1; left < right; left, right = left + 1, right - 1 {
		a[left], a[right] = a[right], a[left]
	}

}

