package Reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	name   string
	age    int
	unknow interface{}
}

func (s *Person) GrowUp() {
	s.age += 1
}

type Animal struct {
	Attack int
	Blood  int
	Person
}

func (s *Animal) GrowUp() {
	s.Attack += 1
	s.Blood += 1
}

// 基础变量的反射
func TestTypeAndValue(t *testing.T) {
	// TypeOf会返回目标数据的类型，比如int/float/struct/指针等
	// valueOf返回目标数据的的值

	var x int = 1
	var y string = "啦啦啦啦啦啦啦啦"

	fmt.Println("type: ", reflect.TypeOf(x))
	fmt.Println("value: ", reflect.ValueOf(x))
	fmt.Println("type: ", reflect.TypeOf(y))
	fmt.Println("value: ", reflect.ValueOf(y))

	var s Person
	s.name = "XDE"
	s.age = 99
	s.unknow = "cat"
	fmt.Println("type: ", reflect.TypeOf(s))
	fmt.Println("value: ", reflect.ValueOf(s))
	fmt.Println("type: ", reflect.TypeOf(s.unknow))
	fmt.Println("value: ", reflect.ValueOf(s.unknow))
}

// struct 的反射
func TestStruct(t *testing.T) {
	x := Person{
		"55k",
		30,
		"给阿姨倒一杯卡布奇诺",
	}
	xType := reflect.TypeOf(x)
	for i := 0; i < xType.NumField(); i++ {

		key := xType.Field(i)
		//fmt.Println(key)
		//fmt.Println("type: ", reflect.TypeOf(key))
		//fmt.Println("value: ", reflect.ValueOf(key))

		// 通过interface方法来获取key所对应的值
		value := reflect.ValueOf(i).Interface()

		fmt.Printf("第%d个字段是：%s:%v = %v \n", i+1, key.Name, key.Type, value)

	}
}

// 匿名、内嵌的反射会是怎样呢
func TestAnonymous(t *testing.T) {
	zombie := Animal{
		Attack: 5,
		Blood: 0,
	}
	zombie.age=30
	zombie.name="jack"
	zombie.unknow="he was a man"

	fmt.Println("type: ", reflect.TypeOf(zombie.Person))
	fmt.Println("value: ", reflect.ValueOf(zombie.Person))

	zT := reflect.TypeOf(zombie)
	fmt.Printf("%#v\n", zT.Field(2))
	// 获取匿名字段的值的详情
	zV := reflect.ValueOf(zombie)
	fmt.Printf("%#v\n", zV.Field(2))

	// 顺便好奇一下group up 函数会不会重定义
	fmt.Println(zombie)
	zombie.GrowUp()
	fmt.Println(zombie)
	zombie.Person.GrowUp()
	fmt.Println(zombie)
}

// 模拟一下可能存在的实际业务场景
func TestKindAndSeT(t *testing.T) {
	XiaoEn := Animal{
		1,
		1,
		Person{
			age: 28,
			name: "Shaun",
			unknow: "Londoner",
		},
	}
	if reflect.TypeOf(XiaoEn).Kind() == reflect.Struct{
		fmt.Println(XiaoEn)

		// 反射只能通过指针修改值
		ptrXE := reflect.ValueOf(&XiaoEn)
		// 修改值必须是指针类型否则不可行
		if ptrXE.Kind() != reflect.Ptr{
			fmt.Println("不是指针类型，没法进行修改操作")
			return
		}

		// 获取指针所指向的元素
		ptrXE = ptrXE.Elem()

		// 获取目标key的Value的封装
		ptrXEA := ptrXE.FieldByName("Attack")

		if ptrXEA.Kind() == reflect.Int{
			ptrXEA.SetInt(10)
		}

		fmt.Printf("%#v \n", XiaoEn)

	}

}
