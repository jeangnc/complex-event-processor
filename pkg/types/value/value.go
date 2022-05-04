package value

import "fmt"

func NewValue(v interface{}) GenericValue {
	switch v := v.(type) {
	case float64:
		return FloatValue{value: v}
	case int64:
		return IntValue{value: v}
	case string:
		return StringValue{value: v}
	default:
		fmt.Println(fmt.Sprintf("I don't know about type %T!\n", v))
		return nil
	}
}

type GenericValue interface {
	ToInt() int64
	ToFloat() float64
	ToString() string
}

type Ordered interface {
	LessThan(GenericValue) bool
	LessThanEqual(GenericValue) bool
	GreaterThan(GenericValue) bool
	GreaterThanEqual(GenericValue) bool
	Equal(GenericValue) bool
	Different(GenericValue) bool
}
