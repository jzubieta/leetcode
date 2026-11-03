package main

import "fmt"

func twoSum(numbers []int, target int)[]int{
seen:= make(map[int]int)
for i,v:= range numbers{
	complement:=target-v
	if index,ok:=seen[complement];ok{
		return []int{i,index}
	}else{
		seen[v]=i
	}
}
return []int{-1,-1}

}

func main(){

	test:= []int{1,2,3,4}
	target:=7

	fmt.Println(twoSum(test,target))
}