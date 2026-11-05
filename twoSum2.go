package main 


func twoSumII ( nums[] int, target int)[]int{

	rigth:=0
	left:=len(nums)-1

	for rigth<left{
		if nums[rigth]+nums[left]> target{
			left--
		}else if nums[rigth]+nums[left]<target{
			rigth++
		}else {
			return []int{rigth+1,left+1}
		
	}
}
return []int{0,0}

}