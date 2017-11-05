package ivsample

import "fmt"

func SwitchJudgeType() {
	i := GetValue()

	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case interface{}:
		fmt.Println("interface")
	default:
		fmt.Println("unknown")
	}
}

/*
编译失败，因为type只能使用在interface
当返回interface的时候是可以去用上述方式去判断的
.\interface.go:100: cannot type switch on non-interface value i (type int)
 */
//func GetValue() int {
//	return 1
//}

func GetValue() interface{} {
	return 1
}

func convertInterfaceToString(a interface{}) string {
	/*
	.\type.go:32: cannot convert a (type interface {}) to type string: need type assertion
	 */
	//return string(a)

	v,ok := a.(string)
	if ok {
		return v
	}else {
		return "err"
	}
}

func ConvertInterfaceToStringDemo() {

	fmt.Println(convertInterfaceToString(1))
	fmt.Println(convertInterfaceToString("my name is jerome"))
}


