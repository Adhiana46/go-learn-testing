package math

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAbsolute(t *testing.T) {
	t.Run("negative number", func(t *testing.T) {
		result := Absolute(-5)
		assert.Equal(t, 5, result, "expect 5, but got %d", result)
	})

	t.Run("positive number", func(t *testing.T) {
		result := Absolute(5)
		assert.Equal(t, 5, result, "expect 5, but got %d", result)
	})
}

func TestAbsolute_FlakyTest(t *testing.T) {
	// skip untuk testing yg kadang berhasil kadang gagal entah kenapa
	t.Skip()
}

func TestAbsolute_Short(t *testing.T) {
	/**
	digunakan untuk skip testing:
	- yang memakan waktu agak lama
	- integration test

	dijalankan dengan perintah
		go test ./... -short
	*/
	if testing.Short() {
		t.Skip()
	}

	<-time.After(5 * time.Second)

	t.Run("negative number", func(t *testing.T) {
		result := Absolute(-5)
		assert.Equal(t, 5, result, "expect 5, but got %d", result)
	})

	t.Run("positive number", func(t *testing.T) {
		result := Absolute(5)
		assert.Equal(t, 5, result, "expect 5, but got %d", result)
	})
}
