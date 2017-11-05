package ivsample

import (
	"fmt"
	"errors"
)

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func tryTheThing() (string,error)  {
	return "",ErrDidNotWork
}

func DoTheThing2(reallyDoIt bool) (err error) {
	var result string
	if reallyDoIt {
		result, err = tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func VariableRange() {
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
	/*
	<nil>
	<nil>
	因为 if 语句块内的 err 变量会遮罩函数作用域内的 err 变量
	 */
	fmt.Println(DoTheThing2(true))
	/*
	did not work
	 */

}
