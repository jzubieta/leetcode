package main



func majorityElement(nums []int) int {
    
	frequency:=make(map[int]int)

		for _,num := range nums{
			frequency[num]++
		}
		for k,v:= range frequency{
			if v> len(nums)/2{
				return k
			}
		}
		return 0
	}
