package main 
func lengthOfLongestSubstring(s string) int {
    count := map[byte]int{}
    left := 0
    best := 0

    for right := 0; right < len(s); right++ {
        // entra s[right]
		count[s[right]]=count[s[right]]+1

        for  count[s[right]]>1{
            // sale s[left]
			count[s[left]]--
            // left++
			left++
        }

        // registra
		best = max(best, right-left+1)
    }
    return best
}