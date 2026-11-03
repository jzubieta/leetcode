package main

import (
	"fmt"
	"sync"
)
func main(){
	fmt.Println(collect([]int{1,2,3}))

}
func collect(items []int) []int {
ch := make(chan int)
var wg sync.WaitGroup
for _, item := range items {
wg.Add(1)
go func(x int) {
defer wg.Done()
ch <- x * 2
}(item)
}
go func(){
wg.Wait()
close(ch)}()
var out []int
for x := range ch {
out = append(out, x)
}
return out
}