package socketsample

import (
	"fmt"
	"reflect"
	"io"
	"log"
	"net"
	"bufio"
	"os"
)


//串行指定读取客户端返回内容大小
func SocketClient1() {

	addr := "wwww.baidu.com:80" //定义主机名
	conn, err := net.Dial("tcp", addr) //拨号操作，需要指定协议。
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("访问公网IP地址是：", conn.RemoteAddr().String()) /*获取“conn”中的公网地址。注意：最好是加上后面的String方法，因为他们的那些是不一样的哟·当然你打印的时候
   可以不加输出结果是一样的，但是你的内心是不一样的哟！*/
	fmt.Printf("客户端链接的地址及端口是：%v\n", conn.LocalAddr()) //获取到本地的访问地址和端口。
	fmt.Println("“conn.LocalAddr()”所对应的数据类型是：", reflect.TypeOf(conn.LocalAddr()))
	fmt.Println("“conn.RemoteAddr().String()”所对应的数据类型是：", reflect.TypeOf(conn.RemoteAddr().String()))
	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n")) //向服务端发送数据。用n接受返回的数据大小，用err接受错误信息。
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("向服务端发送的数据大小是:", n)

	buf := make([]byte, 1024) //定义一个切片的长度是1024。

	n, err = conn.Read(buf) //接收到的内容大小。

	if err != nil && err != io.EOF {
		//io.EOF在网络编程中表示对端把链接关闭了。
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n])) //将接受的内容都读取出来。
	conn.Close()  //断开TCP链接。

}


//按照指定大小循环读取
func SocketClient2() {
	addr := "wwww.baidu.com:80" //定义主机名
	conn, err := net.Dial("tcp", addr) //拨号操作，需要指定协议。
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn.RemoteAddr().String()) //最好是加上后面的String方法，因为他们的那些是不一样的哟·当然你打印的时候可以不加。
	fmt.Println(conn.LocalAddr())
	fmt.Println(reflect.TypeOf(conn.LocalAddr()))
	fmt.Println(reflect.TypeOf(conn.RemoteAddr().String()))
	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n")) //向服务端发送数据。用n接受返回的数据大小，用err接受错误信息。
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("写入的大小是:", n)
	buf := make([]byte, 10) //定义一个切片的长度是1024。
	for {
		n, err = conn.Read(buf) //接收到的内容大小。
		if err == io.EOF {
			conn.Close()
			break
		}
		fmt.Print(string(buf[:n]))
	}
	fmt.Println(string(buf[:n])) //将接受的内容都读取出来。
}

//按行读取
func SocketClient3() {
	addr := "wwww.baidu.com:80" //定义主机名
	conn, err := net.Dial("tcp", addr) //拨号操作，需要指定协议。
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn.RemoteAddr().String()) //最好是加上后面的String方法，因为他们的那些是不一样的哟·当然你打印的时候可以不加。
	fmt.Println(conn.LocalAddr())
	fmt.Println(reflect.TypeOf(conn.LocalAddr()))
	fmt.Println(reflect.TypeOf(conn.RemoteAddr().String()))
	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n")) //向服务端发送数据。用n接受返回的数据大小，用err接受错误信息。
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("写入的大小是:", n)

	r := bufio.NewReader(conn) //将这个链接（connection）包装以下。将conn的内容都放入r中，但是没有进行读取，让步我们一会对其进行操作。
	for {
		line, err := r.ReadString('\n') //将r的内容也就是conn的数据按照换行符进行读取。
		if err == io.EOF {
			conn.Close()
			break
		}
		fmt.Print(line)
	}
}

func SocketClient4() {

	addr := "wwww.baidu.com:80" //定义主机名
	conn, err := net.Dial("tcp", addr) //拨号操作，需要指定协议。
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("访问公网IP地址以及端口是：", conn.RemoteAddr().String()) /*获取“conn”中的公网地址。注意：最好是加上后面的String方法，因为他们的那些是不一样的哟·当然你打印的时候
	27     可以不加输出结果是一样的，但是你的内心是不一样的哟！*/
	fmt.Printf("客户端链接的地址及端口是：%v\n", conn.LocalAddr()) //获取到本地的访问地址和端口。
	fmt.Println("“conn.LocalAddr()”所对应的数据类型是：", reflect.TypeOf(conn.LocalAddr()))
	fmt.Println("“conn.RemoteAddr().String()”所对应的数据类型是：", reflect.TypeOf(conn.RemoteAddr().String()))
	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n")) //向服务端发送数据。用n接受返回的数据大小，用err接受错误信息。
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("写入的大小是:", n)
	io.Copy(os.Stdout, conn)
	conn.Close()
}
