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

func (v FloatValue) LessThan(v2 GenericValue) bool {
	return v.value < v2.ToFloat()
}

func (v FloatValue) LessThanEqual(v2 GenericValue) bool {
	return v.value <= v2.ToFloat()
}

func (v FloatValue) GreaterThan(v2 GenericValue) bool {
	return v.value > v2.ToFloat()
}

func (v FloatValue) GreaterThanEqual(v2 GenericValue) bool {
	return v.value >= v2.ToFloat()
}

func (v FloatValue) Equal(v2 GenericValue) bool {
	return v.value == v2.ToFloat()
}

func (v FloatValue) Different(v2 GenericValue) bool {
	return v.value != v2.ToFloat()
}
