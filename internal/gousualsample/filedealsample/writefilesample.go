package filedealsample

import (
	"os"
	"log"
	"fmt"
	"io/ioutil"
)

//1 打开一个文件，如果没有就创建，如果有这个文件就清空文件内容（相当于python中的"w"）
func FileWriteW() {
	f, err := os.Create("D:\\GoPath\\src\\jerome.com\\gobaselearning\\internal\\gousualsample\\filedealsample\\wa.txt") //姿势一：打开一个文件，如果没有就创建，如果有这个文件就清空文件内容,需要用两个变量接受相应的参数
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("jerome\n") //往文件写入相应的字符串。
	f.Close()
}

//2 以追加的方式打开一个文件（相当于python中的"a"）
func FileWriteA() {
	f, err := os.OpenFile("wa.txt", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0644) //表示最佳的方式打开文件，如果不存在就创建，打开的模式是可读可写，权限是644
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("jerome\n")
	f.Close()
}

//3 修改文件内容-随机写入
func FileWriteSeek() {
	f, err := os.OpenFile("wa.txt", os.O_CREATE | os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("jerome\n")
	f.Seek(1, os.SEEK_SET) //表示文件的其实位置，从第二个字符往后写入。
	f.WriteString("$$$")
	f.Close()
}

//4.ioutil方法创建文件
func FileWriteByioutilcopy() {

	FileName := "D:\\GoPath\\src\\jerome.com\\gobaselearning\\internal\\gousualsample\\filedealsample\\我有一只小毛驴.txt"
	OutputFile := "D:\\GoPath\\src\\jerome.com\\gobaselearning\\internal\\gousualsample\\filedealsample\\复制的小毛驴.txt"

	buf, err := ioutil.ReadFile(FileName) //将整个文件的内容读到一个切片中。
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	//fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(OutputFile, buf, 0x644)    //我们将读取到的内容又重新写入到另外一个OutputFile文件中去。
	if err != nil {
		panic(err.Error())
	}

	/*注意，在执行该代码之后，就会生成一个OutputFile文件，其内容和FileName的内容是一致的哟！*/
}
