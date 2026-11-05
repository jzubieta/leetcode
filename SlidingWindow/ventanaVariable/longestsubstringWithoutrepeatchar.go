package main

func longestSubstring(s string)int{

	m:=make(map[rune]bool)
	l:=0
	longest:=0

	for r:=0;r<len(s);r++{
	

	 for m[rune(s[r])]{
		delete(m,rune(s[l]))
		l++
	 }
	m[rune(s[r])]=true
	 longest=max(longest,r-l+1)

	}
	return longest
}