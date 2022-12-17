package StringToFloat

import (
	"testing"
)

func TestGet(t *testing.T) {
	a := "123.456"
	b := GetFloat64(a)
	if b != 123.456 {
		t.Error("not work")
	}
}
