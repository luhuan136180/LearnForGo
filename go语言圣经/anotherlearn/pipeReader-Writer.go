package main

import (
	"errors"
	"fmt"
	"io"
	"time"
)

func main() {
	pipeReader, pipWriter := io.Pipe()
	go PipWriter(pipWriter)
	go PipRead(pipeReader)
	time.Sleep(30 * time.Second)

}

func PipWriter(writer *io.PipeWriter) {
	data := []byte("Go语言中文网")

	for i := 0; i < 3; i++ {
		n, err := writer.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("写入的字节 %d \n", n)
	}
	writer.CloseWithError(errors.New("写入段已经关闭"))
}

func PipRead(reader *io.PipeReader) {

	buf := make([]byte, 128)

	for {
		fmt.Println("接口端开始阻塞5秒钟...")
		time.Sleep(5 * time.Second)
		fmt.Println("接收端开始接受")
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("收到字节: %d\n buf内容: %s\n", n, buf)
	}

}
