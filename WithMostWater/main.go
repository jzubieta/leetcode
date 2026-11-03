package main

import (
	

)

func withMostWater(heigths []int)int{

	current,biggest,j,i:=0,0,len(heigths)-1,0

	for i<j{
			
			base:= j-i
			
			maxHeigth:=min(heigths[i],heigths[j])
            current= base*maxHeigth
			biggest=max(biggest,current)
			if(heigths[i]<heigths[j]){
				i=i+1
			}else{
				j=j-1
			}
			
		}
		

	
return biggest
}
func main(){
withMostWater([]int{1,8,6,2,5,4,8,3,7})
}