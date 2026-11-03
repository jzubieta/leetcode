package main

import "fmt"

func main(){
nums:= []int{2,3,4,5}
n:=len(nums)
l:=make([]int,n)
r:=make([]int,n)
result:=make([]int,n)

l[0]=1
r[n-1]=1
for i:=1;i<n;i++{
	l[i]=l[i-1]*nums[i-1]
	
}
for i:=n-2;i>=0;i--{
	r[i]=r[i+1]*nums[i+1]
}
for i:=0;i<n;i++{
	result[i]=r[i]*l[i]
}
fmt.Println(result)
}