package prettyany

import (
	"reflect"
	"testing"
)

type A struct {
	Int int
	Map *map[string]int
	Interface interface{}
	Bytes []byte
	Int8s []int8
	Uint8s []uint8
	Runes []rune
}

func TestPrettyAny_Pretty(t *testing.T) {
	a := A{
		Int: 1,
		Map: &map[string]int {
			"123": 123,
			"876": 876,
		},
		Interface: func()error{
			return nil
		},
		Bytes: []byte("123"),
		Int8s: []int8{1,2,3},
		Uint8s: []uint8{'1','2','3'},
		Runes: []rune("你好"),
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
