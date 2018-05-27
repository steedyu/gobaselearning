package ivsample

import (
	"fmt"
	"os"
)

//#1 Range
type Animal struct {
	name string
	legs int
}

func ValuePropertyofRange() {
	zoo := []Animal{
		Animal{"Dog", 4},
		Animal{"Chicken", 2},
		Animal{"Snail", 0},
	}

	fmt.Printf("-> Before update %v\n", zoo)

	/*
	Value property of range (stored here as animal) is a copy of the value from zoo, not a pointer to the value in zoo.
	 */
	for _, animal := range zoo {
		// 🚨 Oppps! `animal` is a copy of an element 😧
		animal.legs = 999
	}
	fmt.Printf("\n-> After update %v\n", zoo)

	for idx := range zoo {
		zoo[idx].legs = 999
	}

	fmt.Printf("\n-> After update %v\n", zoo)
}


//#2 The … thingy
func myFprint(format string, a ...interface{}) {
	if len(a) == 0 {
		fmt.Printf(format)
	} else {

		//fmt.Printf(format, a)

		fmt.Printf(format, a...)

		/*
		In Go, variadic parameters are converted to slices by the compiler
		This means that the variadic argument a is in fact, just a slice. Because of this, the code below is completely valid.
		 */
		// `a` is just a slice!
		for _, elem := range a {
			fmt.Println(elem)
		}
	}
}

func VariadicFeature() {
	myFprint("%s : line %d\n", "file.txt", 49)
}


//#3 Slicing
func CreateNewSlice() {
	data := []int{1, 2, 3}
	slice := data[:2]
	slice[0] = 999

	fmt.Println(data)
	fmt.Println(slice)

	// Option #1
	// appending elements to a nil slice
	// `...` changes slice to arguments for the variadic function `append`
	a := append([]int{}, data[:2]...)
	a[0] = 299
	fmt.Println(data)
	fmt.Println(a)

	// Option #1
	// Create slice with length of 2
	// copy(dest, src)
	a1 := make([]int, 2)
	copy(a1, data[:2])
	a1[0] = 399

	fmt.Println(data)
	fmt.Println(a1)
}

//#4 interface
func retrunerror() error {
	var err error // nil
	return err
}


func returnPathError() error {
	var err *os.PathError // nil
	return err
}

/*
type iface struct {
    tab  *itab
    data unsafe.Pointer
}

itab stand for interface table and is also a structure that holds needed meta information about interface and underlying type:

type itab struct {
    inter  *interfacetype
    _type  *_type
    link   *itab
    bad    int32
    unused int32
    fun    [1]uintptr // variable sized
}

We’re not going to learn the logic of how interface type assertion works,
but what is important is to understand
that interface is a compound of interface and static type information plus pointer to the actual variable (field data in iface)

 */

func InterfaceWithotVariableandVariableEqualsNil() {

	/*
	have an interface with a variable which value equals to nil” and “interface without variable
	 */

	err := retrunerror()	//tab => nil (type) data => nil (no variable)
	fmt.Println("err == nil", err == nil)

	err = returnPathError() //tab => *os.PathError(type)  data => nil (os.PathError)
	fmt.Println("err == nil", err == nil)
}

/*
One of the interface{} related gotchas is the frustration that
you can’t easily assign slice of interfaces to slice of concrete types and vice versa. Something like

func foo() []interface{} {
    return []int{1,2,3}
}

$ go build
cannot use []int literal (type []int) as type []interface {} in return argument

It’s confusing and the beginning. Like, why I can do this conversion with a single variable,
but cannot do with slice? But once you know what is an empty interface (take a look at the picture above again),
it becomes pretty clear, that this “conversion” is actually a quite expensive operation
which involves allocating a bunch of memory and is around O(n) of time and space.
And one of the common approaches in Go design is “if you want to do something expensive - do it explicitly”.

 */
