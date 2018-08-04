package gousualsample

import (
	"fmt"
	"context"
	"time"
	"sync"
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
