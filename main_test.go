package main

// Пишите тесты в этом файле
import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	testCases := []struct {
		size int
	}{
		{size: 0},
		{size: 1},
		{size: 10},
		{size: 1000},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Test case with size %d", tc.size), func(t *testing.T) {
			result := generateRandomElements(tc.size)

			require.Len(t, result, tc.size)

			for _, v := range result {
				assert.GreaterOrEqual(t, v, 0)
				assert.LessOrEqual(t, v, SIZE)
			}

		})
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

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", i+1), func(t *testing.T) {
			result := maximum(tc.data)
			assert.Equal(t, tc.expected, result)
		})
	}
}
