package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//输入一个n行m列的整数矩阵，再输入q个询问，每个询问包含四个整数x1, y1, x2, y2，表示一个子矩阵的左上角坐标和右下角坐标。
//对于每个询问输出子矩阵中所有数的和。

const N = 1001

func main(){
	var n,m,q int
	fmt.Scanf("%d%d%d",&n,&m,&q)
	matrix:=[N][N]int{}
	sc :=bufio.NewScanner(os.Stdin)
	for i:=1;i<=n;i++{
		sc.Scan()
		line :=strings.Split(sc.Text()," ")
		for j:=1;j<=m;j++{
			matrix[i][j],_=strconv.Atoi(line[j-1])
		}
	}
	sum:=[N][N]int{}
	for i:=1;i<=n;i++{
		for j:=1;j<=m;j++{
			sum[i][j]=sum[i-1][j]+sum[i][j-1]-sum[i-1][j-1]+matrix[i][j]
		}
	}
	var x1,y1,x2,y2 int
	for i:=0;i<q;i++{
		sc.Scan()
		line:=strings.Split(sc.Text()," ")
		x1,_=strconv.Atoi(line[0])
		y1,_=strconv.Atoi(line[1])
		x2,_=strconv.Atoi(line[2])
		y2,_=strconv.Atoi(line[3])
		fmt.Println(submatrixSum(sum,x1,y1,x2,y2))
	}
}

func submatrixSum(q [N][N]int,x1,y1,x2,y2 int) int {
	return q[x2][y2]-q[x2][y1-1]-q[x1-1][y2]+q[x1-1][y1-1]
}
