package filedealsample

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"io/ioutil"
)

var (
	FileName string = "D:\\GoPath\\src\\jerome.com\\gobaselearning\\internal\\gousualsample\\filedealsample\\我有一只小毛驴.txt"    //这是我们需要打开的文件，当然你也可以把它定义到从某个配置文件来获取变量。
	InputFile  *os.File    //变量 InputFile 是 *os.File 类型的。该类型是一个结构，表示一个打开文件的描述符（文件句柄）。
	InputError error    //我们使用 os 包里的 Open 函数来打开一个文件。如果文件不存在或者程序没有足够的权限打开这个文件，Open函数会返回一个错误，InputError变量就是用来接收这个错误的。
	Count int            //这个变量是我们用来统计行号的，默认值为0.
)

//顺序读取文件内容  按行读
func FileSequenceRead() {
	//InputFile,InputError = os.OpenFile(FileName,os.O_CREATE|os.O_RDWR,0644) //打开FileName文件，如果不存在就创建新文件，打开的权限是可读可写，权限是644。这种打开方式相对下面的打开方式权限会更大一些。
	InputFile, InputError = os.Open(FileName) //使用 os 包里的 Open 函数来打开一个文件。该函数的参数是文件名，类型为 string 。我们以只读模式打开"FileName"文件。
	if InputError != nil {
		//如果打开文件出错，那么我们可以给用户一些提示，然后在推出函数。
		fmt.Print("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer InputFile.Close()        //defer关键字是用在程序即将结束时执行的代码,确保在程序退出前关闭该文件。
	inputReader := bufio.NewReader(InputFile) //我们使用 bufio.NewReader()函数来获得一个读取器变量（读取器）。我们可以很方便的操作相对高层的 string 对象，而避免了去操作比较底层的字节。
	for {
		Count += 1
		inputString, readerError := inputReader.ReadString('\n')  //我们将inputReader里面的字符串按行进行读取。
		if readerError == io.EOF {
			return  //如果遇到错误就终止循环。
		}
		fmt.Printf("The %d line is: %s", Count, inputString)    //将文件的内容逐行（行结束符'\n'）读取出来。
	}
}

//按列读取数据
func FileColumnRead() {
	filename := "D:\\GoPath\\src\\jerome.com\\gobaselearning\\internal\\gousualsample\\filedealsample\\a.txt"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var Column1, Column2, Column3 []string    //定义3个切片，每个切片用来保存不同列的数据。

	for {
		var FirstRowColumn, SecondRowColumn, ThirdRowColumn string
		_, err := fmt.Fscanln(file, &FirstRowColumn, &SecondRowColumn, &ThirdRowColumn)    //如果数据是按列排列并用空格分隔的，我们可以使用 fmt 包提供的以 FScan 开头的一系列函数来读取他们。
		if err != nil {
			break
		}
		Column1 = append(Column1, FirstRowColumn)  //将第一列的每一行的参数追加到空切片Column1中。以下代码类似。
		Column2 = append(Column2, SecondRowColumn)
		Column3 = append(Column3, ThirdRowColumn)
	}
	fmt.Println(Column1)
	fmt.Println(Column2)
	fmt.Println(Column3)
}


//带缓冲的读取
func FileBufferRead() {

	f, err := os.Open(FileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	ReadSize := make([]byte, 1024) //指定每次读取的大小为1024。
	ReadByte := make([]byte, 4096, 4096) //指定读取到的字节数。
	r := bufio.NewReader(f)
	for {
		ActualSize, err := r.Read(ReadSize)    //回返回每次读取到的实际字节大小。
		if err != nil && err != io.EOF {
			panic(err)
		}
		if ActualSize == 0 {
			break
		}
		ReadByte = append(ReadByte, ReadSize[:ActualSize]...)    //将每次的读取到的内容都追加到我们定义的切片中。
	}
	fmt.Println(string(ReadByte))    //打印我们读取到的内容，注意，不能直接读取，因为我们的切片的类型是字节，需要转换成字符串这样我们读取起来会更方便。

}


//将整个文件的内容读到一个字节切片中
func FileTotalRead() {
	buf, err := ioutil.ReadFile(FileName) //将整个文件的内容读到一个字节切片中。
	fmt.Println(reflect.TypeOf(buf))
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
}





