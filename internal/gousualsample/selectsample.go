package gousualsample

import (
	"fmt"
	"time"
)

func AllCaseExpressionEvaluatedBeforeSelect() {
	var ch3 chan int
	var ch4 chan int
	var chs = []chan int{ch3, ch4}
	var numbers = []int{1,2,3,4,5}

	/*
	所有跟在case关键字右边的发送语句或接收语句中的通道表达式和元素表达式都会先被求值。
	求值的顺序时自上而下、从左到右。
	 */
	select {
	case getChan(0,chs) <- getNumber(2,numbers):
		fmt.Println("1th case is selected.")
	case getChan(1,chs) <- getNumber(3,numbers):
		fmt.Println("2nd case is selected.")
	default :
		fmt.Println("default")

	}
}

func getNumber(i int, numbers []int) int{
	fmt.Printf("numbers[%d]\n",i)
	return numbers[i]
}

func getChan(i int, chs []chan int) chan int {
	fmt.Printf("chs[%d]\n",i)
	return chs[i]
}

func SelectTimeOutSample() {
	ch11 := make(chan int,1000)

	var e int
	ok := true
	for  {
		select {
		case e,ok = <- ch11:
			if !ok {
				fmt.Println("End")
				break
			}else {
				fmt.Printf("%d\n",e)
			}
		case ok = <- func() chan bool {
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
