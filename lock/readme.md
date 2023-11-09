# mutex

## 是什么
> `sync.Mutex` 本质是一个互斥锁, 它可以保证同一时刻只有一个 goroutine 访问某个资源,是一个互斥锁。说的通俗些就是保证同一时刻只有一个 goroutine 访问部分代码片段。

## 使用场景
以下案例的前提都是单机服务，若要在分布式环境下使用，需要使用分布式锁。(核心就是要对一个业务只需要一个线程去处理。)
1. 当我们设计一个微信支付回调接口的时, 就需要使用 `sync.Mutex` 来保证幂等性，防止重复回调。
2. 当我们设计一个秒杀系统时, 就需要使用 `sync.Mutex` 来保证同一时刻只有一个 goroutine 访问某个资源。
3. 当我们设计一个分布式锁时, 就需要使用 `sync.Mutex` 来保证同一时刻只有一个 goroutine 访问某个资源。
4. 当我们设计一个缓存系统时, 就需要使用 `sync.Mutex` 来保证同一时刻只有一个 goroutine 访问某个资源。
5. 比如说要对一个文件只需要一个协程去操作的时候。
   

## 函数列表
sync.Mutex 的是一个结构体, 包含以下方法:
1. `Lock()` 加锁,如果当前线程已经加锁, 则阻塞当前线程, 直到当前线程解锁。
2. `Unlock()` 解锁,如果当前线程没有加锁, 则会 panic。
3. `TryLock()` 尝试加锁,若当前线程已经加锁, 则返回 false, 否则返回 true。

## 如何使用
1. 创建一个 `sync.Mutex` 类型的变量，并将其赋值给一个变量。
2. 在需要加锁的代码块中，调用 `mutex.Lock()` 方法加锁。
3. 在需要解锁的代码块中，调用 `mutex.Unlock()` 方法解锁, 一般使用 defer 关键字。



```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    addLock()
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
```


## 注意事项
1. 对于一个没有加锁的 lock 进行解锁, 会 panic。
```go
package main

import (
	"sync"
	"time"
)

func main() {
	lock := sync.Mutex{}
	lock.Unlock()

	time.Sleep(10 * time.Second)
}

//go run no_unlock.go
//fatal error: sync: unlock of unlocked mutex
//
//goroutine 1 [running]:

```

2. 加锁和解锁必须在同一个 goroutine 中，否则会引发死锁或竞态条件。 
3. 加锁和解锁必须在同一个代码块中，否则会出现资源保护问题。 
4. 如果加锁和解锁不在同一个代码块中，需要使用带有 defer 语句的锁定，以确保在代码块的所有路径上都会释放锁，避免资源泄露。 
5. 加锁和解锁必须在同一个 goroutine 中，否则会引发死锁或竞态条件。


## 原理

## 设计学习

---

# rwmutex

## 是什么
> `sync.RWMutex` 是读写锁, 可以保证同一时刻只有多个 goroutine 读取某个资源,保证可以并发读，但是不能并发读写。
> 1. 在同一时间只能有一个goroutine 获取到写锁。
> 2. 在同一个时间可以有多个 goroutine 获取到读锁。
> 3. 在同一个时间只能存在写锁或读锁。

## 使用场景(解决什么问题)
1. 共享资源：当多个goroutine需要访问共享资源，如文件、内存中的数据结构或网络连接时，可以使用sync.RWMutex来保护资源的并发访问。通过使用读写锁，可以允许多个goroutine同时读取资源，但仅允许一个goroutine进行写入操作。
2. 全局状态：当全局状态需要在多个goroutine之间共享时，可以使用sync.RWMutex来保护对全局状态的访问。这样可以确保在修改状态时不会发生冲突，并保证状态的完整性。
3. 缓存：在需要维护缓存一致性的场景中，可以使用sync.RWMutex来保护对缓存数据的访问。通过读写锁，可以允许多个goroutine同时读取缓存，但仅允许一个goroutine进行写入操作，以确保缓存的一致性和防止数据竞争。 

## 函数列表
1. lock 获取写锁,
2. RLock 获取读锁,若未获取到(如已经加了写锁)则堵塞
3. TryRLock 尝试给加读锁, 若获取到则返回 true,否则则返回flase
4. TryLock 尝试给加写锁, 若获取到则返回 true,否则则返回flase
5. RUnlock 解读锁,若对一个未加锁的结构直接进行解锁, 则返回错误 fatal error: sync: Unlock of unlocked RWMutex
6. Unlock 解写锁,若对一个未加锁的结构直接进行解锁, 则返回错误 fatal error: sync: Unlock of unlocked RWMutex

## 如何使用
1. 创建一个 `sync.RWMutex` 类型的变量，并将其赋值给一个变量。
2. 在需要加锁的代码块中，调用 `mutex.Lock()` 方法加锁。
3. 在需要解锁的代码块中，调用 `mutex.Unlock()` 方法解锁, 一般使用 defer 关键字。

1.  两个协程并行读取，互相不堵塞
```go
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
```
2. 两个协程并发读写,可以看到两个协程是争抢获取读锁和写锁, 当一个锁执行完后另外一个锁才能执行
```go
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
```
3. 两个协程并发读写, 但是当一个协程获取到读锁的时候, 另外一个协程就不能获取到读锁, 直到第一个协程释放读锁
```go
// 可以看到下面这个函数的两个协程会有如下两个结果
// 1， "i am [two] and write", 在输出 "i am [one] and read", 然后休眠2S结束
// 2. "i am [one] and read", 堵塞1S， 在输出"i am [two] and write", 然后在堵塞1S，输出"readAndWriteNotWait end................."
func readAndWriteNotWait() {
	lock := &sync.RWMutex{}
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

```
## 注意事项

## 原理

## 设计学习





