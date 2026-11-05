package main


func longestOnes(nums [] int, k int)int{

	l:=0
	maxSize:=0
	countZeroes:=0

	for r:=0; r< len(nums);r++{
		
		if  nums[r]==0 {
			countZeroes++
		}
		for countZeroes>k{
		
		if  nums[l]==0 {
			countZeroes--
		}
		l++
		}
		maxSize=max(maxSize,r-l+1)
	}
return maxSize
}