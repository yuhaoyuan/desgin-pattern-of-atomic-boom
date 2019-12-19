package factory

import "testing"

func sendMail(factory AtomicBoomFactory, n string, w int, ts int64) int { // 工厂生产函数
	boom := factory.Create()
	boom.SetName(n)
	boom.SetWeight(w)
	boom.SetTime(ts)
	return boom.DfMail()
}

func TestMail(t *testing.T) {
	var (
		factory AtomicBoomFactory
	)

	factory = AtomicBoom17FactoryBase{}
	if sendMail(factory, "b1", 100, 12345) != 17 {
		t.Fatal("error with factory method pattern")
	}

	factory = AtomicBoom21FactoryBase{}
	if sendMail(factory, "b2", 100, 678910) != 21 {
		t.Fatal("error with factory method pattern")
	}
}
