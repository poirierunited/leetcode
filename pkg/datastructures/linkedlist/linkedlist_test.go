package linkedlist

import "testing"

func TestLinkedList_Add(t *testing.T) {
	tests := []struct {
		name           string
		values         []int
		expectedLength int
		expectedHead   int
		expectedTail   int
	}{
		{
			name:           "empty list",
			values:         []int{},
			expectedLength: 0,
			expectedHead:   0,
			expectedTail:   0,
		},
		{
			name:           "single element",
			values:         []int{1},
			expectedLength: 1,
			expectedHead:   1,
			expectedTail:   1,
		},
		{
			name:           "multiple elements",
			values:         []int{1, 2, 3, 4, 5},
			expectedLength: 5,
			expectedHead:   1,
			expectedTail:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := NewLinkedList()

			// Add values to the list
			for _, v := range tt.values {
				ll.Add(v)
			}

			// Check length
			if ll.Length != tt.expectedLength {
				t.Errorf("Length = %v, want %v", ll.Length, tt.expectedLength)
			}

			// Check head
			if tt.expectedLength > 0 {
				if ll.Head == nil {
					t.Error("Head is nil")
				} else if ll.Head.Value != tt.expectedHead {
					t.Errorf("Head.Value = %v, want %v", ll.Head.Value, tt.expectedHead)
				}
			} else if ll.Head != nil {
				t.Error("Head should be nil for empty list")
			}

			// Check tail
			if tt.expectedLength > 0 {
				if ll.Tail == nil {
					t.Error("Tail is nil")
				} else if ll.Tail.Value != tt.expectedTail {
					t.Errorf("Tail.Value = %v, want %v", ll.Tail.Value, tt.expectedTail)
				}
			} else if ll.Tail != nil {
				t.Error("Tail should be nil for empty list")
			}

			// Verify the list structure
			if tt.expectedLength > 0 {
				current := ll.Head
				for i, expected := range tt.values {
					if current == nil {
						t.Errorf("expected node at index %d, got nil", i)
						return
					}
					if current.Value != expected {
						t.Errorf("at index %d: expected %d, got %d", i, expected, current.Value)
					}
					current = current.Next
				}
				if current != nil {
					t.Error("list has more elements than expected")
				}
			}
		})
	}
}

func TestLinkedList_Pop(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		popCount       int
		expectedLength int
		expectedHead   *int
		expectedTail   *int
	}{
		{
			name:           "empty list",
			initialValues:  []int{},
			popCount:       1,
			expectedLength: 0,
			expectedHead:   nil,
			expectedTail:   nil,
		},
		{
			name:           "single element",
			initialValues:  []int{1},
			popCount:       1,
			expectedLength: 0,
			expectedHead:   nil,
			expectedTail:   nil,
		},
		{
			name:           "multiple elements - pop one",
			initialValues:  []int{1, 2, 3},
			popCount:       1,
			expectedLength: 2,
			expectedHead:   intPtr(1),
			expectedTail:   intPtr(2),
		},
		{
			name:           "multiple elements - pop all",
			initialValues:  []int{1, 2, 3},
			popCount:       3,
			expectedLength: 0,
			expectedHead:   nil,
			expectedTail:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := NewLinkedList()

			// Add initial values
			for _, v := range tt.initialValues {
				ll.Add(v)
			}

			// Pop elements
			for i := 0; i < tt.popCount; i++ {
				ll.Pop()
			}

			// Check length
			if ll.Length != tt.expectedLength {
				t.Errorf("Length = %v, want %v", ll.Length, tt.expectedLength)
			}

			// Check head
			if tt.expectedHead == nil {
				if ll.Head != nil {
					t.Error("Head should be nil")
				}
			} else {
				if ll.Head == nil {
					t.Error("Head is nil")
				} else if ll.Head.Value != *tt.expectedHead {
					t.Errorf("Head.Value = %v, want %v", ll.Head.Value, *tt.expectedHead)
				}
			}

			// Check tail
			if tt.expectedTail == nil {
				if ll.Tail != nil {
					t.Error("Tail should be nil")
				}
			} else {
				if ll.Tail == nil {
					t.Error("Tail is nil")
				} else if ll.Tail.Value != *tt.expectedTail {
					t.Errorf("Tail.Value = %v, want %v", ll.Tail.Value, *tt.expectedTail)
				}
			}
		})
	}
}

// Helper function to create int pointer
func intPtr(i int) *int {
	return &i
}
