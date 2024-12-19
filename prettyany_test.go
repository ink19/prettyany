package prettyany

import (
	"reflect"
	"testing"
)

type A struct {
	A int
	B *map[string]int
	C interface{}
	D []byte
	E []int8
	F []uint8
	G []rune
}

func TestPrettyAny_Pretty(t *testing.T) {
	a := A{
		A: 1,
		B: &map[string]int {
			"123": 123,
			"876": 876,
		},
		C: func()error{
			return nil
		},
		D: []byte("123"),
		E: []int8{1,2,3},
		F: []uint8{'1','2','3'},
		G: []rune("你好"),
	}
	t.Log(NewPrettyAny().Pretty(a))
}

func TestReflect(t *testing.T) {
	t1 := reflect.TypeOf([]rune("你好"))
	t2 := reflect.TypeOf([]int32{1,2,3})

	t.Log(t1)
	t.Log(t2)
	t.Log(t1 == t2)
}
