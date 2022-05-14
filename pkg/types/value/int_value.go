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

func (v IntValue) LessThan(v2 GenericValue) bool {
	return v.value < v2.ToInt()
}

func (v IntValue) LessThanEqual(v2 GenericValue) bool {
	return v.value <= v2.ToInt()
}

func (v IntValue) GreaterThan(v2 GenericValue) bool {
	return v.value > v2.ToInt()
}

func (v IntValue) GreaterThanEqual(v2 GenericValue) bool {
	return v.value >= v2.ToInt()
}

func (v IntValue) Equal(v2 GenericValue) bool {
	return v.value == v2.ToInt()
}

func (v IntValue) Different(v2 GenericValue) bool {
	return v.value != v2.ToInt()
}
