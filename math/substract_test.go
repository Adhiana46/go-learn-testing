package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstract(t *testing.T) {
	// use test table
	testCases := []struct {
		name           string
		a, b, expected int
	}{
		{
			name:     "positive x negative",
			a:        1,
			b:        -1,
			expected: 2,
		},
		{
			name:     "positive x positive",
			a:        1,
			b:        1,
			expected: 0,
		},
		{
			name:     "negative x negative",
			a:        -1,
			b:        -1,
			expected: 0,
		},
		{
			name:     "negative x positive",
			a:        -1,
			b:        1,
			expected: -2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Substract(testCase.a, testCase.b)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
