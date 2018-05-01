package ivsample

import (
	"fmt"
	"runtime"
	"sync"
)

/*两个case io 都不阻塞情况下，
无论设置GOMAXPROCS,都会随机选择select的一条路径
所以有一定机会造成panic
 */
func ChannelSelect() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{}) // 既然是迭代就会要求set.s全部可以遍历一次。但是chan是为缓存的，那就代表这写入一次就会阻塞。 我们把代码恢复为可以运行的方式，看看效果
	//ch := make(chan interface{},len(set.s))
	go func() {
		set.RLock()

		for elem, value := range set.s {
			ch <- elem
			println("Iter:", elem, value)
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func ChannelIterator(){
	th:=threadSafeSet{
		s:[]interface{}{"1","2"},
	}
	v:=<-th.Iter()
	fmt.Sprintf("%s%v","ch",v)
}


/*
All operations on unbuffered channels block the execution until both sender and receiver are ready to communicate.
That’s why unbuffered channels are also called synchronous.
In case a channel has a buffer all read operations succeed without blocking if the buffer is not empty, and write operations
- if the buffer is not full. These channels are called asynchronous.
 */
func ChannelSyncDeadLock() {
	//fatal error: all goroutines are asleep - deadlock!
	c := make(chan int)
	c <- 42    // write to a channel
	val := <-c // read from a channel
	println(val)
}


/*
Closing a channel has one more useful feature
- reading operations on closed channels do not block and always return default value for a channel type
*/
func ChannelClosedUseful() {
	done := make(chan bool)

	go func() {
		println("goroutine message")

		// We are only interested in the fact of sending itself,
		// but not in data being sent.
		//done <- true


		/* Here the done channel is only used to synchronize the execution but not for sending data.
		There is a kind of pattern for such cases
		 */
		close(done)
	}()

	println("main function message")
	<-done
}
