package gousualsample

import (
	"fmt"
	"context"
)


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
