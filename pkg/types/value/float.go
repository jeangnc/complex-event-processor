package value

import "fmt"

type FloatValue struct {
	value float64
}

func (v FloatValue) ToInt() int64 {
	return int64(v.value)
}

func (v FloatValue) ToFloat() float64 {
	return v.value
}

func (v FloatValue) ToString() string {
	return fmt.Sprintf("%v", v.value)
}
