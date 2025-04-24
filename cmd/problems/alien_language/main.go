// In an alien languague, they use english lowecase letters, but possible in a different order.
// The order of the alphabet is some permutation of lowecase letters.
// Given a sequence of words in the alien languague, and the order of the alphabet,
// Return true if the are sorted lexicographically in this alien languague

package main

import "fmt"

func main() {

	order := "huabcdefgijklmnopqrstvwxyz"
	input1 := []string{"hello", "uber"}

	fmt.Println(isLexicoSorted(order, input1))
	// output -> true

	input3 := []string{"hello", "hell"}
	fmt.Println(isLexicoSorted(order, input3))
	// output -> false

	order = "worldabcefghijkmnpqstuvxyz"
	input2 := []string{"word", "world", "row"}

	fmt.Println(isLexicoSorted(order, input2))
	// output -> false

}

func isLexicoSorted(order string, words []string) bool {
	rank := make(map[byte]int)

	// O(26) --> O(1)
	for i := 0; i < len(order); i++ {
		rank[order[i]] = i
	}

	// O(n * m) ->
	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]

		for j := 0; j < len(w1) && j < len(w2); j++ {
			c1, c2 := w1[j], w2[j]

			if c1 != c2 {
				if rank[c1] > rank[c2] {
					return false
				}

				break
			}
		}

		if (len(w1) > len(w2)) && (w1[:len(w2)] == w2) {
			return false
		}
	}

	return true
}
