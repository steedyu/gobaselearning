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


func CondSample1() {
	// mailbox 代表信箱。
	// 0代表信箱是空的，1代表信箱是满的。
	var mailbox uint8
	// lock 代表信箱上的锁。
	var lock sync.Mutex
	// sendCond 代表专用于发信的条件变量。
	sendCond := sync.NewCond(&lock)
	// recvCond 代表专用于收信的条件变量。
	recvCond := sync.NewCond(&lock)

	// send 代表用于发信的函数。
	send := func(id, index int) {
		lock.Lock()
		for mailbox == 1 {
			/*
			条件变量的Wait方法主要做了四件事。
			1 把调用它的 goroutine（也就是当前的 goroutine）加入到当前条件变量的通知队列中。
			2 解锁当前的条件变量基于的那个互斥锁。
			3 让当前的 goroutine 处于等待状态，等到通知到来时再决定是否唤醒它。此时，这个 goroutine 就会阻塞在调用这个Wait方法的那行代码上。
			4 如果通知到来并且决定唤醒这个 goroutine，那么就在唤醒它之后重新锁定当前条件变量基于的互斥锁。自此之后，当前的 goroutine 就会继续执行后面的代码了。
			 */
			sendCond.Wait()
		}
		fmt.Printf("sender [%d-%d]: the mailbox is empty.",
			id, index)
		mailbox = 1
		fmt.Printf("sender [%d-%d]: the letter has been sent.",
			id, index)
		lock.Unlock()
		recvCond.Broadcast()
	}

	// recv 代表用于收信的函数。
	recv := func(id, index int) {
		lock.Lock()
		for mailbox == 0 {
			recvCond.Wait()
		}
		fmt.Printf("receiver [%d-%d]: the mailbox is full.",
			id, index)
		mailbox = 0
		fmt.Printf("receiver [%d-%d]: the letter has been received.",
			id, index)
		lock.Unlock()
		sendCond.Signal() // 确定只会有一个发信的goroutine。
	}

	// sign 用于传递演示完成的信号。
	sign := make(chan struct{}, 3)
	max := 6
	go func(id, max int) { // 用于发信。
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			send(id, i)
		}
	}(0, max)
	go func(id, max int) { // 用于收信。
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max; j++ {
			time.Sleep(time.Millisecond * 200)
			recv(id, j)
		}
	}(1, max/2)
	go func(id, max int) { // 用于收信。
		defer func() {
			sign <- struct{}{}
		}()
		for k := 1; k <= max; k++ {
			time.Sleep(time.Millisecond * 200)
			recv(id, k)
		}
	}(2, max/2)

	<-sign
	<-sign
	<-sign
}