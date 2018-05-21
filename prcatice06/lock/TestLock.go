package main

import (
	"math/rand"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//互斥锁
var lock sync.Mutex
var rwLock sync.RWMutex

func main()  {
	//testLock()
	testRWLock()
}

func testLock()  {

	a := map[int]int{
		1:1,
		2:2,
		3:3,
		4:4,
		5:5,
	}

	for i := 0 ; i < 2; i++{
		go func(b map[int]int) {
			lock.Lock()
			defer lock.Unlock()
			b[1] = rand.Intn(100)
		}(a)
	}

	lock.Lock()
	defer lock.Unlock()
	fmt.Println(a)
}

//读写锁比互斥锁性能提升了100倍
func testRWLock()  {

	var count int32
	a := map[int]int{
		1:1,
		2:2,
		3:3,
		4:4,
		5:5,
	}

	for i := 0 ; i < 2; i++{
		go func(b map[int]int) {
			//rwLock.Lock()
			lock.Lock()
			b[1] = rand.Intn(100)
			//模拟写 10ms
			time.Sleep(10 * time.Microsecond)
			//rwLock.Unlock()
			lock.Unlock()
		}(a)
	}

	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			//for循环
			for {
				//rwLock.RLock()
				lock.Lock()
				//fmt.Println(a)
				//模拟读 1ms
				time.Sleep(time.Microsecond)
				//rwLock.RUnlock()
				lock.Unlock()
				//原子操作
				atomic.AddInt32(&count, 1)
			}
		}(a)
	}

	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count))
}