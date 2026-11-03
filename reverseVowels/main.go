package main

import "fmt"




func reverseVowels(s string) string {
   vowels:=[] rune{}




for _,letter:= range s{

	if letter=='a'||letter=='e'||letter=='i'||letter=='o'||letter=='u'||letter=='A'||letter=='E'||letter=='I'||letter=='O'||letter=='U'{

vowels=append(vowels,letter)
}
}
ru:=[]rune(s)
for i:=0;i<len(ru);i++{
letter:=s[i]

	if len(vowels) > 0 {
    ultimo := vowels[len(vowels)-1]
    
	if letter=='a'||letter=='e'||letter=='i'||letter=='o'||letter=='u'||letter=='A'||letter=='E'||letter=='I'||letter=='O'||letter=='U'{
ru[i]=ultimo
vowels = vowels[:len(vowels)-1]
	}

}
}

return string(ru)
}
func main(){
	fmt.Println(reverseVowels("IceCreAm"))
}
