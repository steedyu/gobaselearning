package ivsample

import "fmt"

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

func IotaCase() {
	fmt.Println(x,y,z,k,p)
}


const cl  = 100
var bl    = 123

func takeConstVarAddress()  {
	println(&bl,bl)
	/*
	常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用，
	.\constiota.go:24: cannot take the address of cl
	 */
	//println(&cl,cl)
}