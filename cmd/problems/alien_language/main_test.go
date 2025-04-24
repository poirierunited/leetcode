package main

import "testing"

func TestIsLexicoSorted(t *testing.T) {
	tests := []struct {
		name  string
		order string
		words []string
		want  bool
	}{
		{
			name:  "basic sorted case",
			order: "huabcdefgijklmnopqrstvwxyz",
			words: []string{"hello", "uber"},
			want:  true,
		},
		{
			name:  "unsorted case",
			order: "worldabcefghijkmnpqstuvxyz",
			words: []string{"word", "world", "row"},
			want:  false,
		},
		{
			name:  "prefix case",
			order: "huabcdefgijklmnopqrstvwxyz",
			words: []string{"hello", "hell"},
			want:  false,
		},
		{
			name:  "empty words",
			order: "abcdefghijklmnopqrstuvwxyz",
			words: []string{},
			want:  true,
		},
		{
			name:  "single word",
			order: "abcdefghijklmnopqrstuvwxyz",
			words: []string{"hello"},
			want:  true,
		},
		{
			name:  "same words",
			order: "abcdefghijklmnopqrstuvwxyz",
			words: []string{"hello", "hello"},
			want:  true,
		},
		{
			name:  "different lengths",
			order: "abcdefghijklmnopqrstuvwxyz",
			words: []string{"a", "ab", "abc"},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isLexicoSorted(tt.order, tt.words)
			if got != tt.want {
				t.Errorf("isLexicoSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
