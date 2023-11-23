package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello word"
	reader := strings.NewReader(s)
	p := make([]byte, 5)
	n, err := reader.Read(p)
	fmt.Println(n, err, p)
}
