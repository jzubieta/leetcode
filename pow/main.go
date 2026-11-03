package main

import (
	"fmt"

	
)

func myPow(x float64, n int) float64 {
   if n==0{
	return 1
   }
	if n>0{
		return pow(x,n)
	}
	return 1/pow(x,n*-1)

}

func pow(x float64,n int) float64{
	
var result float64 =1
	for i:=1;i<=n;i++{
			result=result*x
			
		}
		return result
	}

func main(){
	fmt.Print(myPow(1,-2))
}