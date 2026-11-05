package main

func sumaArray(nums []int ,target int)int{

	l:= 0
	lenMin:=len(nums)+1
	sumaVentana:=0
	for r:=0;r<len(nums);r++{
		sumaVentana=sumaVentana+nums[r]
		for sumaVentana>= target{
			lenMin=min(lenMin,r-l+1)
			sumaVentana=sumaVentana-nums[l]
			l++
		}

	}
	if lenMin== len(nums)+1{
		return 0
	}
	return lenMin

}