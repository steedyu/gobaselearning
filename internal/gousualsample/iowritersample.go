package gousualsample

import (
	"io"
	"os"
	"fmt"
	"bytes"
)

func write(writer io.Writer, data []byte) (num int, err error) {
	return writer.Write(data)
}

func StdOutWrite() {
	write(os.Stdout,[]byte("Hello World"))
}

func WriterAtSample() {
	file, err := os.Create("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Golang中文社区——这里是多余的")
	n, err := file.WriteAt([]byte("Go语言中文网"), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

func WriteToSample() {
	reader := bytes.NewReader([]byte("Go语言中文网"))
	reader.WriteTo(os.Stdout)
}

