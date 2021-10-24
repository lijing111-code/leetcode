package main

import "fmt"

func main() {
	var str1, str2 string
	fmt.Scanf("%s", &str1)
	fmt.Scanf("%s", &str2)
	A := make([]int, 0)
	B := make([]int, 0)
	for i := len(str1) - 1; i >= 0; i-- {
		A = append(A, int(str1[i]-'0'))
	}
	for i := len(str2) - 1; i >= 0; i-- {
		B = append(B, int(str2[i]-'0'))
	}
	C := add(A, B)
	for i := len(C) - 1; i >= 0; i-- {
		fmt.Print(C[i])
	}
}
func add(A, B []int) []int {
	if len(A) < len(B) {
		return add(B, A)
	}
	C := make([]int, 0)
	up := 0
	for i:=0;i < len(A);i++ {
		up += A[i]
		if i < len(B) {
			up += B[i]
		}
		C = append(C, up%10)
		up = up / 10
	}
	if up != 0 {
		C = append(C, 1)
	}
	return C
}
