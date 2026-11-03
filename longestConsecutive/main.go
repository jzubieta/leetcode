package main

import "fmt"

func main() {

	nums := []int{9, 1, 4, 7, 3, 2, 8}
	set := make(map[int]struct{})
	best := 0
	for _, n := range nums {
		set[n] = struct{}{}
	}
	for n := range set {
		if _, ok := set[n-1]; ok {
			continue
		}
		length := 1
		

		for {
	if _, ok := set[n+length]; !ok {
		break
	}
	length++
}	
		
		best = max(best, length)
	}
	fmt.Println(best)
}