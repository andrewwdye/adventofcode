package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNext(t *testing.T) {
	assert.Equal(t, 18, next([]int{0, 3, 6, 9, 12, 15}))
	assert.Equal(t, 28, next([]int{1, 3, 6, 10, 15, 21}))
	assert.Equal(t, 68, next([]int{10, 13, 16, 21, 30, 45}))
}
