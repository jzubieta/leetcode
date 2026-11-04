package main

/*
Dadas dos cadenas, s y p, devuelve un arreglo con 
todos los índices de inicio de los anagramas de p en s. 
El tamaño de la ventana fija es la longitud de p.
*/

func findAllAnagrams(s,p string)[]int{
	k:= len(p)
	result:=[]int{}
	runs:=[]rune{}

	for i:=0 ;i< k ;i++{
		runs=append(runs,rune(s[i]))
		

	}
	if isAnagram(string(runs),p){
			result=append(result, 0)
		}

for rigth:=k ;rigth<len(s);rigth++{
	runs=runs[1:]
	runs = append(runs, rune(s[rigth]))
	if isAnagram(string(runs),p){
		result=append(result,rigth - k + 1)
	}


}
return result
}

func isAnagram(a,b string)bool{

	m:=make(map[rune]int)
for _,char:= range a{
	m[char]++
}
for _,char:=range b{
	m[char]--
}
for _,v:=range m{
	if v!=0{
		return false
	}
}

return true
}