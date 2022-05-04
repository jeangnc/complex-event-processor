package value

import (
	"testing"
)

func TestFloatToInt(t *testing.T) {
	v := FloatValue{value: 1.0}
	i := v.ToInt()

	if i != 1 {
		t.Fatalf(`Failed to cast to int: %v %v`, v, i)
	}
}

func TestFloatToFloat(t *testing.T) {
	v := FloatValue{value: 1.0}
	i := v.ToFloat()

	if i != 1.0 {
		t.Fatalf(`Failed to cast to int: %v %v`, v, i)
	}
}

func TestFloatToString(t *testing.T) {
	v := FloatValue{value: 1.0}
	i := v.ToString()

	if i != "1" {
		t.Fatalf(`Failed to cast to int: %v %v`, v, i)
	}
}
