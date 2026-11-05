package main

import (
	"math"

)

func squares (nums [] int)[]int{
	result:=make([]int,len(nums))
    write:=len(result)-1
	l:=0
	r:= len(nums)-1
	for l<=r{
		if math.Pow(float64(nums[l]),2)>math.Pow(float64(nums[r]),2){
					result[write]=int(math.Pow(float64(nums[l]),2))
					l++		
		}else{
			result[write]=int(math.Pow(float64(nums[r]),2))
			r--
		}
		write--
	}
	return result
}