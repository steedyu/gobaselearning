package apipe

import (
	"fmt"
	"os/exec"
	"bytes"
	"io"
	"bufio"
)

func Demo1(useBufferIo bool){

	fmt.Println("Run command `echo -n \"My first command from golang.\"`: ")
	cmd0 := exec.Command("echo", "-n", "My first command from golang.")
	stdout0, err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: Can not obtain the stdout pipe for command No.0: %s\n", err)
		return
	}
	if err := cmd0.Start();err != nil {
		fmt.Printf("Error: The command No.0 can not be startup: %s\n", err)
		return
	}
	if !useBufferIo {
		var outputBuf0 bytes.Buffer
		for {
			tempOutput := make([]byte, 5)
			n, err := stdout0.Read(tempOutput)
			if err != nil {
				if err == io.EOF {
					break
				}else {
					fmt.Printf("Error: Can not read data from the pipe: %s\n", err)
					return
				}
			}
			if n > 0 {
				outputBuf0.Write(tempOutput[:n])
			}
		}
		fmt.Printf("%s\n", outputBuf0.String())
	}else {
		outputbuf0 := bufio.NewReader(stdout0)
		output0,_,err := outputbuf0.ReadLine()
		if err != nil {
			fmt.Printf("Error: Can not read data from the pipe: %s\n", err)
		}
		fmt.Printf("%s\n", string(output0))
	}
}

func Demo2(){
	fmt.Println("Run command `ps aux | grep apipe`: ")
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "apipe")

	stdout1, err := cmd1.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: Can not obtain the stdout pipe for command: %s", err)
		return
	}
	if err := cmd1.Start(); err != nil {
		fmt.Printf("Error: The command can not running: %s\n", err)
		return
	}
	outputBuf1 := bufio.NewReader(stdout1)

	stdin2, err := cmd2.StdinPipe()
	if err != nil {
		fmt.Printf("Error: Can not obtain the stdin pipe for command: %s\n", err)
		return
	}
	//把与cmd1链接的输出管道中的数据全部写入到这个输入管道中
	outputBuf1.WriteTo(stdin2)

	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2
	//建立了对应的输出管道就不能使用Run方法来启动该命令，而应该使用Start方法
	if err := cmd2.Start(); err != nil {
		fmt.Printf("Error: The command can not be startup: %s\n", err)
		return
	}
	//有些命令会等到输入管道被关闭之后才结束运行，所以，再这种情况下，我们就需要在数据被读取之后尽早地手动关闭输入管道
	err = stdin2.Close()
	if err != nil {
		fmt.Printf("Error: Can not close the stdio pipe: %s\n", err)
		return
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Printf("Error: Can not wait for the command: %s\n", err)
		return
	}
	fmt.Printf("%s\n", outputBuf2.Bytes())
}
