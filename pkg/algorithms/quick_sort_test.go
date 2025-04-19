package algorithms

// func Test_quickSort(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    []int
// 		expected []int
// 	}{
// 		// {
// 		// 	name:     "empty array",
// 		// 	input:    []int{},
// 		// 	expected: []int{},
// 		// },
// 		// {
// 		// 	name:     "single element",
// 		// 	input:    []int{1},
// 		// 	expected: []int{1},
// 		// },
// 		// {
// 		// 	name:     "already sorted array",
// 		// 	input:    []int{1, 2, 3, 4, 5},
// 		// 	expected: []int{1, 2, 3, 4, 5},
// 		// },
// 		{
// 			name:     "reverse sorted array",
// 			input:    []int{5, 4, 3, 2, 1},
// 			expected: []int{1, 2, 3, 4, 5},
// 		},
// 		// {
// 		// 	name:     "array with duplicates",
// 		// 	input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
// 		// 	expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
// 		// },
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			input := make([]int, len(tt.input))
// 			copy(input, tt.input)

// 			if len(input) > 0 {
// 				quickSort(input, 0, len(input)-1)
// 			}

// 			if !reflect.DeepEqual(input, tt.expected) {
// 				t.Errorf("quickSort() = %v, want %v", input, tt.expected)
// 			}
// 		})
// 	}
// }
