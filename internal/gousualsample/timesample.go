package gousualsample

import (
	"fmt"
	"time"
)

func TimeSample() {

	//获取当前时间
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.UTC())
	fmt.Println(now.Date())
	fmt.Println(now.Hour(), now.Minute(), now.Second(), now.Nanosecond())

	//获取当前时间戳  注意这里是UTC时间为基准计算的
	fmt.Println(time.Now().Unix())                //the number of seconds elapsed since January 1, 1970 UTC.
	fmt.Println(time.Now().UnixNano())            //the number of nanoseconds elapsed since January 1, 1970 UTC.
	fmt.Println(time.Now().UnixNano() / 1e6)      //将纳秒转换为毫秒
	fmt.Println(time.Now().UnixNano() / 1e9)      //将纳秒转换为秒

	//从时间戳转换为日期时间
	c1 := time.Unix(now.Unix(),0)
	c2 := time.Unix(now.UnixNano()/1e9, 0)  //将毫秒转换为 time 类型
	fmt.Println(c1)
	fmt.Println(c2)

	c5 := time.Unix(0, now.UnixNano()) //将毫秒转换为 time 类型
	fmt.Println(c5.String())                      //输出当前英文时间戳格式
	c6 := time.Unix(now.Unix(), now.UnixNano()) //将毫秒转换为 time 类型
	fmt.Println(c6.String())                      //输出当前英文时间戳格式



	date := time.Now()
	time.Sleep(10 * time.Second)
	span := time.Since(date)
	fmt.Println(span.Hours(), span.Minutes(), span.Seconds(), span.Nanoseconds(), span.String())


}
