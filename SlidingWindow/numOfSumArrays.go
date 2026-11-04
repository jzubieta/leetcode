package main 

/*
Dado un slice de enteros arr y dos enteros k y threshold, 
devuelve cuántos subarreglos de tamaño k tienen un 
promedio mayor o igual a threshold.

*/
func numOfSubarrays(arr []int, k int, threshold int) int{
 sum:=0
 current:=0
 

 for i:=0;i<k;i++{
sum=sum+arr[i]
if sum/k>=threshold{
	current++
}}

for right:=k;right<len(arr);right++{
	sum=sum+arr[right]-arr[right-k]
	if sum/k>=threshold{
	current++
}

}

 
 return current
}


