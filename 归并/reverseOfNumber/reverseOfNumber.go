package main

import "fmt"

func main(){
	var n int
	fmt.Scanf("%d",&n)
	q :=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scanf("%d",&q[i])
	}
	fmt.Println(reverseOfNumber(q,0,n-1))
}
func reverseOfNumber(q []int ,l ,r int) int{
	if l>=r{
		return 0
	}
	mid :=(l+r)>>1
	cnt:=reverseOfNumber(q,l,mid)+reverseOfNumber(q,mid+1,r)
	tmp :=make([]int,0)
	i,j:=l,mid+1
	for i<=mid && j<=r{
		if q[i]<=q[j]{
			tmp = append(tmp, q[i])
			cnt+=mid-i+1
			i++
		}else{
			tmp = append(tmp, q[j])
			j++
		}
	}
	for i<=mid{
		tmp = append(tmp, q[i])
		i++
	}
	for j<=r{
		tmp = append(tmp, q[j])
		j++
	}

	//赋值回原数组
	i,j=l,0
	for i<=r{
		q[i]=tmp[j]
		i++
		j++
	}
	return cnt
}