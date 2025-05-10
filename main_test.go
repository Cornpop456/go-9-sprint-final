package main

// Пишите тесты в этом файле
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	result := generateRandomElements(0)
	assert.Equal(t, []int{}, result)

	result = generateRandomElements(10)
	assert.Len(t, result, 10)

	for _, v := range result {
		assert.GreaterOrEqual(t, v, 0)
		assert.LessOrEqual(t, v, SIZE)
	}
}

func TestMaximum(t *testing.T) {
	testCases := []struct {
		data     []int
		expected int
	}{
		{data: []int{}, expected: 0},
		{data: []int{1}, expected: 1},
		{data: []int{2, 2, 2, 2}, expected: 2},
		{data: []int{1, 2, 3, 4, 5, 6}, expected: 6},
		{data: []int{6, 5, 4, 3, 2, 1}, expected: 6},
	}

	for _, tc := range testCases {
		result := maximum(tc.data)
		assert.Equal(t, tc.expected, result)
	}
}
