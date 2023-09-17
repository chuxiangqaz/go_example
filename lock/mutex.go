package main

import (
	"fmt"
	"sync"
	"time"
)

// mutex 互斥锁
// 作用：保证同一时间段只有一个 goroutine 持有锁,这就保护了在同一时间段内有且仅有一个 goroutine 访问共享资源, 其他申请锁的 goroutine  将会被堵塞直到锁被释放,可以实现串行化
// 方法：
//  1. lock 锁
//  2. unlock 解锁
//  3. TryLock 以非堵塞摸索取锁, true 代表加锁成功, false 代表加锁失败
// 应用场景：
//	1. 通过互斥锁来实现幂等性，如添加外部联系人成功后我们进行一些操作。
//  2. 比如说要对一个文件只需要一个协程去操作的时候。
//  3. 核心就是要对一个业务只需要一个线程去处理。

func main() {
	// 并发进行操作随机抢到锁
	//TwoGoroutineStrongLock()

	// 启动 100个协程对变量进行累加
	//addLock()
	//add()

	tryLock()
}

func tryLock() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go func() {
			// tryLock 以非堵塞摸索取锁, true 代表加锁成功
			if !lock.TryLock() {
				fmt.Println("未抢到锁")
				return
			}
			defer lock.Unlock()
			fmt.Println("添加外部联系人成功")
			time.Sleep(1 * time.Second)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("tryLock end ...")
}

// 启动 100个协程对变量进行累加
// 当不加锁的时候时候就会出现有可能大不到100的情况, 这是因为 count++ 不是一个原子操作,对应好几个汇编语句。就会导致CPU错乱运行
// 可以对比下 add
func addLock() {
	count := 0
	lock := sync.Mutex{}
	for i := 0; i < 100; i++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			count++
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Printf("addLock end count=%d\n", count)
}

func add() {
	count := 0
	for i := 0; i < 100; i++ {
		go func() {
			count++
		}()
	}

	fmt.Printf("add end count=%d\n", count)
}

// 并发进行操作随机抢到锁
func TwoGoroutineStrongLock() {
	lock := sync.Mutex{}

	go func() {
		for i := 0; i < 10; i++ {
			doIt(&lock, "one", i)
		}

		fmt.Println("fu 1 end..")
	}()

	go func() {
		for i := 0; i < 10; i++ {
			doIt(&lock, "two", i)
		}

		fmt.Println("fu 2 end..")
	}()

	time.Sleep(20 * time.Second)
	fmt.Println("end.............")
}

func doIt(lock *sync.Mutex, fnName string, index int) {
	lock.Lock()
	defer func() {
		fmt.Println("fu " + fnName + " unlock")
		lock.Unlock()
	}()
	fmt.Println("hello fu ["+fnName+"] get lock , and index is", index)
	time.Sleep(1 * time.Second)
}
