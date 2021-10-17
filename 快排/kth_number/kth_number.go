package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d%d", &n, &k)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	fmt.Print(get_kth_numeber(0, n-1, k-1, q))
}

func get_kth_numeber(l, r, k int, q []int) int {
	if l == r {
		return q[l]
	}
	x := q[(l+r)>>1]
	i, j := l-1, r+1
	for i < j {
		for {
			i++
			if q[i] >= x {
				break
			}
		}
		for {
			j--
			if q[j] <= x {
				break
			}
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}
	if k<=j{
		return get_kth_numeber(l,j,k,q)
	}
	return get_kth_numeber(j+1, r, k, q)
}
