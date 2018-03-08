package main

import (
	//"fmt"
	//"math"
	"jerome.com/gobaselearning/internal/gousualsample/encryption"
	"jerome.com/gobaselearning/internal/gousualsample"
)

func main() {
	//ivsample.DeferCallOrder()
	//ivsample.ForrangePointer()
	//ivsample.GoroutineClosure()
	//ivsample.EmbeddedStructMethod()
	//ivsample.ChannelSelect()
	//ivsample.DeferParamterMethod()
	//ivsample.DeferParamterMethod2()
	//ivsample.Sliceappend()
	//ivsample.UseStructImplementInterface()
	//ivsample.UseStructImplementInterface2()
	//ivsample.UseStructImplementInterface3()
	//ivsample.ChannelIterator()
	//ivsample.SwitchJudgeType()
	//ivsample.ConvertInterfaceToStringDemo()
	//ivsample.DeferInMethodWithReturnValue()
	//ivsample.SliceNew()
	//ivsample.Sliceappend2()
	//ivsample.CompareStruct()
	//ivsample.InterfaceNilIssue()
	//ivsample.InterfaceNilIssue1()
	//ivsample.InterfaceNilIssue2()
	//ivsample.InterfaceNilIssue3()
	//ivsample.StringNil()
	//ivsample.IotaCase()
	//ivsample.GotoLoop()
	//ivsample.VariableRange()
	//ivsample.ClosureCase1()
	//ivsample.ClosureCase2()
	//ivsample.PanicCase1()
	//ivsample.PanicCase2()

	//n := littleexercise.ReverseIntNum(-124)
	//fmt.Println(n)
	//n = littleexercise.ReverseIntNum(124)
	//fmt.Println(n)

	//fmt.Println(gousualsample.IsBid("111a"))
	//fmt.Println(gousualsample.IsBid("e17f36d12932672e60bb8605ee76b371"))
	//fmt.Println(gousualsample.IsBid("11713611293267216011860511761371"))
	//gousualsample.Xx()

	//gousualsample.ReaderExample()
	//gousualsample.StdOutWrite()
	//gousualsample.ReaderAtSample()
	//gousualsample.WriterAtSample()
	//gousualsample.ReaderFromSample()

	//gousualsample.StringSample1()
	//gousualsample.StringSample2()
	//gousualsample.StringSample3()
	//gousualsample.StringForRangeSample()
	//gousualsample.StringByteRuneSample()


	//apipe.Demo1(false)
	//apipe.Demo1(true)

	//npipe.InMemorySyncPipe()
	//npipe.FileBasedPipe()

	//socket.SocketDemo()
	//oneway.PhandlerDemo()


	//gousualsample.CheckChannelCloses()
	//gousualsample.ReciveMsgFromChan()
	//gousualsample.ReciveMsgFromChanIsCompletelyCopy()
	//gousualsample.CloseChanDemo()

	//gousualsample.AllCaseExpressionEvaluatedBeforeSelect()
	//gousualsample.SelectTimeOutSample()
	//gousualsample.SelectTimeOutSample2()

	//ivsample.ForrangeAppend()
	//gousualsample.StructZeroUsefulSample()

	//gousualsample.TimeSample()

	//gousualsample.Base64Sample()

	gousualsample.EsAnalysisSample3()

	//var byteArr []byte = make([]byte,0)
	//
	//byS1 := "115 115 58 49 49 50 55 48 57 52 51 52 48 50 52 48 48 55 52 58 52 54 54 55 50 50 57 56 58 115 116 111 99 107 115 111 114 116 58 49"
	//bys1Arr := strings.Split(byS1, " ")
	//for _, item := range bys1Arr {
	//	n,_ := strconv.Atoi(item)
	//	byteArr = append(byteArr, byte(n))
	//}
	//fmt.Println(string(byteArr))
	//
	//byS1 = "54 48 48 54 56 56 124 48 49 124 48 49"
	//bys1Arr = strings.Split(byS1, " ")
	//for _, item := range bys1Arr {
	//	n,_ := strconv.Atoi(item)
	//	byteArr = append(byteArr, byte(n))
	//}
	//fmt.Println(string(byteArr))

	//str1 := "601313ST,601360ST,601361ST"
	//strArr1 := strings.Split(str1, "601313ST")
	//if strings.TrimSpace(strArr1[0]) == "" {
	//	fmt.Println(strings.TrimLeft(strArr1[1],","))
	//}
	//
	//str2 := "601360ST,601361ST,601313ST"
	//strArr2 := strings.Split(str2, "601313ST")
	//if strings.TrimSpace(strArr2[1]) == "" {
	//	fmt.Println(strings.TrimRight(strArr2[0],","))
	//}
	//
	//
	//str3 := "601360ST,601361ST,601313ST,601362ST"
	//strArr3 := strings.Split(str3, "601313ST")
	//fmt.Println(strings.TrimLeft(strArr3[1],",") + "," + strings.TrimRight(strArr3[0],","))

	//encryption.GenRsakeyDemo()
	//encryption.RsaEnDecrypt()
	//encryption.DesDemo()
	encryption.AesEnDeCryptDemo()
}



