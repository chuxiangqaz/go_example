package main

import (
	"fmt"
	"sync"
	"time"
)

// rwmutex 读取锁
// 作用：保证可以并发读，但是不能并发读写,
//  1. 在同一时间只能有一个goroutine 获取到写锁
//  2. 在同一个时间可以有多个 goroutine 获取到读锁
//  3. 在同一个时间只能存在写锁或读锁
//
// 方法
//  1. lock 获取写锁,
//  2. RLock 获取读锁,若为获取到则堵塞
//  3. TryRLock 获取读锁, 若获取到则返回 true,否则则返回flase
//  4. TryLock 获取写锁, 若获取到则返回 true,否则则返回flase
//  5. RUnlock 解读锁,若对一个未加锁的结构直接进行解锁, 则返回错误 fatal error: sync: Unlock of unlocked RWMutex
//  6. Unlock 解写锁,若对一个未加锁的结构直接进行解锁, 则返回错误 fatal error: sync: Unlock of unlocked RWMutex
func main() {
	//readAndRead()
	//readAndWrite()
	readAndWriteNotWait()
}

// 两个协程并行读取，互相不堵塞
func readAndRead() {
	lock := &sync.RWMutex{}
	lock.Unlock()
	go func() {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("i am [one] and  read ....")
		time.Sleep(1 * time.Second)
	}()

	go func() {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("i am [two] and read ....")
		time.Sleep(1 * time.Second)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("readAndRead end.................")
}

func readAndWrite() {
	lock := &sync.RWMutex{}

	go func() {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("i am [one] and  read ....")
		time.Sleep(1 * time.Second)
	}()

	go func() {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("i am [two] and wite ....")
		time.Sleep(1 * time.Second)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("readAndWrite end.................")
}

// 读锁和写锁等待执行
func readAndWriteNotWait() {
	lock := &sync.RWMutex{}
	lock.Unlock()
	go func() {
		if !lock.TryRLock() {
			fmt.Println("i am [one] not get read lock")
			return
		}

		defer lock.RUnlock()
		fmt.Println("i am [one] and  read")
		time.Sleep(1 * time.Second)
	}()

	go func() {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("i am [two] and write")
		time.Sleep(1 * time.Second)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("readAndWrite end.................")
}
