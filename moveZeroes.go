package main

func moveZeroes(nums[]int){

	contador:=0

	for _,num:=range nums{

		if num!=0{
			nums[contador]=num
			contador++
		}

	}
for i:=contador;i<len(nums);i++{
	nums[i]=0
}

}