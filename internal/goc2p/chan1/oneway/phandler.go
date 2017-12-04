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
	origsCap := cap(origs)
	buffered := origsCap > 0
	goTicketTotal := origsCap / 2
	goTicket := initGoTicket(goTicketTotal)
	go func() {
		for {
			p, ok := fetchPerson1()
			if !ok {
				for {
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
				<-goTicket
				go func() {
					origs <- p
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


