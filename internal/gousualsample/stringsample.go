package gousualsample

import (
	"fmt"
	"unicode/utf8"
	"time"

	"bytes"
	"math/rand"
	"strings"
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


