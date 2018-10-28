package gousualsample

import (
	"fmt"
	"context"
	"time"
	"sync"
	"sync/atomic"
)


/*
genLeakGoroutine
NoLeadkGoroutineSample
NoLeakGoroutineSample2
演示了goroutine造成memoryleak 以及 使用context这个结构体解决这个问题
 */
// gen is a broken generator that will leak a goroutine.
func genLeakGoroutine() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
		}
	}()
	return ch
}

func LeakGoroutineSample() {
	// The call site of gen doesn't have a
	for n := range genLeakGoroutine() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

// gen is a generator that can be cancellable by cancelling the ctx.
func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return // avoid leaking of this goroutine when ctx is done.
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}

func NoLeadkGoroutineSample() {
	/*
	context.Context是一个interface类型，在下面这段代码中
	WithCancel中返回的是指针类型
	所以传入gen的ctx 是指向同一区域
	 */
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // make sure all paths cancel the context to avoid context leak

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}
}

var (
	wg sync.WaitGroup
)

func work(ctx context.Context) error {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Doing some work ", i)

		// we received the signal of cancelation in this channel
		case <-ctx.Done():
			fmt.Println("Cancel the context ", i)
			return ctx.Err()
		}
	}
	return nil
}

func NoLeakGoroutineSample2() {
	//这里设置了一个4s的timeout
	ctx, cancel := context.WithTimeout(context.Background(), 4 * time.Second)
	defer cancel()

	fmt.Println("Hey, I'm going to do some work")

	wg.Add(1)
	go work(ctx)
	wg.Wait()

	fmt.Println("Finished. I'm going home")
}


/*
trigger函数会不断地获取一个名叫count的变量的值，并判断该值是否与参数i的值相同。如果相同，那么就立即调用fn代表的函数，然后把count变量的值加1，最后显式地退出当前的循环。否则，我们就先让当前的 goroutine“睡眠”一个纳秒再进入下一个迭代。

注意，我操作变量count的时候使用的都是原子操作。这是由于trigger函数会被多个 goroutine 并发地调用，所以它用到的非本地变量count，就被多个用户级线程共用了。因此，对它的操作就产生了竞态条件（race condition），破坏了程序的并发安全性。
 */
func ControlGoroutineOrder() {

	var count uint32 = 0

	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}

	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}

	/*
	最后要说的是，因为我依然想让主 goroutine 最后一个运行完毕，所以还需要加一行代码。不过既然有了trigger函数，我就没有再使用通道。
	 */
	trigger(10, func(){})

}