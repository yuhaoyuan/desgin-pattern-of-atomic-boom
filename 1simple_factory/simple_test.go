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

}
