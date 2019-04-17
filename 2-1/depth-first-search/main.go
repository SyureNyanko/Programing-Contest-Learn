package main

import (
	"fmt"
)

const (
	k = 99
)

var ints []int
var n int

/*
sum == kになるまで全探索
*/
func dfs(i int, sum int) bool {
	if i == n {
		return sum == k
	}
	/* not choose */
	if dfs(i+1, sum) {
		return true
	}
	/* choose */
	if dfs(i+1, sum+ints[i]) {
		return true
	}
	return false
}

func main() {
	ints = []int{1, 5, 7, 9, 5, 3, 2, 10}
	n = len(ints)
	if dfs(0, 0) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
