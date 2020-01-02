package x_GoSync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type taksSt struct {
	name string
	id   int
}

func producer(cond *sync.Cond, pool *sync.Pool, wg *sync.WaitGroup) {
	go func() {
		for i := 1; i <= 6; i++ {
			time.Sleep(time.Second) //
			fmt.Println("produce: ", i)
			pool.Put(taksSt{
				name: "p-task",
				id:   i,
			})
			//cond.Broadcast() // 如果不在这里唤醒， 那么会死锁（因为所有goroutine都阻塞
		}
	}()
}

func consumer(cond *sync.Cond, pool *sync.Pool, wg *sync.WaitGroup) {
	go func() {
		for true {
			fmt.Println("consumer - ", )
			//cond.L.Lock()
			vTask := pool.Get()
			for ; vTask == nil; {
				vTask = pool.Get()
				if vTask != nil {
					fmt.Println("for v != nil", vTask)
					wg.Done()
				}
			}
			//cond.Wait() // 阻塞
			//fmt.Println("cond release vTask = ", vTask)
			//cond.L.Unlock()
			//pool.Put(vTask)
			//cond.Broadcast()
			//cond.Signal()
		}
	}()
}

func TestPoolWithCond(t *testing.T) {
	//lock := sync.Mutex{}
	//cond := sync.NewCond(&lock)
	pool := &sync.Pool{} //sync.Pool主要用途是增加临时对象的重用率，减少GC负担
	//wg := &sync.WaitGroup{}

	pool.Put(taksSt{
		"ttt 1",
		1,
	})
	x := pool.Get()
	fmt.Println(x)
	x = taksSt{
		name: "new t",
		id:   2,
	}
	pool.Put(x)

	x = pool.Get()
	fmt.Println(x)
	x = taksSt{
		name: "new t",
		id:   3,
	}
	pool.Put(x)

	x = pool.Get()
	x = "asda"
	fmt.Println(x)
	pool.Put(x)

	//wg.Add(6)
	//producer(cond, pool, wg)
	//for i:=0;i<6;i++ {
	//	consumer(cond, pool, wg)
	//}
	//wg.Wait()
	//time.Sleep(time.Second*10)
}
