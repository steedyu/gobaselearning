package gousualsample

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

func OnceDoDemo() {
	var num int
	sign := make(chan bool)
	var once sync.Once
	f := func(ii int) func() {
		return func() {
			num = (num + ii * 2)
			sign <- true
		}
	}
	for i := 0; i < 3; i++ {
		fi := f(i + 1)
		go once.Do(fi())
	}
	for j := 0; j < 3; j++ {
		select {
		case <-sign:
			fmt.Println()
		case <-time.After(100 * time.Millisecond):
			fmt.Println("Timeout!")
		}
	}
	fmt.Println("Num:%d. \n", num)
}


/*
首先，我们不能对通过Get方法获取到的对象值有任何假设，都应该是无状态的或者状态一致的
临时对象池中的仍和对象值都有可能在任何时候被移除掉，并且根本不会通知该池的使用方。

临时对象池的一些适用场景（比如作为临时且状态无关的数据暂存处），以及一些不适用的场景（比如用来存放数据库连接的实例）
 */
func PoolDemo() {

	// 禁用GC，并保证在main函数执行结束前恢复GC
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var count int32
	newFunc := func() interface{} {
		return atomic.AddInt32(&count, 1)
	}
	pool := sync.Pool{New: newFunc}

	// New 字段值的作用
	v1 := pool.Get()
	fmt.Printf("v1: %v\n", v1)

	// 临时对象池的存取
	pool.Put(newFunc())
	pool.Put(newFunc())
	pool.Put(newFunc())
	v2 := pool.Get()
	fmt.Printf("v2: %v\n", v2)

	// 垃圾回收对临时对象池的影响
	debug.SetGCPercent(100)
	runtime.GC()
	v3 := pool.Get()
	fmt.Printf("v3: %v\n", v3)
	pool.New = nil
	v4 := pool.Get()
	fmt.Printf("v4: %v\n", v4)
}