[TOC]



# IO

>  io 包对IO原语进行基本的抽象接口，包括文件、网络、管道、套接字等。


## 变量

### 1. EOF
> EOF 是一个特殊的变量，表示一个没有更多数据的EOF。
> 在 Go 语言中，当读取或写入一个空的输入输出数据流时，会返回 io.EOF（End Of File）错误。io.EOF 是 io 包中的一个特殊错误值，表示输入输出操作已经到达了数据流的末尾或无法访问的文件位置。 当读取输入流（如从文件、网络连接或标准输入）时，如果到达了数据流的末尾，再进行读取操作时会立即返回 io.EOF 错误。同样，当写入输出流时，如果到达了数据流的末尾，再进行写入操作时也会立即返回 io.EOF 错误。

```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	useEOF()
}

// io.EOF: 表示读取到达结束就会返回该错误
func useEOF() {
	data := bytes.NewBufferString("hello word")
	for i := 0; i < 3; i++ {
		dst := make([]byte, 11)
		l, err := data.Read(dst)
		fmt.Printf("第%d次读取结果:data=%s,len=%d,err=%v,\n\n", i+1, dst, l, err)
	}

}

```



### 2. ErrClosedPipe
> 当从一个已关闭的Pipe读取或者写入时，会返回ErrClosedPipe。


3. ErrNoProgress 
> 某些使用io.Reader接口的客户端如果多次调用Read都不返回数据也不返回错误时，就会返回本错误，一般来说是io.Reader的实现有问题的标志。


4. ErrShortBuffer 
5. ErrShortWrite 
6. ErrUnexpectedEOF 

## 类型

### 1.  Reader

> 定义：Read 将 len(p) 个字节读取到 p 中。它返回读取的字节数 n（0 <= n <= len(p)） 以及任何遇到的错误。
>
> 1. 即使 `Reader` 返回的 `n  < len(p)` 本方法在被调用时仍时仍使用 `p` 的全部长度为暂存空间。
> 2. 如果有部分数据不够, `Reader` 按惯例会返回可以读取到的数据，而不是等待更多数据。如100个字节的 `Read` 调用返回了50个字节，则会返回 `50, nil` 错误。
> 3. 调用 `Read` 后，如果没有更多的数据，`err` 就会返回 `io.EOF。 

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```



#### 实现类型

- IO 包本身的接口
  1. `LimitedReader`
  2. `PipeReader`
  3. `SectionReader`
- 其他标准库类型
  1. `os.File` 类型：变量有`os.Stdin`, `os.Stdout`, `os.Stderr` ，功能主要是读取文件内容到 p 上。
  2. `strings.Reader`：字符串的读取器，方法主要功能是读取该结构体的内容到 p 上。
  3. `bytes.Buffer` ：
  4. `bufio.Reader`
  5. `bytes.Reader`
  6. `compress/gzip`
  7. `crypto/cipher.StreamReader`
  8. `crypto/tls.Conn`
  9. `encoding/csv.Reader`
  10. `mime/multipart.Part`
  11. `net/conn`



#### 代码示例

- 能正常读取到内容时仍使用 `p` 的全部长度为暂存空间（代码参见：reader.go）

  ```go
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
  ```

  

- 如4个字节的内容， `Read` 调用返回了2个字节。（代码参见：reader.go）

  ```go
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
  ```

  

- 再次读取返回错误。（代码参见：reader.go）

  ```go
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
  ```

  

### 2. Writer

> 定义：`Write` 将 len(p) 个字节从 p 中写入到基本数据流中，它返回从 p 中被写入的字节数 n（0 <= n <= len(p)）以及任何遇到的引起写入提前停止的错误。
>
> 1. `Write` 返回的 n < len(p)，它就必须返回一个 非nil 的错误。

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```



#### 实现类型

1. `PipeWriter` : IO 接口
2. `LimitedReader`: IO 接口
3. `SectionReader`: IOjie'k
4. `os.File`
5. `bufio.Writer`
6. `bytes.Buffer`
7. `compress/gzip.Writer `
8. `crypto/tls.Conn`
9. `encoding/csv.Writer`
10. `net/conn `



#### 代码示例

- 写入成功,  n = len(p)

  ```go
  func writeSuccess() {
  	src := "hello word!\n"
  	p := []byte(src)
  	// 由于 os.Stdout 是一个 *file 类型, 所以实现了write方法,也就是将内容写入到标准错误中
  	len1, err1 := os.Stdout.Write(p)
  	fmt.Println("wirte success, len=%d,err=%v", len1, err1)
  }
  ```

  

- 写入失败，n < len(p)

  ```go
  // TODO  这个例子是错误的,因为buffer是无限制容器的,所以不会报错
  func writeFail() {
  	var buf = bytes.Buffer{}
  	data := make([]byte, 1024*1024*10)
  	len1, err1 := buf.Write(data)
  	fmt.Println("wirte success, len=%d,err=%v", len1, err1)
  }
  
  ```

  

### 3. Closer

> 该接口比较简单，只有一个关闭的方法，用于关闭数据流。

````go
type Closer interface {
    Close() error
}
````



