package main

import (
	"fmt"
	"slices"
)

func threeSum(nums[]int)[][]int{
	k:=0
	i,j:=k+1,len(nums)-1

	r:=[][]int{}
	slices.Sort(nums)
for k < len(nums)-2{
	if k>0{
		for k<len(nums)-2&&nums[k]==nums[k-1]{
			k=k+1
		}
	}
	i=k+1
	j=len(nums)-1
	for i<j{
		sum:=nums[k]+nums[i]+nums[j]
		if sum==0{
		r = append(r, []int{nums[k], nums[i], nums[j]})
		i=i+1
		for i<len(nums)-1&&nums[i]==nums[i-1]{
		i=i+1}
		
		j=j-1
		for j>0&& nums[j]==nums[j+1]{
			j=j-1
		}

		
		}else{
			if sum>0{
				j=j-1
			}else{
				i=1+i
			}
		}
	}
	
	
		k=k+1
	
	}
	return r
}


func main() {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"vacío", []int{}, [][]int{}},
		{"menos de tres números", []int{0, 0}, [][]int{}},
		{"un trío de ceros", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{"ceros repetidos", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{"duplicados en los extremos", []int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
		{"dos tríos", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"dos parejas posibles", []int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
		{"sin solución", []int{1, 2, -2, -1}, [][]int{}},
	}

	for _, test := range tests {
		got := threeSum(test.nums)
		fmt.Printf("%s\n  obtuvo:   %v\n  esperado: %v\n", test.name, got, test.want)
	}
}