package socketsample

import (
	"time"
	"net"
	"log"
)


//串行服务端
func SocketServer1() {

	addr := "0.0.0.0:1212" //表示监听本地所有ip的8080端口，也可以这样写：addr := ":8080"
	listener, err := net.Listen("tcp", addr) //使用协议是tcp，监听的地址是addr
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close() //关闭监听的端口
	for {
		conn, err := listener.Accept() //用conn接收链接
		if err != nil {
			log.Fatal(err)
		}
		conn.Write([]byte("Yinzhengjie\n"))  //通过conn的wirte方法将这些数据返回给客户端。
		conn.Write([]byte("hello Golang\n"))
		time.Sleep(3 * time.Second) //在结束这个链接之前需要睡一分钟在结束当前循环。
		conn.Close() //与客户端断开连接。
	}
}

//并发服务端
func SocketServer2() {

	Handle_conn := func(conn net.Conn) {
		conn.Write([]byte("Yinzhengjie\n"))  //通过conn的wirte方法将这些数据返回给客户端。
		conn.Write([]byte("尹正杰是一个好男孩！\n"))
		time.Sleep(time.Second * 3)
		conn.Close() //与客户端断开连接。
	}

	addr := "0.0.0.0:1212" //表示监听本地所有ip的8080端口，也可以这样写：addr := ":8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept() //用conn接收链接
		if err != nil {
		}
		go Handle_conn(conn)  //开启多个协程。
	}
}

//web并发服务器
func SocketServer3() {

	var content = `

 <html>
 <head  name="尹正杰" age="25">  <!--标签的开头,其和面跟的内容（name="尹正杰"）是标签的属性,其属性可以定义多个。-->
     <meta charset="UTF-8"/>     <!--指定页面编码，-->
     <meta http-equiv="refresh" content="30; Url=http://www.cnblogs.com/yinzhengjie/"> <!--这是做了一个界面的跳转，表示30s不运行的话就跳转到指定的URL-->
     <title>尹正杰的个人主页</title> <!--定义头部的标题-->
 </head> <!--标签的结尾-->

 <body>
 <h1 style="color:red">尹正杰</h1>
 <h1 style="color:green">hello golang</h1>

 </body>
 </html>
 `

	Handle_conn := func(conn net.Conn) {
		conn.Write([]byte(content)) //将html的代码返回给客户端，这样客户端在web上访问就可以拿到指定字符。
		conn.Close()
	}

	addr := "0.0.0.0:1212" //表示监听本地所有ip的8080端口，也可以这样写：addr := ":8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept() //用conn接收链接
		if err != nil {
		}
		go Handle_conn(conn)  //开启多个协程。
	}
}