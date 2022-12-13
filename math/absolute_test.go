package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsolute(t *testing.T) {
	result := Absolute(-5)

	assert.Equal(t, 5, result, "expect 5, but got %d", result)
}

func TestAbsolute_WithPositiveNum(t *testing.T) {
	result := Absolute(5)

	assert.Equal(t, 5, result, "expect 5, but got %d", result)
}
