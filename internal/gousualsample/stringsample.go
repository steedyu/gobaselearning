package gousualsample

import (
	"fmt"
	"unicode/utf8"
	"time"

	"bytes"
	"math/rand"
	"strings"
	"strconv"
)

const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func StringSample1() {

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)

}

func StringSample2() {

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%q ", sample[i])
	}
}

const placeOfInterest = `⌘`

func StringSample3() {

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")

	fmt.Printf("DecodeRune: ")
	r, s := utf8.DecodeRuneInString(placeOfInterest)
	fmt.Printf("rune:%d(十六进制:%x), size:%d", r, r, s)
	fmt.Printf("\n")

}

const nihongo = "日本語"
const nihongoSepcial = "日本\xbd語"

func StringForRangeSample() {
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	for index, runeValue := range nihongoSepcial {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}

func UnicodeUtf8Sample() {
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}

func StringByteRuneSample() {
	hello := "Hello,世界"

	fmt.Println("------------for string-----------------------")
	count := 0
	for i := range hello {
		count++
		fmt.Printf("Time:%d,Index:%d -> %v\n", count, i, string(hello[i]))
	}

	fmt.Println("------------forrange string-----------------------")
	count = 0
	for i, c := range hello {
		count++
		fmt.Printf("Time:%d,Index:%d -> %v\n", count, i, string(c))

	}

	fmt.Println("-------------------unicode/utf8---------------------")
	s := []byte(hello)
	fmt.Println(utf8.RuneCount(s))
	// RuneCount returns the number of runes in p
	/*
	这里最后一个界打印不出来，因为界是最后一个rune，故runecount = 1
	 */
	for utf8.RuneCount(s) > 1 {
		r, size := utf8.DecodeRune(s)    //解码 s 中的第一个字符，返回解码后的字符和 p 中被解码的字节数
		fmt.Printf("string:%v -> %x ->charcter length:%d;", string(r), r, size)

		s = s[size:]
		nextR, size := utf8.DecodeRune(s)
		fmt.Printf("r == nextR:%v \n", r == nextR)
	}
}


func GenRandStringDemo() {
	fmt.Println(genRandString())
}

func genRandString() string {
	var buff bytes.Buffer
	var prev string
	var curr string
	for i := 0; buff.Len() < 3; i++ {
		//string可以将数字直接转为数字对应Ascii的字符
		curr = string(genRandAZAscii())
		if curr == prev {
			continue
		}
		prev = curr
		buff.WriteString(curr)
	}
	return buff.String()
}

func genRandAZAscii() int {
	min := 65 // A
	max := 90 // Z
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func MultilineStrings() {

	/*
	``的功能类似于Net中的@""
	 */
	str := `This is a\a
multiline
string.`

	str1 := "a\\ab"
	fmt.Println(str)
	fmt.Println(str1)
}

func EfficientConcatenationStrings() {

	var b bytes.Buffer

	for i := 0; i < 1000; i++ {
		b.WriteString(genRandString())
	}

	fmt.Println(b.String())

	//join来连接字符串时，需要字符提前已经生成
	var strs []string
	for i := 0; i < 1000; i++ {
		strs = append(strs, genRandString())
	}

	fmt.Println(strings.Join(strs, ""))
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
func GenRandomString(length int) string {

	var source = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)

}

func StringUsualFunctionDemo() {

	//"hellogo"中是否包含"hello", 包含返回true， 不包含返回false
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("hellogo", "abc"))

	//Joins 组合
	s := []string{"abc", "hello", "mike", "go"}
	buf := strings.Join(s, "x")
	fmt.Println("buf = ", buf)

	//Index, 查找子串的位置
	fmt.Println(strings.Index("abcdhello", "hello"))
	fmt.Println(strings.Index("abcdhello", "go")) //不包含子串返回-1

	buf = strings.Repeat("go", 3)
	fmt.Println("buf = ", buf) //"gogogo"

	//Split 以指定的分隔符拆分
	buf = "hello@abc@go@mike"
	s2 := strings.Split(buf, "@")
	fmt.Println("s2 = ", s2)

	//Trim去掉两头的字符
	buf = strings.Trim("      are u ok?          ", " ") //去掉2头空格
	fmt.Printf("buf = #%s#\n", buf)

	//去掉空格，把元素放入切片中
	s3 := strings.Fields("      are u ok?          ")
	//fmt.Println("s3 = ", s3)
	for i, data := range s3 {
		fmt.Println(i, ", ", data)
	}

}

/*
fmt.Errorf("%s", "this is normal err")   等同于 errors.New(fmt.Sprintf("%s", "this is normal err"))
 */

/*
严格的讲，应该是在把 int，float等类型转换为字符串时，不要用 fmt.Sprintf，更好的做法是用标准库函数。fmt.Sprintf 的用途是格式化字符串，接受的类型是
 interface{}，内部使用了反射。所以，与相应的标准库函数相比，fmt.Sprintf 需要更大的开销。大多数类型转换的函数都可以在 strconv 包里找到。
 */
func StringConvertDemo() {

	//转换为字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	//第二个数为要追加的数，第3个为指定10进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abcgohello")

	fmt.Println("slice = ", string(slice)) //转换string后再打印

	//其它类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	//'f' 指打印格式，以小数方式， -1指小数点位数(紧缩模式)， 64以float64处理
	str = strconv.FormatFloat(3.14, 'f', -1, 64)

	//整型转字符串，常用
	str = strconv.Itoa(6666)
	fmt.Println("str = ", str)

	//字符串转其它类型
	var flag bool
	var err error
	flag, err = strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag = ", flag)
	} else {
		fmt.Println("err = ", err)
	}

	//把字符串转换为整型
	a, _ := strconv.Atoi("567")
	fmt.Println("a = ", a)
}


