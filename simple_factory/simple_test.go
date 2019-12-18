package simple_factory

import (
	"fmt"
	"testing"
)

func TestThe1(t *testing.T){  //
	b1 := NewAtomicBoom(1)
	rsAnswer1 := b1.SetTime(12345)
	fmt.Println(rsAnswer1)

	b2 := NewAtomicBoom(1)
	rsAnswer2 := b2.SetTime(12345)
	fmt.Println(rsAnswer2)

	////
	b3 := fatGuy{}
	b3.SetTime(43210)
}
/*
我想要一个原子弹，只需要知道New能创建，Set能创建时间，并在那个时间爆炸即可。
这个例子上还可以加上GetTs、updateTs等等
*/