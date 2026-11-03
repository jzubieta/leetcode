package main
func kidsWithCandies(candies []int, extraCandies int) []bool {
    greatest:=0
result:=make([]bool,len(candies))
	for _,v:= range candies{
		if v+extraCandies>greatest{
			greatest=v+extraCandies
			result = append(result, true)
		}else{
			result=append(result,false)
		}
	}
	return result
}