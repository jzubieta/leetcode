
package main
func longestPalindrome(s string) int {
	counts := make(map[rune]int)

	for _, r := range s {
		counts[r]++
	}

	length := 0
	hasOdd := false

	for _, frequency := range counts {
		length += (frequency / 2) * 2
		if frequency%2 == 1 {
			hasOdd = true
		}
	}

	if hasOdd {
		length++
	}

	return length
}