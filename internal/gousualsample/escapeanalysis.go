package gousualsample

import (
	"fmt"
	"time"
)

func EsAnalysisSample1() {
	s := "hello"
	fmt.Println(s)
}

type S struct{}

func ref(z S) *S {
	return &z
}

func EsAnalysisSample2() {
	var x S
	_ = *ref(x)
}

type StkContent struct {
	Price      string
	Timeticket int64
}

func UseStkIdContentMap(stkIdContentMap map[string]*StkContent) {
	tmpStkContent := StkContent{}
	for i := 0; i < 20; i++ {
		if v, ok := stkIdContentMap[fmt.Sprintf("%v",i)];ok {
			tmpStkContent.Price = v.Price
			tmpStkContent.Timeticket = v.Timeticket
			fmt.Printf("%+v \r\n",tmpStkContent)
		}
	}
}

func CompseStkIdContentMap() map[string]*StkContent {
	var stkIdContentMap map[string]*StkContent = make(map[string]*StkContent, 0)
	for i := 0; i < 20; i++ {
		stkIdContentMap[fmt.Sprintf("%v",i)] = &StkContent{
			Price:"11.1",
			Timeticket:time.Now().UnixNano(),
		}
	}
	return stkIdContentMap
}

func EsAnalysisSample3() {
	stkIdContentMap := CompseStkIdContentMap()
	UseStkIdContentMap(stkIdContentMap)
}

