package gousualsample

import (
	"fmt"
	"sync"
	"log"
	"time"
	"math/rand"
	"strconv"
	"runtime"
)

/*
其实这个例子比较特殊，因为它是一个无缓冲的channel，并且没有给传递值进去
如果没有关闭，它是会阻塞，读不出值的

在后面的多个Sender和一个receiver中就可以使用这个作为停止channel，由receiver关闭停止channel，
然后在Sender中进行判断，
 */
func isChanelClosed(ch <-chan int) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}

func CheckChannelCloses() {
	c := make(chan int)
	fmt.Println(isChanelClosed(c)) // false
	close(c)
	fmt.Println(isChanelClosed(c)) // true
}

func OneSenderMultiReceiversCloseChannelDemo() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 100000
	const NumReceivers = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int, 100)

	// the sender
	go func() {
		for {
			if value := rand.Intn(MaxRandomNumber); value == 0 {
				// the only sender can close the channel safely.
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			// receive values until dataCh is closed and
			// the value buffer queue of dataCh is empty.
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}

func MultiSendersOneReceiverCloseChannelDemo() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 100000
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the receiver of channel dataCh.
	// Its reveivers are the senders of channel dataCh.

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				value := rand.Intn(MaxRandomNumber)

				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}()
	}

	// the receiver
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == MaxRandomNumber - 1 {
				// the receiver of the dataCh channel is
				// also the sender of the stopCh cahnnel.
				// It is safe to close the stop channel here.
				close(stopCh)
				return
			}

			log.Println(value)
		}
	}()

	// ...
	wgReceivers.Wait()
}

func MultiSendersMultiReceiversCloseChannelDemo() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 100000
	const NumReceivers = 10
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the moderator goroutine shown below.
	// Its reveivers are all senders and receivers of dataCh.

	/*
	请注意channel toStop的缓冲大小是1.这是为了避免当mederator goroutine 准备好之前第一个通知就已经发送了，导致丢失。
	 */
	toStop := make(chan string, 1)
	// the channel toStop is used to notify the moderator
	// to close the additional signal channel (stopCh).
	// Its senders are any senders and receivers of dataCh.
	// Its reveiver is the moderator goroutine shown below.

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop // part of the trick used to notify the moderator
		// to close the additional signal channel.
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(MaxRandomNumber)
				if value == 0 {
					// here, a trick is used to notify the moderator
					// to close the additional signal channel.
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				// the first select here is to try to exit the
				// goroutine as early as possible.
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			for {
				// same as senders, the first select here is to
				// try to exit the goroutine as early as possible.
				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case value := <-dataCh:
					if value == MaxRandomNumber - 1 {
						// the same trick is used to notify the moderator
						// to close the additional signal channel.
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}

					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}

func ReciveMsgFromChan() {

	//从chan中接受值的两个方式，第二个方式可以判断，chan是否关闭
	var chan1 chan int = make(chan int, 1)
	chan1 <- 1

	value1 := <-chan1
	fmt.Println("receive result", value1)
	close(chan1)

	value, ok := <-chan1
	fmt.Println("receive result", value, ok)

	//接受未初始化的chan，会导致永久堵塞
	var chan2 chan int
	go func() {
		fmt.Println("here start")
		v3 := <-chan2
		fmt.Println("here end", v3)
	}()
	runtime.Gosched()

	go func() {
		fmt.Println("make chan2 start")
		chan2 = make(chan int, 1)
		chan2 <- 1
		fmt.Println("make chan2 end")
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("end")
}

func ReciveMsgFromChanIsCompletelyCopy() {
	type Addr struct {
		city     string
		district string
	}
	type Person struct {
		Name    string
		Age     uint8
		Address Addr
	}
	/*
	在发送过程中，进行的元素值复制并非浅表复制，而属于完全复制。者也保证了我们使用通道传递的值不变性
	 */

	var personChan = make(chan Person, 1)
	p1 := Person{"Harry", 32, Addr{"Beijing", "Haidian"}}
	fmt.Printf("p1 (1):%v\n", p1)
	personChan <- p1
	p1.Address.district = "Shijingshan"
	fmt.Printf("p1 (2):%v\n", p1)
	p1_copy := <-personChan
	fmt.Printf("p1_copy :%v\n", p1_copy)
}

func CloseChanDemo() {
	ch := make(chan int, 5)
	sign := make(chan byte, 2)
	go func() {
		for i := 0; i < 5; i ++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
		close(ch)
		fmt.Println("The channel is closed.")
		sign <- 0
	}()
	go func() {
		for {
			/*
			运行时系统并没有在通道ch被关闭之后立即把false作为相应接受操作的第二个结果，
			而是等到接受端把自己在通道中的所有元素都接受到之后才这样做。这确保了在发送端关闭通道的安全性。
			调用close函数只是让相应的通道进入关闭状态而不是立即组织对它的一切操作
			 */
			e, ok := <-ch
			fmt.Printf("%d (%v)\n", e, ok)
			if !ok {
				break
			}
			time.Sleep(2 * time.Second)
		}
		fmt.Println("Done")
		sign <- 1
	}()
	<-sign
	<-sign
}