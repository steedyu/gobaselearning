package ivsample

import "fmt"

func GotoLoop() {

	for i := 0; i < 10; i++ {
		//.\gotoloop.go:9: goto loop jumps into block starting at .\gotoloop.go:5
		//loop:
		fmt.Println(i)
	}
	goto loop
	loop:
	fmt.Println("end")
}