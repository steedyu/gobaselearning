package ivsample

import "fmt"

func getValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	//.\string.go:9: cannot use nil as type string in return argument
	//return nil, false
	return "不存在数据", false
}

func StringNil()  {
	intmap:=map[int]string{
		1:"a",
		2:"bb",
		3:"ccc",
	}

	v,err:=getValue(intmap,3)
	fmt.Println(v,err)
}