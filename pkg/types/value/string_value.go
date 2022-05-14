package value

import (
	"fmt"
	"strconv"
)

type StringValue struct {
	value string
}

func (v StringValue) ToInt() int64 {
	i, err := strconv.ParseInt(v.value, 10, 64)

	if err != nil {
		fmt.Println(err)
	}

	return i
}

func (v StringValue) ToFloat() float64 {
	f, err := strconv.ParseFloat(v.value, 64)

	if err != nil {
		fmt.Println(err)
	}

	return f
}

func (v StringValue) ToString() string {
	return v.value
}

func (v StringValue) LessThan(v2 GenericValue) bool {
	return v.value < v2.ToString()
}

func (v StringValue) LessThanEqual(v2 GenericValue) bool {
	return v.value <= v2.ToString()
}

func (v StringValue) GreaterThan(v2 GenericValue) bool {
	return v.value > v2.ToString()
}

func (v StringValue) GreaterThanEqual(v2 GenericValue) bool {
	return v.value >= v2.ToString()
}

func (v StringValue) Equal(v2 GenericValue) bool {
	return v.value == v2.ToString()
}

func (v StringValue) Different(v2 GenericValue) bool {
	return v.value != v2.ToString()
}
