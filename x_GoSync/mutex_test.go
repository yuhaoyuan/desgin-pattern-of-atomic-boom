package x_GoSync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func LockWithPrint(name string, locker sync.Locker) {
	fmt.Println(name, "ready lock")
	locker.Lock()
	fmt.Println(name, " locked")
}

func TestMutex(t *testing.T) {
	var m0 sync.Mutex
	wg := sync.WaitGroup{}
	LockWithPrint("m0", &m0)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lockName := fmt.Sprintf("gr_%d", i)
			LockWithPrint(lockName, &m0)
			m0.Unlock()
			fmt.Println(lockName, "unlocked")
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("只要我按一下这个开关，所有的g都会开始获锁")

	m0.Unlock()
	fmt.Println("m0 unlocked!")

	wg.Wait()
	fmt.Println("test done!")
}


func RwRead(i int, m *sync.RWMutex, wg *sync.WaitGroup){
	fmt.Println("i= ", i, "read-lock-ready")
	m.RLock()
	fmt.Println("i= ", i, "read-locked!")

	step := 0
	for{
		fmt.Println("i= ", i, "read-lock - ", "read someth step-", step)
		step++
		if step == 10{
			break
		}
	}
	m.RUnlock()
	wg.Done()
}

func RwWrite(i int, m *sync.RWMutex, wg *sync.WaitGroup){
	fmt.Println("i= ", i, " write-lock-ready")
	m.Lock()
	fmt.Println("i= ", i, "write-locked!")

	step := 0
	for{
		fmt.Println("i= ", i, "lock - ", "write someth step-", step)
		step++
		if step == 10{
			break
		}
	}
	m.Unlock()
	wg.Done()
}

func TestRwMutex(t *testing.T) {
	var rwM sync.RWMutex

	wg := sync.WaitGroup{}
	wg.Add(14)

	go RwRead(1, &rwM, &wg)
	go RwWrite(2, &rwM, &wg)
	go RwRead(3, &rwM, &wg)
	go RwWrite(4, &rwM, &wg)
	go RwRead(50, &rwM, &wg)
	go RwRead(51, &rwM, &wg)
	go RwRead(52, &rwM, &wg)
	go RwRead(53, &rwM, &wg)
	go RwRead(54, &rwM, &wg)
	go RwWrite(6, &rwM, &wg)
	go RwRead(7, &rwM, &wg)
	go RwWrite(8, &rwM, &wg)
	go RwRead(9, &rwM, &wg)
	go RwWrite(10, &rwM, &wg)

	wg.Wait()
}
