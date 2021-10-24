package main

import (
	"fmt"
	"math"
)
var l ,r float64=-10000,10000

//数的三次方根：给定一个浮点数，求它的三次方根
func main(){
	var n float64
	fmt.Scanf("%f",&n)

	fmt.Printf("%6f",getCubicRoots(l,r,n))
}

func getCubicRoots(l,r ,n float64) float64{
	for r-l>math.Pow(10,-8){
		mid:=(l+r)/2
		if math.Pow(mid,3)>=n{
			r=mid
		}else{
			l=mid
		}
	}
	return l
}
