package main
func minSubArrayLen(target int, nums []int) int {
 sum:=0
 minLen:=320000
 best:=0
l:=0;
for right:=0;right<len(nums);right++{
	 sum=nums[right]+sum

	for sum <=target{
		sum=sum-nums[l]
		l++
	}
	best=min(minLen,l-right+1)
}
return best
}