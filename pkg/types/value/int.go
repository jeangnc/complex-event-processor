package value

import "fmt"

type IntValue struct {
	value int64
}

func (v IntValue) ToInt() int64 {
	return v.value
}

func (v IntValue) ToFloat() float64 {
	return float64(v.value)
}

func (v IntValue) ToString() string {
	return fmt.Sprintf("%v", v.value)
}
