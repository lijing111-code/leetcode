package main

import "fmt"
//归并排序：
//1、确定分界点：mid=(l+r)/2
//2、递归出子问题
//3、再归并（将两个有序数组合并成一个）-------重点
func main(){
	var n int
	fmt.Scanf("%d",&n)
	q :=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scanf("%d",&q[i])
	}
	mergeSort(q,0,n-1)
	for _,num:=range q{
		fmt.Printf("%d ",num)
	}
}

func mergeSort(q []int,l,r int){
	if l>=r{
		return
	}
	mid :=(l+r)>>1
	mergeSort(q,l,mid)
	mergeSort(q,mid+1,r)

	//双指针法合并有序数组
	tmp :=make([]int,0)
	i,j:=l,mid+1
	for i<=mid && j<=r{
		if q[i]<=q[j]{
			tmp = append(tmp, q[i])
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
}
