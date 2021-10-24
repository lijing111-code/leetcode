package main

import "fmt"

func main(){
	var m,n int
	fmt.Scanf("%d%d",&m,&n)
	q:=make([]int,m)
	for i:=0;i<m;i++{
		fmt.Scanf("%d",&q[i])
	}
	s:=make([]int ,0)
	s = append(s, 0)
	sum:=0
	for i:=0;i<m;i++{
		sum+=q[i]
		s = append(s, sum)
	}
	for i:=0;i<n;i++{
		var l ,r int
		fmt.Scanf("%d%d",&l,&r)
		fmt.Println(prefixSum(s,l,r))
	}
}

func prefixSum(q []int,l,r int)int {
	return q[r]-q[l-1]
}

