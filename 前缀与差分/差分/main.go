package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//输入一个长度为 n 的整数序列。
//接下来输入 m 个操作，每个操作包含三个整数 l,r,c，表示将序列中 [l,r] 之间的每个数加上 c。
//请你输出进行完所有操作后的序列
func main(){
	var n,m int
	fmt.Scanf("%d%d",&n,&m)
	q:=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scanf("%d",&q[i])
	}

	//原q数组是处理后的前缀和
	for i:= len(q)-1;i>=1;i--{
		q[i]=q[i]-q[i-1]
	}

	sc :=bufio.NewScanner(os.Stdin)
	for i:=0;i<m;i++{
		var l,r,c int
		sc.Scan()
		line:=strings.Split(sc.Text()," ")
		l,_=strconv.Atoi(line[0])
		r,_=strconv.Atoi(line[1])
		c,_=strconv.Atoi(line[2])
		differential(q,l,r,c)
	}
	sum:=0
	for i:=0;i<n;i++{
		sum+=q[i]
		fmt.Printf("%d ",sum)
	}
}

func differential(q []int,l,r,c int){
	q[l-1]+=c
	if r< len(q) {
		q[r] += -c
	}
}
