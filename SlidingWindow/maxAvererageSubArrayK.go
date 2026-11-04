package main

func findMaxAverage(nums []int, k int) float64{
	 sum:=0
	var average float64=0
	var maxAverage float64=0
	 for i:=0;i<k;i++{
		sum=nums[i]+sum
	 }
	 average=float64(sum)/float64(k)
	 maxAverage=average

	 for right:=k; right<len(nums);right++{
		sum= sum+nums[right]-nums[right-k]
		average=float64(sum)/float64(k)
		maxAverage=max(average,maxAverage)

	 }
	 return maxAverage
}