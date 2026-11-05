package main


func isVowelthis(c rune) bool{

	

		if c=='a'||c=='e'|| c=='i'||c=='o'||c=='u'{
			return true
		}
		return false
	}


	func maxVowels(s string ,k int)int{

		count:=0
		maxCounter:=0
	

		for i:=0 ; i <k;i++{
			if isVowelthis(rune(s[i])){
				count++
			}
		}
		maxCounter=count
	for rigth:=k; rigth<len(s);rigth++{
		if isVowelthis(rune(s[rigth])){
			count++
		}
		if isVowelthis(rune(s[rigth-k])){
			count--
		}


maxCounter=max(maxCounter,count)

	}
	return maxCounter
}