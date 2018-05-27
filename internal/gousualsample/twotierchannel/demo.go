package twotierchannel

import (
	"sync"
	"time"
)

func TwoTierChannelSystemDemo() {
	/*
	其实这个例子，是控制了同时处理job的goroutine数量，以及控制jobqueue队列长度
	从代码的理解上看，它并不是解决速度跟不上的问题，而是解决突然间流量增大，而不会开启很多的goroutine来处理，使得一瞬间资源耗尽
	控制这些goroutine和channel

	We have decided to utilize a common pattern when using Go channels, in order to create a 2-tier channel system,
	one for queuing jobs and another to control how many workers operate on the JobQueue concurrently.

	具体实现方式
	一个队列JobQueue 接受提交的请求分配到该channel

	另外一个dispatcher 和 worker
	dispatcher中规定了 启动多少个 worker 用于处理JobQueue
	设计上：dispatcher中的　workpool 作为一个总的 job channel 的 pool ，实际是用channel类型的channel来实现的， 从pool中取出一个channel 放入元素
	而worker中也有一个field引用了workpool，  worker中的处理方式是将一个非缓冲channel放入该pool中，然后 监控该channel有数据读出便处理
	而dispatcher中则从workpool中获取一个channel，将job 推送到该channel，那woker就可以接收并处理，
	两个组件通过workpool来进行衔接

	其实这里思考一下，如果只用一个channel在各个worker中进行循环效果也一样，因为开启worker数目当中并不会变化，没有用到的也是当场阻塞

	 */

	var max int = 2

	JobQueue = make(chan Job, 10)

	dispatcher := NewDispatcher(max)
	dispatcher.Run()

	for i := 0; i < 11; i++ {
		job := Job{Flag:i}
		job.Payload = Payload{}
		JobQueue <- job
	}


	time.Sleep(time.Minute * 1)
}

type WaitWrapper struct {
	sync.WaitGroup
}

func (w *WaitWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		defer w.Done()
		cb()
	}()
}





