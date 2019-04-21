package main

import (
	"fmt"
)

var (
	R            = 10
	Points []int = []int{1, 7, 15, 20, 30, 50}
	Ans    []int = []int{0, 0, 0, 0, 0, 0}
)

func main() {
	p := 0
	i := 0
	for i < len(Points)-1 {
		for Points[i]-Points[p] <= R {
			i = i + 1
		}
		p = i
		Ans[i-1] = 1
	}

	fmt.Println(Ans)
}
