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
	var C []int
	if cmp(A, B) {
		C = minus(A, B)
	} else {
		fmt.Print("-")
		C = minus(B, A)
	}
	for i := len(C) - 1; i >= 0; i-- {
		fmt.Print(C[i])
	}
}

func cmp(A, B []int) bool {
	if len(A) > len(B) {
		return true
	} else if len(A) < len(B) {
		return false
	} else {
		for i := len(A) - 1; i >= 0; i-- {
			if A[i] == B[i] {
				continue
			} else {
				return A[i] > B[i]
			}
		}
	}
	return true
}

func minus(A, B []int) []int {
	C := make([]int, 0)
	up := 0
	for i := 0; i < len(A); i++ {
		if i < len(B) {
			up += A[i] - B[i] + 10
		} else {
			up += A[i] + 10
		}
		C = append(C, up%10)
		if up >= 10 {
			up = 0
		} else {
			up = -1
		}
	}
	i := len(C) - 1
	for ; i > 0 && C[i] == 0; i-- {
	}
	return C[:i+1]
}
