package gousualsample

import (
	"io"
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"bufio"
)

func IoReaderExample() {

}

/*
Read 将 len(p) 个字节读取到 p 中。它返回读取的字节数 n（0 <= n <= len(p)） 以及任何遇到的错误。即使 Read 返回的 n < len(p)，它也会在调用过程中使用 p 的全部作为暂存空间。
若一些数据可用但不到 len(p) 个字节，Read 会照例返回可用的数据，而不是等待更多数据。
当 Read 在成功读取 n > 0 个字节后遇到一个错误或 EOF (end-of-file)，它就会返回读取的字节数。它会从相同的调用中返回（非nil的）错误或从随后的调用中返回错误（同时 n == 0）。
一般情况的一个例子就是 Reader 在输入流结束时会返回一个非零的字节数，同时返回的 err 不是 EOF 就是 nil。无论如何，下一个 Read 都应当返回 0, EOF。

调用者应当总在考虑到错误 err 前处理 n > 0 的字节。这样做可以在读取一些字节，以及允许的 EOF 行为后正确地处理 I/O 错误。
也就是说，当 Read 方法返回错误时，不代表没有读取到任何数据。调用者应该处理返回的任何数据，之后才处理可能的错误。
io.EOF 变量的定义：var EOF = errors.New("EOF")，是 error 类型。
根据 reader 接口的说明，在 n > 0 且数据被读完了的情况下，返回的 error 有可能是 EOF 也有可能是 nil。

 */

func ReaderExample() {
	FOREND:
	for {
		readerMenu()

		var ch string
		fmt.Scanln(&ch)
		var (
			data []byte
			err error
		)
		switch strings.ToLower(ch) {
		case "1":
			fmt.Println("请输入不多于9个字符，以回车结束：")
			data, err = read(os.Stdin, 11)
		case "2":
			//file, err := os.Open(util.GetProjectRoot() + "/src/chapter01/io/01.txt")
			//if err != nil {
			//	fmt.Println("打开文件 01.txt 错误:", err)
			//	continue
			//}
			//data, err = ReadFrom(file, 9)
			//file.Close()
			fmt.Println("暂未实现！")
		case "3":
			data, err = read(strings.NewReader("from string"), 12)
		case "4":
			fmt.Println("暂未实现！")
		case "b":
			fmt.Println("返回上级菜单！")
			break FOREND
		case "q":
			fmt.Println("程序退出！")
			os.Exit(0)
		default:
			fmt.Println("输入错误！")
			continue
		}

		if err != nil {
			fmt.Println("数据读取失败，可以试试从其他输入源读取！")
		} else {
			fmt.Printf("读取到的数据是：%s\n", data)
		}
	}
}

func read(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func readerMenu() {
	fmt.Println("")
	fmt.Println("*******从不同来源读取数据*********")
	fmt.Println("*******请选择数据源，请输入：*********")
	fmt.Println("1 表示 标准输入")
	fmt.Println("2 表示 普通文件")
	fmt.Println("3 表示 从字符串")
	fmt.Println("4 表示 从网络")
	fmt.Println("b 返回上级菜单")
	fmt.Println("q 退出")
	fmt.Println("***********************************")
}

/*
ReadAt 从基本输入源的偏移量 off 处开始，将 len(p) 个字节读取到 p 中。它返回读取的字节数 n（0 <= n <= len(p)）以及任何遇到的错误。
当 ReadAt 返回的 n < len(p) 时，它就会返回一个 非nil 的错误来解释 为什么没有返回更多的字节。在这一点上，ReadAt 比 Read 更严格。
即使 ReadAt 返回的 n < len(p)，它也会在调用过程中使用 p 的全部作为暂存空间。若一些数据可用但不到 len(p) 字节，ReadAt 就会阻塞直到所有数据都可用或产生一个错误。 在这一点上 ReadAt 不同于 Read。
若 n = len(p) 个字节在输入源的的结尾处由 ReadAt 返回，那么这时 err == EOF 或者 err == nil。
若 ReadAt 按查找偏移量从输入源读取，ReadAt 应当既不影响基本查找偏移量也不被它所影响。
ReadAt 的客户端可对相同的输入源并行执行 ReadAt 调用。
 */
func ReaderAtSample() {
	reader := strings.NewReader("Go语言中文网")
	p := make([]byte, 6)
	n, err := reader.ReadAt(p, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d", p, n)
}

func ReaderFromSample() {
	file,err := os.Open("writeAt.txt")
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()
}

func ReaderFromSample2() {
	ioutil.ReadFile("writeAt.txt")

}

