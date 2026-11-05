package main
func intersection2Array(nums1,nums2 [] int)[]int{
result:=[]int{}
m:= make(map[int]bool)

for _,num:= range nums1{
	m[num]=true
}
for _,num2:= range nums2{
	if _,ok:=m[num2];ok{
		result = append(result, num2)
		delete(m,num2)

	}
}
return result
}