package singleton

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	ins1 := GetSingleton()
	ins2 := GetSingleton()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	} else {
		fmt.Println("the same!")
	}
}
