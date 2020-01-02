package x_GoSync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
* sync.Cond
* 条件变量的作用并不是保证在同一时刻仅有一个线程访问某一个共享数据，而是在某一个条件发生时，通知阻塞在该条件上的goroutine(线程)
*
* 条件变量+互斥量(锁)
* a. 互斥量为共享数据的访问提供互斥支持
* b. 条件变量就数据状态的变化向相关线程发出通知(goroutine)
*
* sync.Cond提供的三个相关方法:
* 1. wait: 阻塞当前线程(goroutine)，直到收到该条件变量发来的通知
* 2. signal: 单发通知，让该条件变量向至少一个正在等待它的goroutine发送通知，表示共享数据的状态已经改变
* 3. broadcast: 广播通知，让条件变量给正在等待它的所有goroutine发送通知，告知共享数据的状态已经改变
*
* sync.Cond struct
* // Each Cond has an associated Locker L (often a *Mutex or *RWMutex),
* // which must be held when changing the condition and
* // when calling the Wait method.
* //
* // A Cond must not be copied after first use.
* type Cond struct {
*   noCopy noCopy
*
*   // L is held while observing or changing the condition
*   L Locker
*
*   notify  notifyList
*   checker copyChecker
* }
*
* 通过sync.Cond的定义，我们需要注意以下几点:
* 1. Cond内部存在一个Locker(Mutex或RWMutex)，在Cond状态条件改变或调用Wait方法时，必须被锁住。即Locker是对
*    Wait, Signal，Broadcast进行保护，确保在发送信号的时候不会有新的goroutine进入wait而阻塞。
* 2. Cond变量在第一次创建之后不应该被copy。
* 3. 在调用Signal，Broadcast函数之前，应该确保目标进入wait阻塞状态。
*
 */

func TestCond(t *testing.T) {
	var wg sync.WaitGroup
	cd := sync.NewCond(new(sync.Mutex))
	//cd2 := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cd.L.Lock()
			fmt.Println("goroutine_", i, " start in t:", int(time.Now().Unix()))
			cd.Wait() // 阻塞住等通知
			cd.L.Unlock()
			fmt.Println("goroutine_", i, " end in t:", int(time.Now().Unix()))
		}(i)
	}
	//sw:=0
	sw := 1
	if sw == 0 {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			cd.Signal() // 唤醒一个goroutine
		}
	} else {
		time.Sleep(time.Second)
		cd.Broadcast()
	}
	wg.Wait()
	fmt.Println("end test")
}
