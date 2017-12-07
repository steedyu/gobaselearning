package oneway

import (
	"fmt"
	"time"
)

type Person struct {
	Name    string
	Age     uint8
	Address Addr
}

type Addr struct {
	city     string
	district string
}

type PersonHandler interface {
	Batch(origs <-chan Person) <-chan Person
	Handle(origs *Person)
}

type PersonHandlerImpl struct {}

/*
1 基于通道通讯时在多个Goroutine之间进行同步的重要手段。通道既能被用来在多个Goroutine之间传递数据，又能够在数据传递过程起到同步的作用
2 在相关通道被关闭时，当此时的通道中海油一流元素值时，运行时系统会等for语句把他们全部接受后再结束该语句的执行
 */
func (handler PersonHandlerImpl) Batch(origs <-chan Person) <-chan Person{
	dests := make(chan Person,100)
	go func() {
		for p := range origs{
			handler.Handle(&p)
			dests <- p
		}
		fmt.Println("All the information has been handled.")
		close(dests)
	}()
	return dests
}

func (handler PersonHandlerImpl) Handle(orig *Person) {
	if orig.Address.district == "Haidian" {
		orig.Address.district = "Shijingshan"
	}
}

var personTotal = 200
var persons []Person = make([]Person, personTotal)
var personCount int

func init() {
	for i := 0; i < 200; i++ {
		name := fmt.Sprintf("%s%d", "P", i)
		p := Person{name, 32, Addr{"Beijing", "Haidian"}}
		persons[i] = p
	}
}

func PhandlerDemo() {
	handler := getPersonHandler()
	origs := make(chan Person, 100)
	dests := handler.Batch(origs)
	fetchPerson(origs)
	sign := savePerson(dests)
	<-sign
}

func getPersonHandler() PersonHandler {
	return PersonHandlerImpl{}
}

func savePerson(dest <-chan Person) <-chan byte {
	sign := make(chan byte, 1)
	go func() {
		for {
			p, ok := <-dest
			if !ok {
				fmt.Println("All the information has been saved.")
				sign <- 0
				break
			}
			savePerson1(p)
		}
	}()
	return sign
}

func fetchPerson(origs chan<- Person) {
	/*
	判断channel是否带有缓存
	非缓冲通道只能同步地传递元素值
	在收发两端都有并发需求的情况下，使用非缓冲通道作为元素值传输介质是不合适的
	  */
	origsCap := cap(origs)
	buffered := origsCap > 0
	/*
	goTicket通道实际上我们为了限制该程序启用的Goroutine的数量而声明的一个缓冲通道
	这是使用缓冲通道作为Goroutine票池的典型做法
	 */
	goTicketTotal := origsCap / 2
	goTicket := initGoTicket(goTicketTotal)
	go func() {
		for {
			p, ok := fetchPerson1()
			if !ok {
				for {
					/*
					保证安全的情况下，关闭origs
					1 buffered判断，如果origs通道是非缓冲的，我们没必要做检查直接关闭
					2 goTicket长度是否和其容量相等，相等，说明goTicket中令牌都已被收回，所有相关Goroutine都已经执行完毕，就可以关闭了
					 */
					if !buffered || len(goTicket) == goTicketTotal {
						break
					}
					time.Sleep(time.Nanosecond)
				}
				fmt.Println("All the information has been fetched.")
				close(origs)
				break
			}
			if buffered {
				/*
				每当我们要启用一个Goroutine的时候，就从该通道中接受一个元素值，以表示可被启用的Goroutine减少一个
				 */
				<-goTicket
				go func() {
					origs <- p
					/*
					每当一个被启用Goroutine的运行即将结束的时候，我们就应该向该通道发送一个元素值，以表示可被启用的Goroutine增加一个
					 */
					goTicket <- 1
				}()
			} else {
				origs <- p
			}
		}
	}()
}

func initGoTicket(total int) chan byte {
	var goTicket chan byte
	if total == 0 {
		return goTicket
	}
	goTicket = make(chan byte, total)
	for i := 0; i < total; i++ {
		goTicket <- 1
	}
	return goTicket
}

func fetchPerson1() (Person, bool) {
	if personCount < personTotal {
		p := persons[personCount]
		personCount++
		return p, true
	}
	return Person{}, false
}

func savePerson1(p Person) bool {
	return true
}


