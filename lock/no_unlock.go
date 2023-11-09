package main

import (
	"sync"
	"time"
)

func main() {
	lock := sync.Mutex{}
	go func() {
		lock.Lock()
		time.Sleep(time.Second)

	}()

	go func() {
		time.Sleep(3 * time.Second)
		lock.Unlock()
	}()

	time.Sleep(10 * time.Second)
}
