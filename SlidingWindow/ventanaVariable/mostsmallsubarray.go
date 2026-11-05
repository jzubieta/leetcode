package main





func mostSmall(arr[]int ,target int) int{

l:=0
minim:=len(arr)+1
windowSum:=0
for rigth:=0;rigth<len(arr);rigth++{
	windowSum=windowSum+arr[rigth]
	for windowSum>=target{
		minim=min(minim, rigth-l+1)
		windowSum=windowSum-arr[l]
		l++
	}
}
if minim==len(arr)+1{
	return 0
}
return int(minim)
}