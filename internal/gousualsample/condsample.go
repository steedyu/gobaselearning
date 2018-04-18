package gousualsample

import (
	"sync"
	"fmt"
	"time"
)

func CondDemo() {
	/*
	Cond在Locker的基础上增加的一个消息通知的功能。但是它只能按照顺序去使一个goroutine解除阻塞。

	Cond中内部是一个带头尾指针队列的数据结构
	 */

	var wait sync.WaitGroup
	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)

	for i := 0; i < 3; i++ {
		go func(i int) {
			wait.Add(1)
			defer wait.Done()
			cond.L.Lock()
			fmt.Println("Waiting start...", i)
			/*
			在调用wait方法内部，会先检查cond是否被复制
			加入通知列表
			释放锁
			等待通知
			被通知了，获取锁，继续运行
			 */
			cond.Wait()
			fmt.Println("Waiting end...")
			cond.L.Unlock()

			fmt.Println("Goroutine run. Number:", i)

		}(i)
	}

	/*
	每执行一次Signal就会执行一个goroutine
	通知的顺序是和加进去的顺序一致的，cond 内部最后调用的　runtime/sema.go 中的方法
	 */
	for i := 0; i < 3; i++ {
		time.Sleep(2e9)
		cond.L.Lock()
		cond.Signal()
		cond.L.Unlock()
	}

	/*
	如果想让所有的goroutine执行，那么将所有的Signal换成一个Broadcast方法可以。
	 */
	//time.Sleep(2e9)
	//cond.L.Lock()
	//cond.Broadcast()
	//cond.L.Unlock()

	wait.Wait()

}