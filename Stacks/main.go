package main

import "fmt"

func isValid(s string) bool{
stack :=[]rune{}

matches:=map[rune]rune{
	')':'(',
	']':'[',
	'}':'{',
}

for _,char:=range s{
	if char=='('||char=='{'||char=='['{
		stack=append(stack,char)
	}
	if char==')'||char=='}'||char==']'{
		if len(stack)==0{
			return false
		}
		lastOpen:= stack[len(stack)-1]
		if matches[char]!=lastOpen{
			return false
		}else{
			stack = stack[:len(stack)-1]
		}

	}
	
}
if len(stack)==0{
		return true
	}
	return false

}

func main(){
	fmt.Println(isValid("(())"))
}