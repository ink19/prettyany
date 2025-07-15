### Usage

```go
type A struct {
	S int
	B *map[string]int
	T interface{}
}

func print() {
	a := A{
		S: 1,
		B: &map[string]int {
			"123": 123,
			"876": 876,
		},
		T: func()error{
			return nil
		},
	}
	fmt.Printf("%v\n", reflect.ValueOf(a).Field(0))
	fmt.Printf("%s\n", prettyany.NewPrettyAny().Pretty(a))
}

```
