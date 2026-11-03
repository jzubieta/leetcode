package main

import (
	"fmt"
	"math"
)
func angleClock(hour int, minutes int) float64 {
angleByMinute:= 360/60
angleByHour:=360/12

ah := float64(hour%12)*float64(angleByHour) + float64(minutes)*0.5
am := float64(minutes) * float64(angleByMinute)

angle :=math.Abs(float64(ah)-float64(am))

if angle>180{
	return 360-angle
}
    return angle
}

func main(){
	fmt.Println(angleClock(12,30))
}