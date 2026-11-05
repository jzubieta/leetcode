package main
func checkInclusion(s1,s2 string) bool{

	k:=len(s1)
	window:=[]rune{}
	
	if len(s1)>len(s2){
		return false
	}
for i:=0;i<k;i++{
	window=append(window, rune(s2[i]))
}
if (isPermutation(s1,string(window))){
	return true
}
for r:=k;r<len(s2);r++{
	window= window[1:]
	window=append(window, rune(s2[r]))
	if isPermutation(s1,string(window)){
return true
	}
}
return false
}


func isPermutation(s1,window string)bool{
	s:=make(map[rune]int)
	w:=make(map[rune]int)
	for _,cs1:=range s1{
		s[cs1]++
	}
for _,w1:= range window{
	w[w1]++
}

for k,v:= range w{
	if s[k]!=v{
		return false
	}
}
return len(s1)==len(window)
}