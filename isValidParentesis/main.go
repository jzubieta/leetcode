package main

import "fmt"

func isValid(s string) bool {
stack:= [] rune{}

   matches := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }
	for _,v:=range s{
		if v=='('||v=='['||v=='{'{
			stack=append(stack, v)
		}
		if v==')'||v==']'||v=='}'{
			if len(stack)==0{
				return false
			}
			lastOpen:=stack[len(stack)-1]
			t:=matches[v]
			if t!=lastOpen{
				return false
			}
			stack=stack[:len(stack)-1]

		}

}
if len(stack)==0{
	return true
}else{
	return false
}}


func main(){
	fmt.Println(isValid(")))]](([["))
}