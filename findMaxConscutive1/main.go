package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
    counter:=0
    maxc:=0
    for _,n:=range nums{
        if n==1{
            counter=counter+1
        }else{
             
             counter=0
        }
		 maxc=max(maxc,counter)
    }

return maxc
}

func main(){

	fmt.Println(findMaxConsecutiveOnes([]int{1,1,0,1,1,1}))
}