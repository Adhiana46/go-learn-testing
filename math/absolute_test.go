package math

import "testing"

func TestAbsolute(t *testing.T) {
	result := Absolute(-5)

	if result != 5 {
		t.Logf("expect 5, but got %d", result)
		t.Fail()
	}
}

func TestAbsolute_WithPositiveNum(t *testing.T) {
	result := Absolute(5)

	if result != 5 {
		t.Logf("expect 5, but got %d", result)
		t.Fail()
	}
}
