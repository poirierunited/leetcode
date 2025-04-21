package searching

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "empty array",
			nums:     []int{},
			target:   5,
			expected: -1,
		},
		{
			name:     "single element found",
			nums:     []int{5},
			target:   5,
			expected: 0,
		},
		{
			name:     "single element not found",
			nums:     []int{5},
			target:   3,
			expected: -1,
		},
		{
			name:     "target in middle",
			nums:     []int{1, 3, 5, 7, 9},
			target:   5,
			expected: 2,
		},
		{
			name:     "target at start",
			nums:     []int{1, 3, 5, 7, 9},
			target:   1,
			expected: 0,
		},
		{
			name:     "target at end",
			nums:     []int{1, 3, 5, 7, 9},
			target:   9,
			expected: 4,
		},
		{
			name:     "target not found",
			nums:     []int{1, 3, 5, 7, 9},
			target:   4,
			expected: -1,
		},
		{
			name:     "duplicate elements",
			nums:     []int{1, 2, 2, 2, 3},
			target:   2,
			expected: 2, // returns one of the occurrences
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("BinarySearch(%v, %d) = %d; want %d", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
