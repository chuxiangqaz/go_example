package main

import (
	"bytes"
	"fmt"
)

func main() {
	data := []byte("hello world")
	reader := bytes.NewReader(data)
	dst := make([]byte, 5)
	n, err := reader.Read(dst)
	fmt.Println(n, err, string(dst))
}
