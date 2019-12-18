package simple_factory

import (
	"fmt"
)

//  简单工厂---让我们来造原子弹把
type AtomicBoom interface {
	SetTime(timestamp int64) string
}

type littleBoy struct {  // 注意这里是小写(私有函数)
	boomTs int64
}
func (t *littleBoy) SetTime(timestamp int64) string {
	t.boomTs=timestamp
	return fmt.Sprintf("the little boy will boom in the %d", timestamp)
}

type fatGuy struct {
	boomTs int64
}
func (t *fatGuy) SetTime(timestamp int64) string {
	t.boomTs = timestamp
	return fmt.Sprintf("the fat guy will boom in the %d", timestamp)
}


func NewAtomicBoom(switcH int) AtomicBoom {
	if switcH == 1{
		return &littleBoy{}
	} else if switcH ==2 {
		return &fatGuy{}
	}
	return nil
}

