package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	//writeSuccess()
	writeFail()
}

func writeSuccess() {
	src := "hello word!\n"
	p := []byte(src)
	// 由于 os.Stdout 是一个 *file 类型, 所以实现了write方法,也就是将内容写入到标准错误中
	len1, err1 := os.Stdout.Write(p)
	fmt.Println("wirte success, len=%d,err=%v", len1, err1)
}

// TODO  这个例子是错误的,因为buffer是无限制容器的,所以不会报错
func writeFail() {
	var buf = bytes.Buffer{}
	data := make([]byte, 1024*1024*10)
	len1, err1 := buf.Write(data)
	fmt.Println("wirte success, len=%d,err=%v", len1, err1)
}
