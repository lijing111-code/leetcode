package main

import "fmt"

//模板
func main() {
	var n, p int
	fmt.Scanf("%d%d", &n, &p)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	num:=make([]int ,p)
	for i:=0;i<p;i++{
		fmt.Scanf("%d",&num[i])
	}
	for _,i:= range num{
		left :=getleft(q,0,n-1,i)
		right :=getright(q,0,n-1,i)
		if q[left]==i{
			fmt.Println(left,right)
		}else{
			fmt.Println(-1,-1)
		}
	}
}

func getleft(q []int, l, r, num int) int{
	for l<r{
		mid :=(l+r)/2
		if q[mid]>=num{
			r = mid
		}else{
			l=mid+1
		}
	}
	return l
}

func getright(q []int, l, r, num int) int {
	for l<r{
		mid :=(l+r+1)/2
		if q[mid]<=num{
			l = mid
		}else{
			r=mid-1
		}
	}
	return l
}


//自己写的版本-------------思路没问题，处理边界条件太多


//package main
//
//import "fmt"

//func main() {
//	var n, p int
//	fmt.Scanf("%d%d", &n, &p)
//	q := make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scanf("%d", &q[i])
//	}
//	num:=make([]int ,p)
//	for i:=0;i<p;i++{
//		fmt.Scanf("%d",&num[i])
//	}
//	for _,i:= range num{
//		fmt.Println(rangeOfNumber(q,0,n-1,i))
//	}
//}
//
//func rangeOfNumber(q []int, l, r, num int)(int ,int) {
//	if q[l]==q[r]{
//		if q[l]==num{
//			return l,r
//		}
//		return -1,-1
//	}
//	if num<q[l]||num>q[r] || l>r{
//		return -1,-1
//	}
//	if l==r{
//		if q[l]!=num{
//			return -1,-1
//		}
//		return l,l
//	}
//
//	mid :=(l+r+1)>>1
//	midN :=q[mid]
//	var l1, r1 int
//	if midN>num{
//		l1 ,r1 =rangeOfNumber(q,l,mid-1,num)
//	}else if midN==num{
//		l1,_ =rangeOfNumber(q,l,mid-1,num)
//		if l1==-1{
//			l1=mid
//		}
//		_,r1=rangeOfNumber(q,mid,r,num)
//	}else{
//		l1 ,r1 =rangeOfNumber(q,mid+1,r,num)
//	}
//
//
//	return l1, r1
//}