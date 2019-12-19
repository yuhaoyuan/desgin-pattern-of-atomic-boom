package singleton

import "sync"

// 导弹发射器只有一个
type BoomControl struct {

}

var singleton *BoomControl

var Once sync.Once

func GetSingleton() *BoomControl {
	Once.Do(func (){
		singleton = &BoomControl{}
	})
	return singleton
}
