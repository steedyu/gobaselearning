package gousualsample

import (
	"os"
	"io"
	"fmt"
)

func SeekSample() {
	file, err := os.Open("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	/*
	whence 的值，在 io 包中定义了相应的常量，应该使用这些常量
	原先 os 包中的常量已经被标注为Deprecated
	 */
	file.Seek(2,io.SeekStart)
	buf := make([]byte,50)
	n, err := file.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n, ":", string(buf))
}
