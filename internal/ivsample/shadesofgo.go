package ivsample

import (
	"fmt"
	"unicode/utf8"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"sync"
	"time"
)

//http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html


//1 Opening Brace Can't Be Placed on a Separate Line
//func() main {
//
//}

// Unused Variables ,Unused Imports

//Short Variable Declarations Can Be Used Only Inside Functions
/*

myvar := 1 //error

Compile Error:
func main() {
}

Works:
var myvar = 1

func main() {
}
 */


//Redeclaring Variables Using Short Variable Declarations
func RedeclaingVaiablesUsingShortVariable() {

	one := 0
	//one := 1 //error

	one, two := 1, 2
	one, two = two, one
}


//Accidental Variable Shadowing
/*
You can use the vet command to find some of these problems. By default,  vet will not perform any shadowed variable checks.
Make sure to use the  -shadow flag: go tool vet -shadow your_file.go

Note that the vet command will not report all shadowed variables. Use  go-nyet for more aggressive shadowed variable detection.
 */
func AccidentalVariableShadowing() {
	x := 1
	fmt.Println(x)     //prints 1
	{
		fmt.Println(x) //prints 1
		x := 2
		fmt.Println(x) //prints 2
	}
	fmt.Println(x)     //prints 1 (bad if you need 2)
}

//Can't Use "nil" to Initialize a Variable Without an Explicit Type
//var x = nil //error

//Map Capacity
/*
You can specify the map capacity when it's created, but you can't use the cap() function on maps.
 */
func MapCapacity() {
	m := make(map[string]int, 99)
	//cap(m) //error
	fmt.Println(m)
}

//Strings Can't Be "nil"

//Array Function Arguments
func ArrayFunctionArguments() {

	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr) //prints [7 2 3]
	}(x)

	fmt.Println(x) //prints [1 2 3] (not ok if you need [7 2 3])

	func(arr *[3]int) {
		(*arr)[0] = 7
		fmt.Println(arr) //prints &[7 2 3]
	}(&x)

	fmt.Println(x) //prints [7 2 3]


	/*
	Another option is to use slices.
	Even though your function gets a copy of the slice variable it still references the original data.
	 */
	x1 := []int{1, 2, 3}

	func(arr []int) {
		arr[0] = 7
		fmt.Println(arr) //prints [7 2 3]
	}(x1)

	fmt.Println(x1) //prints [7 2 3]
}

//Slices and Arrays Are One-Dimensional
/*
Creating a dynamic multi-dimensional array using slices of "independent" slices is a two step process.
 First, you have to create the outer slice.
 Then, you have to allocate each inner slice. The inner slices are independent of each other. You can grow and shrink them without affecting other inner slices.
 */
func MultiDimensionalArray() {
	h, w := 2, 4

	raw := make([]int, h * w)
	for i := range raw {
		raw[i] = i
	}
	fmt.Println(raw, &raw[4])
	//prints: [0 1 2 3 4 5 6 7] <ptr_addr_x>

	table := make([][]int, h)
	for i := range table {
		table[i] = raw[i * w:i * w + w]
	}

	fmt.Println(table, &table[1][0])
	//prints: [[0 1 2 3] [4 5 6 7]] <ptr_addr_x>
}

//Accessing Non-Existing Map Keys


//Strings Are Immutable
func StringsUpdate() {
	x := "text"
	//x[0] = 'T'  error
	xbytes := []byte(x)
	xbytes[0] = 'T'

	/*
	Note that this isn't really the right way to update characters in a text string because a given character could be stored in multiple bytes.
	If you do need to make updates to a text string convert it to a rune sclice first.
	Even with rune slices a single character might span multiple runes, which can happen if you have characters with grave accent,
	for example. This complicated and ambiguous nature of "characters" is the reason why Go strings are represented as byte sequences.
	 */

	fmt.Println(string(xbytes)) //prints Text
}


//Strings Are Not Always UTF8 Text
func StringAreNoteAlwaysUTF8() {
	data1 := "ABC"
	fmt.Println(utf8.ValidString(data1)) //prints: true

	data2 := "A\xfeC"
	fmt.Println(utf8.ValidString(data2)) //prints: false
}

/*
The built-in len() function returns the number of bytes instead of
the number of characters like it's done for unicode strings in Python.
 */
func StringLength() {
	data := "♥"
	fmt.Println(utf8.RuneCountInString(data)) //prints: 1

	/*
	Technically the RuneCountInString() function doesn't return the number of characters
	because a single character may span multiple runes.
	 */
	data1 := "é"
	fmt.Println(len(data1))                    //prints: 3
	fmt.Println(utf8.RuneCountInString(data1)) //prints: 2
	fmt.Printf("%+q", data1)
}

/*
Missing Comma In Multi-Line Slice, Array, and Map Literals
 x := []int{
    1,
    2 //error
    }

 */

//Iteration Values For Strings in "range" Clauses
func IterationValuesForStringsInRangeClauses() {

	/*
	The index value (the first value returned by the "range" operation) is the index of the first byte
	for the current "character" (unicode code point/rune) returned in the second value.
	It's not the index for the current "character" like it's done in other languages.
	Note that an actual character might be represented by multiple runes.
	ake sure to check out the "norm" package (golang.org/x/text/unicode/norm) if you need to work with characters.
	 */

	data := "A\xfe\x02\xff\x04"
	for _, v := range data {
		fmt.Printf("%#x ", v)
	}
	//prints: 0x41 0xfffd 0x2 0xfffd 0x4 (not ok)

	/*
	The for range clauses with string variables will try to interpret the data as UTF8 text.
	For any byte sequences it doesn't understand it will return 0xfffd runes (aka unicode replacement characters)
	instead of the actual data. If you have arbitrary (non-UTF8 text) data stored in your string variables,
	 make sure to convert them to byte slices to get all stored data as is.
	 */

	fmt.Println()
	for _, v := range []byte(data) {
		fmt.Printf("%#x ", v)
	}
	//prints: 0x41 0xfe 0x2 0xff 0x4 (good)

	/*
	下面这串sample的字符串，按照每个字节来读取数据，可以还正确输出数据
	如果按照rune string来获取 会有几个不认，但是最后一个特殊字符 确实由代码点3个字节组成的
	 */

	fmt.Println()
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	for _, v := range []byte(sample) {
		fmt.Printf("%#x,%q ", v, v)
	}
	fmt.Println(utf8.ValidString(sample))
	for _, v := range sample {
		fmt.Printf("%#x,%q,%v ", v, v, utf8.ValidRune(v))

	}
	fmt.Println()

	r, s := utf8.DecodeRuneInString(sample)
	fmt.Println(r, s)

	t := "🐶"
	fmt.Printf("\n %#x, %+q", t, t)
	fmt.Println(len(t))

}


//Fallthrough Behavior in "switch" Statements
func FallthroughInSwitch() {

	isSpace1 := func(ch byte) bool {
		switch(ch) {
		case ' ': //error
		case '\t':
			return true
		}
		return false
	}

	fmt.Println(isSpace1('\t')) //prints true (ok)
	fmt.Println(isSpace1(' '))  //prints false (not ok)


	isSpace2 := func(ch byte) bool {
		switch(ch) {
		case ' ', '\t':
			return true
		}
		return false
	}

	fmt.Println(isSpace2('\t')) //prints true (ok)
	fmt.Println(isSpace2(' '))  //prints true (ok)

	/*
	Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch,
	但是可以使用fallthrough强制执行后面的case代码，fallthrough不会判断下一条case的expr结果是否为true。
	 */
	isSpace3 := func(ch byte) bool {
		switch(ch) {
		case ' ': //error
			fallthrough
		case '\t':
			return true
		}
		return false
	}

	fmt.Println(isSpace3('\t')) //prints true (ok)
	fmt.Println(isSpace3(' '))  //prints true (ok)
}


//Increments and Decrements
/*
Unlike other languages, Go doesn't support the prefix version of the operations.
 You also can't use these two operators in expressions.
 */
func IncrandDecr() {
	data := []int{1, 2, 3}
	i := 0
	//++i //error  syntax error: unexpected ++, expecting }
	i++
	//fmt.Println(data[i++]) //error syntax error: unexpected ++, expecting ]
	fmt.Println(data[i])
}


//Bitwise NOT Operator
//TODO:

//Operator Precedence Differences
//TODO:


//App Exits With Active Goroutines
//Sending to an Unbuffered Channel Returns As Soon As the Target Receiver Is Ready
//Sending to an Closed Channel Causes a Panic

//Methods with Value Receivers Can't Change the Original Value
type data1 struct {
	num   int
	key   *string
	items map[string]bool
}

func (this *data1) pmethod() {
	this.num = 7
}

func (this data1) vmethod() {
	this.num = 8
	*this.key = "v.key"
	this.items["vmethod"] = true
}

/*
Method receivers are like regular function arguments.
If it's declared to be a value then your function/method gets a copy of your receiver argument.
This means making changes to the receiver will not affect the original value unless
your receiver is a map or slice variable and
you are updating the items in the collection or
the fields you are updating in the receiver are pointers.
 */
func MethodsValueReceiverCannotChangetheOriginalVal() {
	key := "key.1"
	d := data1{1, &key, make(map[string]bool)}

	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)
	//prints num=1 key=key.1 items=map[]

	d.pmethod()
	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)
	//prints num=7 key=key.1 items=map[]

	d.vmethod()
	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)
	//prints num=7 key=v.key items=map[vmethod:true]
}



//Closing HTTP Connections
/*
Some HTTP servers keep network connections open for a while (based on the HTTP 1.1 spec and the server "keep-alive" configurations).
 By default, the standard http library will close the network connections only when the target HTTP server asks for it.
 This means your app may run out of sockets/file descriptors under certain conditions.
 */
func HTTPConnections() {
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Close = true
	//or do this:
	//req.Header.Add("Connection", "close")

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(string(body)))
}

func HTTPConnections2() {
	/*
	You can also disable http connection reuse globally. You'll need to create a custom http transport configuration for it.
	 */
	tr := &http.Transport{DisableKeepAlives: true}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("http://golang.org")
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(string(body)))
}
/*
If you send a lot of requests to the same HTTP server it's ok to keep the network connection open.
However, if your app sends one or two requests to many different HTTP servers in a short period of time
it's a good idea to close the network connections right after your app receives the responses.
Increasing the open file limit might be a good idea too. The correct solution depends on your application though.
 */


//Unmarshalling JSON Numbers into Interface Values
func UnmarshallingJsonNumbersIntoInterfaceValues() {
	var data = []byte(`{"status": 200}`)

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error:", err)
		return
	}

	/*
	By default, Go treats numeric values in JSON as float64 numbers when you decode/unmarshal JSON data into an interface.
	 */

	//var status = result["status"].(int) //error  panic: interface conversion: interface is float64, not int
	//var status,_ = result["status"].(json.Number).Int64() //ok  interface conversion: interface is float64, not json.Number
	var status = uint64(result["status"].(float64)) //ok
	fmt.Println("status value:", status)
}

//use a Decoder type to unmarshal JSON and tell it to represent JSON numbers using the Number interface type.
func UnmarshallingJsonNumbersIntoInterfaceValues2() {
	var data = []byte(`{"status": 200}`)

	var result map[string]interface{}
	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	if err := decoder.Decode(&result); err != nil {
		fmt.Println("error:", err)
		return
	}

	var status, _ = result["status"].(json.Number).Int64() //ok
	fmt.Println("status value:", status)
}

//use a struct type that maps your numeric value to the numeric type you need
func UnmarshallingJsonNumbersIntoInterfaceValues3() {

	var data = []byte(`{"status": 200}`)

	var result struct {
		Status uint64 `json:"status"`
	}

	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&result); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("result => %+v", result)
	//prints: result => {Status:200}
}


//This option is useful if you have to perform conditional JSON field decoding where the field type or structure might change.
func UnmarshallingJsonNumbersIntoInterfaceValuesV() {
	records := [][]byte{
		[]byte(`{"status": 200, "tag":"one"}`),
		[]byte(`{"status":"ok", "tag":"two"}`),
	}

	/*
	type Command struct {
	    ID   int
	    Cmd  string
	    Args *json.RawMessage
	}
	使用json.RawMessage的话，Args字段在Unmarshal时不会被解析，直接将字节数据赋值给Args。
	我们可以能先解包第一层的JSON数据，然后根据Cmd的值，再确定Args的具体类型进行第二次Unmarshal。
	这里要注意的是，一定要使用指针类型*json.RawMessage，否则在Args会被认为是[]byte类型，在打包时会被打包成base64编码的字符串。
	 */

	for idx, record := range records {
		var result struct {
			StatusCode uint64
			StatusName string
			Status     json.RawMessage `json:"status"`
			Tag        string             `json:"tag"`
		}

		if err := json.NewDecoder(bytes.NewReader(record)).Decode(&result); err != nil {
			fmt.Println("error:", err)
			return
		}

		var sstatus string
		if err := json.Unmarshal(result.Status, &sstatus); err == nil {
			result.StatusName = sstatus
		}

		var nstatus uint64
		if err := json.Unmarshal(result.Status, &nstatus); err == nil {
			result.StatusCode = nstatus
		}

		fmt.Printf("[%v] result => %+v\n", idx, result)
	}
}

func JsonRawMessageDemo() {
	raw := json.RawMessage(`{"foo":"bar"}`)

	//j, err := json.Marshal(raw)

	//The value passed to json.Marshal must be a pointer for json.RawMessage to work properly:
	j, err := json.Marshal(&raw)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
}


//Comparing Structs, Arrays, Slices, and Maps
/*
You can use the equality operator, ==, to compare struct variables
 if each structure field can be compared with the equality operator.
 */
//TODO:
type comparedata struct {
	num     int
	fp      float32
	complex complex64
	str     string
	char    rune
	yes     bool
	events  <-chan string
	handler interface{}
	ref     *byte
	raw     [10]byte
}

func CompareStruct1() {
	v1 := comparedata{}
	v2 := comparedata{}
	fmt.Println("v1 == v2:", v1 == v2) //prints: v1 == v2: true
}


//Recovering From a Panic
func RecoverIncorrectUsed() {
	recover() //doesn't do anything
	panic("not good")
	recover() //won't be executed :)
	fmt.Println("ok")
}
/*
The recover() function can be used to catch/intercept a panic.
Calling  recover() will do the trick only when it's done in a deferred function.
 */

func doRecover() {
	fmt.Println("recovered =>", recover()) //prints: recovered => <nil>
}

func RecoverIncorrectUsed2() {
	defer func() {
		doRecover() //panic is not recovered
	}()

	panic("not good")
}

//The call to recover() works only if it's called directly in your deferred function.
func RecoverCorrectUsed() {
	defer func() {
		fmt.Println("recovered:", recover())
	}()

	panic("not good")
}

//"Hidden" Data in Slices
/*
When you reslice a slice, the new slice will reference the array of the original slice.
If you forget about this behavior it can lead to unexpected memory usage
if your application allocates large temporary slices creating new slices from them to refer to small sections of the original data.
 */
func getHiddenSliceData1() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0]) //prints: 10000 10000 <byte_addr_x>
	return raw[:3]
}

/*
To avoid this trap make sure to copy the data you need from the temporary slice (instead of reslicing it).
 */
func getHiddenSliceData2() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0]) //prints: 10000 10000 <byte_addr_x>
	res := make([]byte, 3)
	copy(res, raw[:3])
	return res
}

func HiddenSliceDataDemo() {
	data := getHiddenSliceData1()
	fmt.Println(len(data), cap(data), fmt.Sprintf("%p", data)) //prints: 3 10000 <byte_addr_x>


	data1 := getHiddenSliceData2()
	fmt.Println(len(data1), cap(data1), fmt.Sprintf("%p", data1)) //prints: 3 3 <byte_addr_y>
}

//Slice Data "Corruption"

func SliceDataCorruption() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	dir1 := path[:sepIndex]
	dir2 := path[sepIndex + 1:]
	fmt.Println("dir1 =>", string(dir1), len(dir1), cap(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2), len(dir2), cap(dir2)) //prints: dir2 => BBBBBBBBB

	fmt.Println("before path append=>", string(path), len(path), cap(path), &path[0])

	//情况1
	//dir1 = append(dir1, "suffix"...)

	//情况2
	/*
	这种情况，增加元素后，并没有超过原来数组的cap，不重新分配；而append之后，并没有重新指向新区域，dir1仍然指向原来区域，
	但是修改的数组生肖，path和dir2的指向内容即原来的数组发生了改变
	 */
	//_ = append(dir1, "suffix"...)

	//情况3
	/*
	这种情况，由于增加的元素个数超过了dir1所指数组的cap大小，所以会重新分配数组给dir1，而原来path和dir2指向的的数组并不会进行改变
	 */
	dir1 = append(dir1, "suffixsuffixsuffixsuffix"...)

	fmt.Println("after path append=>", string(path), len(path), cap(path), &path[0])

	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	/*
	It didn't work as you expected. Instead of "AAAAsuffix/BBBBBBBBB" you ended up with "AAAAsuffix/uffixBBBB".
	It happened because both directory slices referenced the same underlying array data from the original path slice.
	This means that the original path is also modified. Depending on your application this might be a problem too.
	 */

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB (not ok)

	fmt.Println("new path =>", string(path))

	//This problem can fixed by allocating new slices and copying the data you need. Another option is to use the full slice expression.
}

func SliceAppendDemo() {
	slice1 := make([]int, 0, 10)
	slice1 = append(slice1, []int{1, 2, 3, 4, 5, 6, 7}...)
	fmt.Println(slice1, len(slice1), cap(slice1))
	//[1 2 3 4 5 6 7] 7 10
	_ = append(slice1, 8)
	fmt.Println(slice1, len(slice1), cap(slice1))
	//[1 2 3 4 5 6 7] 7 10

	/*
	理解：
	下面slice2 初始化指向一个 []int{1, 2, 3, 4, 5} len 5 cap 5
	slice3指向slice2其中一部分区域
	在append slice3，会对其指向区域中的数组(也是slice2指向的数组)进行元素修改
	如果append slice3的结果没有赋值给slice3，slice3并不会改变其指向的区域
	当增加的元素个数超过slice2指向区域，而slice2如果不被重新赋值，也仍然指向原来区域
	 */
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice2 len cap =>", len(slice2), cap(slice2))
	slice3 := slice2[:2]
	_ = append(slice3, 6, 7)
	fmt.Println(slice2, slice3)
	slice4 := append(slice3, 8, 9, 10, 11)
	fmt.Println(slice2, slice3, slice4)
}

func SliceDataCorruption2() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	dir1 := path[:sepIndex:sepIndex] //full slice expression
	dir2 := path[sepIndex + 1:]
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB

	/*
	The extra parameter in the full slice expression controls the capacity for the new slice.
	Now appending to that slice will trigger a new buffer allocation instead of overwriting the data in the second slice.
	 */

	dir1 = append(dir1, "suffix"...)
	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB (ok now)

	fmt.Println("new path =>", string(path))
}

//"Stale" Slices
/*
Multiple slices can reference the same data. This can happen when you create a new slice from an existing slice,
for example. If your application relies on this behavior to function properly then you'll need to worry about "stale" slices.

At some point adding data to one of the slices will result in a new array allocation
when the original array can't hold any more new data. Now other slices will point to the old array (with old data).
 */
func StaleSlices() {

	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) //prints 3 3 [1 2 3]

	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) //prints 2 2 [2 3]

	for i := range s2 {
		s2[i] += 20
	}

	//still referencing the same array
	fmt.Println(s1) //prints [1 22 23]
	fmt.Println(s2) //prints [22 23]

	s2 = append(s2, 4)

	for i := range s2 {
		s2[i] += 10
	}

	//s1 is now "stale"
	fmt.Println(s1) //prints [1 22 23]
	fmt.Println(s2) //prints [32 33 14]
}

//#Type Declarations and Methods
type myMutex sync.Mutex
type myLocker struct {
	sync.Mutex
}
type myLocker2 sync.Locker

func TypeDeclarationsAndMethods() {

	//var mtx myMutex
	//mtx.Lock() //error
	//mtx.Unlock() //error
	/*
	If you do need the methods from the original type
	 you can define a new struct type embedding the original type as an anonymous field.
	 */
	var lock myLocker
	lock.Lock() //ok
	lock.Unlock() //ok

	//Interface type declarations also retain their method sets.
	var lock2 myLocker2 = new(sync.Mutex)
	lock2.Lock() //ok
	lock2.Unlock() //ok
}



//# Breaking Out of "for switch" and "for select" Code Blocks
func BreakForSwitchSelect() {
	loop:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			break loop
		}
	}

	fmt.Println("out!")
}

//# Iteration Variables and Closures in "for" Statements
func IterationVariableandClosuresInforStatements() {
	data := []string{"one", "two", "three"}

	for _, v := range data {
		//vcopy := v //
		go func() {
			//fmt.Println(vcopy)
			fmt.Println(v)
		}()

		//go func(in string) {
		//	fmt.Println(in)
		//}(v)
	}

	time.Sleep(3 * time.Second)
	//goroutines print: three, three, three
}

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func IterationVariableandClosuresInforStatements2() {
	data := []field{{"one"}, {"two"}, {"three"}}

	for _, v := range data {
		fmt.Printf("%p\n", &v)
		go v.print()
		//v := v
		//go v.print()
	}

	time.Sleep(3 * time.Second)
	//goroutines print: three, three, three
}

func IterationVariableandClosuresInforStatements3() {
	data := []*field{{"one"}, {"two"}, {"three"}}
	// Redundant Type declare
	//data := []*field{&field{"one"}, &field{"two"}, &field{"three"}}

	/*
	此处当传入一个指针类型的数组的时候，for range 遍历到的元素v的地址就是数组元素本身的地址，并没有新建一个临时元素同时指向那个实际值
	go v.print()  与 go func() {v.print()}() 还是会不一样，因为这里还是会存在闭包，虽然v每一次的指向地址是不同的
	 */
	fmt.Println("data[0], data[1], data[2] => ", fmt.Sprintf("%p, %p, %p", data[0], data[1], data[2]))
	fmt.Println("&data[0], &data[1], &data[2] => ", fmt.Sprintf("%p, %p, %p", &data[0], &data[1], &data[2]))

	for _, v := range data {
		/*
		当data数组是一个指针类型数组，v也是该指针类型，
		 */
		fmt.Println("&v, v =>", fmt.Sprintf("%p %p", &v, v))
		//这里能够正常输出和go 这个语法应该有关
		go v.print()
		//goroutines print: one, two, three

		/*
		闭包的问题仍然是存在的
		虽然v是一个指针类型，但是循环中，该指针指向的区域，每次循环中在改变，而其v本身的地址(&v)并没有改变
		 */
		//go func() {
		//	v.print()
		//}()

	}

	time.Sleep(3 * time.Second)
}



//# Deferred Function Call Argument Evaluation
/*
Arguments for a deferred function call are evaluated
when the defer statement is evaluated (not when the function is actually executing).
 */
func DeferedFunctionCallArgmentEvaluation() {
	var i int = 1

	// func() int { return i * 2 }() 结果作为defer 方法的参数，在一开始就会被执行确定下来
	defer fmt.Println("result =>", func() int {
		return i * 2
	}())
	i++
	//prints: result => 2 (not ok if you expected 4)
}

//# Deferred Function Call Execution
/*
The deferred calls are executed at the end of the containing function and not at the end of the containing code block.
It's an easy mistake to make for new Go developers confusing the deferred code execution rules with the variable scoping rules.
It can become a problem if you have a long running function with a for loop that tries to defer resource cleanup calls in each iteration.

    for _,target := range targets {
        f, err := os.Open(target)
        if err != nil {
            fmt.Println("bad target:",target,"error:",err) //prints error: too many open files
            break
        }
        defer f.Close() //will not be closed at the end of this code block
        //do something with the file...
    }

One way to solve the problem is by wrapping the code block in a function.

    for _,target := range targets {
        func() {
            f, err := os.Open(target)
            if err != nil {
                fmt.Println("bad target:",target,"error:",err)
                return
            }
            defer f.Close() //ok
            //do something with the file...
        }()
    }
 */


/*
Advanced Beginner:
 */
//# Using Pointer Receiver Methods On Value Instances
/*
It's OK to call a pointer receiver method on a value as long as the value is addressable.
 In other words, you don't need to have a value receiver version of the method in some cases.

Not every variable is addressable though.
Map elements are not addressable. Variables referenced through interfaces are also not addressable.

 */
type printdata struct {
	name string
}

func (p *printdata) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func UsingPointerReceiverMethodsOnValueInstances() {
	d1 := printdata{"one"}
	d1.print() //ok

	/*
	Compile Errors:
	cannot use printdata literal (type printdata) as type printer in assignment:
	printdata does not implement printer (print method has pointer receiver)
	 */
	//var in printer = printdata{"two"} //error
	//in.print()

	/*
	Compile Errors:
	cannot call pointer method on m["x"]
	cannot take the address of m["x"]
	 */
	//m := map[string]printdata {"x":printdata{"three"}}
	//m["x"].print() //error
}

//# Updating Map Value Fields
type upmapdata struct {
	name string
}

func UpdateingMapValueFields() {
	/*
	Compile Error:
	cannot assign to struct field m["x"].name in map

	It doesn't work because map elements are not addressable.
	 */
	//m := map[string]upmapdata {"x":{"one"}}
	//m["x"].name = "two" //error

	//slice elements are addressable.
	s := []upmapdata{{"one"}}
	s[0].name = "two" //ok
	fmt.Println(s)    //prints: [{two}]

	//The first work around is to use a temporary variable.
	m := map[string]upmapdata{"x":{"one"}}
	r := m["x"]
	r.name = "two"
	m["x"] = r
	fmt.Printf("%v", m) //prints: map[x:{two}]

	//Another workaround is to use a map of pointers.
	m1 := map[string]*upmapdata{"x":{"one"}}
	m1["x"].name = "two" //ok
	fmt.Println(m1["x"]) //prints: &{two}

	//m2 := map[string]*upmapdata {"x":{"one"}}
	////panic: runtime error: invalid memory address or nil pointer dereference
	//m2["z"].name = "what?" //???


}


//# "nil" Interfaces and "nil" Interfaces Values
/*
This is the second most common gotcha in Go because interfaces are not pointers even though they may look like pointers.
Interface variables will be "nil" only when their type and value fields are "nil".

The interface type and value fields are
populated based on the type and value of the variable used to create the corresponding interface variable.
This can lead to unexpected behavior when you are trying to check if an interface variable equals to "nil".
 */
func NilInterfacesandNilInterfacesValues() {
	var data *byte
	var in interface{}

	fmt.Println(data, data == nil) //prints: <nil> true
	fmt.Println(in, in == nil)     //prints: <nil> true

	in = data
	fmt.Println(in, in == nil)     //prints: <nil> false
	//'data' is 'nil', but 'in' is not 'nil'


	doit := func(arg int) interface{} {
		var result *struct{} = nil

		//if (arg > 0) {
		//	result = &struct{}{}
		//}
		if(arg > 0) {
			result = &struct{}{}
		} else {
			return nil //return an explicit 'nil'
		}

		return result
	}

	if res := doit(-1); res != nil {
		fmt.Println("good result:", res) //prints: good result: <nil>
		//'res' is not 'nil', but its value is 'nil'
	}else {
		fmt.Println("bad result (res is nil)") //here as expected
	}
}


//# Stack and Heap Variables
/*
You don't always know if your variable is allocated on the stack or heap.
In C++ creating variables using the new operator always means that you have a heap variable.
In Go the compiler decides where the variable will be allocated even if the new() or make() functions are used.
The compiler picks the location to store the variable based on its size and the result of "escape analysis".
This also means that it's ok to return references to local variables, which is not ok in other languages like C or C++.

If you need to know where your variables are allocated pass the "-m" gc flag to "go build" or "go run" (e.g., go run -gcflags -m app.go).
 */

//# GOMAXPROCS, Concurrency, and Parallelism
//TODO:

	