package main

func maxvowel(s string,k int)int{

	currentCount:=0
	maxCount:=0
	for i:=0;i<k;i++{
		if isVowel(rune(s[i])){
			currentCount=currentCount+1
		}
	}
	maxCount=currentCount
	for right:=k;right<len(s);right++{
		if isVowel(rune(s[right])){
			currentCount++
		}
		 if isVowel(rune(s[right-k])){
			currentCount--
		}
		maxCount=max(maxCount,currentCount)
	}

return maxCount

}

func isVowel( char rune)bool{
 if char=='a'||char=='e'||char=='i'||char=='o'||char=='u'{
	return true
 }
 return false
}