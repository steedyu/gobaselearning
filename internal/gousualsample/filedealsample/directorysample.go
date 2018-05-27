package filedealsample

import (
	"os"
	"fmt"
)

func FileDelete() {

	deleteFile := "D:\\GoPath\\src\\jerome.com\\gobaselearning\\internal\\gousualsample\\filedealsample\\复制的小毛驴.txt"
	os.Remove(deleteFile)
}

func DirCreateDelete() {

	OneLevelDirectory := "yinzhengjie"
	MultilevelDirectory := "yinzhengjie/golang/code"

	os.Mkdir(OneLevelDirectory, 0777)    //创建名称为OneLevelDirectory的目录，设置权限为0777。相当于Linux系统中的“mkdir yinzhengjie”
	os.MkdirAll(MultilevelDirectory, 0777)    //创建MultilevelDirectory多级子目录，设置权限为0777。相当于Linux中的 “mkdir -p yinzhengjie/golang/code”

	err := os.Remove(MultilevelDirectory) //删除名称为OneLevelDirectory的目录，当目录下有文件或者其他目录是会出错。
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll(OneLevelDirectory) //根据path删除多级子目录，如果path是单个名称，那么该目录不删除。
}