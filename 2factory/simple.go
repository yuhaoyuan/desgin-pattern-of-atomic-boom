package factory

import "fmt"

// 原子弹接口
type AtomicBoom interface {
	SetName(name string)
	SetWeight(weight int)
	SetTime(boomTs int64) string
	DfMail() int
}

// BoomBase 是原子弹的基类
type BoomBase struct {
	name string  // private
	weight int
	boomTs int64
}
func (t *BoomBase) SetName(name string) {
	t.name = name
}
func (t *BoomBase) SetWeight(weight int) {
	t.weight = weight
}
func (t *BoomBase) SetTime(boomTs int64) string {
	t.boomTs = boomTs
	return fmt.Sprintf("the %s  will boom in the %d", t.name, boomTs)
}

// DF17 "继承"（匿名函数的方式） 基类中所有的成员函数
type DF17 struct {
	*BoomBase
}
func (t *DF17) DfMail() int {
	fmt.Println("send DF-17 Mail" )
	return 17
}

// Df21
type DF21 struct {
	*BoomBase
}
func (t *DF21) DfMail() int {
	fmt.Println("send DF-21 Mail" )
	return 21
}

// 原子弹工厂接口
type AtomicBoomFactory interface {
	Create() AtomicBoom
}

// 17号导弹工厂类
type AtomicBoom17FactoryBase struct {

}
func (AtomicBoom17FactoryBase) Create() AtomicBoom {
	return &DF17{
		BoomBase: &BoomBase{},
	}
}


// 21号导弹工厂类
type AtomicBoom21FactoryBase struct {

}
func (AtomicBoom21FactoryBase) Create() AtomicBoom {
	return &DF21{
		BoomBase: &BoomBase{},
	}
}


