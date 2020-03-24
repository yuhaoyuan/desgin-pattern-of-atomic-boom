package channel

import (
	"fmt"
	"testing"
	"time"
)

func doRequest(xChan *chan int){
	time.Sleep(time.Second)
	*xChan <- 10
	time.Sleep(time.Second * time.Duration(3))
	*xChan <- 20
}

func TestM(t *testing.T){
	xChan := make(chan int, 10)
	timer := time.NewTicker(time.Duration(1) * time.Second)

	go doRequest(&xChan)
	for {
		select {
		case <-timer.C:
			fmt.Println(time.Now())
		case x :=<-xChan:
			fmt.Println("xChan out = ", x)
		}
	}

}