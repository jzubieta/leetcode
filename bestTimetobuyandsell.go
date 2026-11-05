package main

func bttobas( prices[]int)int{
if len(prices)==0{return 0}
	minprice:=prices[0]
	profit:=0
	maxProfit:=0
	for i:=0;i<len(prices);i++{
		profit= prices[i]-minprice
 maxProfit=max(maxProfit,profit)
 minprice=min(minprice,prices[i])
	}

	return maxProfit
}