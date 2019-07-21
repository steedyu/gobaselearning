package gousualsample

import (
	"fmt"
	"time"
	"math/rand"
)

func AllCaseExpressionEvaluatedBeforeSelect() {

	/*
	chs[0]
	numbers[2]
	chs[1]
	numbers[3]
	default

	下方代码，只会出现default结果
	 */
	//var ch3 chan int
	//var ch4 chan int

	/*
	chs[0]
	numbers[2]
	chs[1]
	numbers[3]
	1th case is selected.

	chs[0]
	numbers[2]
	chs[1]
	numbers[3]
	2nd case is selected.

	下方代码两种结果都有可能出现
	*/
	var ch3 chan int = make(chan int, 1)
	var ch4 chan int = make(chan int, 1)
	var chs = []chan int{ch3, ch4}
	var numbers = []int{1, 2, 3, 4, 5}

	/*
	所有跟在case关键字右边的发送语句或接收语句中的通道表达式和元素表达式都会先被求值。
	求值的顺序时自上而下、从左到右。
	 */
	select {
	case getChan(0, chs) <- getNumber(2, numbers):
		fmt.Println("1th case is selected.")
	case getChan(1, chs) <- getNumber(3, numbers):
		fmt.Println("2nd case is selected.")
	default:
		fmt.Println("default")

	}
}

func getNumber(i int, numbers []int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int, chs []chan int) chan int {
	fmt.Printf("chs[%d]\n", i)
	return chs[i]
}

func SelectTimeOutSample() {
	ch11 := make(chan int, 1000)

	var e int
	ok := true
	for {
		select {
		case e, ok = <-ch11:
			if !ok {
				fmt.Println("End")
				break
			} else {
				fmt.Printf("%d\n", e)
			}
		case ok = <-func() chan bool {
			timeout := make(chan bool, 1)

			go func() {
				time.Sleep(time.Millisecond)
				timeout <- false
			}()
			return timeout
		}():
			fmt.Println("Timeout.")
			break

		}
		if !ok {
			break
		}
	}
}

func SelectTimeOutSample2() {
	ch11 := make(chan int, 1000)

	var e int
	ok := true
	//to := time.NewTimer(time.Millisecond)
	var timer *time.Timer
	for {
		//to.Reset(time.Millisecond) //这样做的确可以不重复创建timer 但是这个超时的时间，其实据不精确了
		select {
		case e, ok = <-ch11:
			if !ok {
				fmt.Println("End")
				break
			} else {
				fmt.Printf("%d\n", e)
			}
		//case <-to.C:
		case <-func() <-chan time.Time {
			//这样做能够精确 因为在 执行到case的时候 才会去初始化或者重置
			if timer == nil {
				timer = time.NewTimer(time.Millisecond)
			} else {
				timer.Reset(time.Millisecond)
			}
			return timer.C
		}():
			fmt.Println("Timeout.")
			break
		}
		if !ok {
			break
		}
	}
}

func SelectSample() {
	selectexample1()
	selectexample2()
}


// 示例1。
func selectexample1() {
	// 准备好几个通道。
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	// 随机选择一个通道，并向它发送元素值。
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index
	// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
	select {
	case <-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	case elem := <-intChannels[2]:
		fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
	default:
		fmt.Println("No candidate case is selected!")
	}
}

// 示例2。
func selectexample2() {
	intChan := make(chan int, 1)
	// 一秒后关闭通道。
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("The candidate case is closed.")
			break
		}
		fmt.Println("The candidate case is selected.")
	}
}
