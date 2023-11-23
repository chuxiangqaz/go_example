package main

import (
	"fmt"
	"os"
)

func main() {
	//returnContent()
	//returnEOF()
	returnSomeLen()
}

// 验证读取到真正的内容
func returnContent() {
	file, err := os.OpenFile("./1.txt", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	p := make([]byte, 20)

	// 由于file实现了 io.Reader(p)
	// 读取文件的20个字节到p中
	n, err := file.Read(p)
	if err != nil {
		panic(err)
	}

	sp := string(p)
	fmt.Printf("read file len, len=%d,data=%s, data len = %d, byte data len=%d\n", n, sp, len(sp), len(p))
}

// 验证返回读取到EOF
func returnEOF() {
	file, err := os.OpenFile("./1.txt", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	p := make([]byte, 20)

	// 由于file实现了 io.Reader(p)
	// 读取文件的20个字节到p中
	n, err := file.Read(p)
	if err != nil {
		panic(err)
	}

	sp := string(p)
	fmt.Printf("first read file len, len=%d,data=%s, data len = %d, byte data len=%d\n", n, sp, len(sp), len(p))

	n2, err2 := file.Read(p)
	fmt.Printf("second read file , len=%d, err=%v\n", n2, err2)
}

// 验证读取部分
func returnSomeLen() {
	file, err := os.OpenFile("./1.txt", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	p := make([]byte, 2)

	// 由于file实现了 io.Reader(p)
	// 读取文件的2个字节到p中
	n, err := file.Read(p)
	if err != nil {
		panic(err)
	}

	sp := string(p)
	fmt.Printf("read file len, len=%d,data=%s, data len = %d, byte data len=%d\n", n, sp, len(sp), len(p))
}
