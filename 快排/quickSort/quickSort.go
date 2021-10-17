package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d",&n)
	arr :=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scanf("%d",&arr[i])
	}
	quickSort(0,n-1,arr)
	fmt.Println(arr)
}

func quickSort(l ,r int, arr []int){
	if l>=r{// 起始条件
		return
	}
	x :=arr[l]
	i,j:=l-1,r+1// 两个指针，因为do while要先自增/自减
	for i<j{
		for{
			i++
			if arr[i]>=x{
				break
			}
		}
		for{
			j--
			if arr[j]<=x{
				break
			}
		}
		if i<j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	quickSort(l,j,arr)// 递归处理左右两段
	quickSort(j+1,r,arr)
}
