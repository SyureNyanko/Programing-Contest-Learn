package main

import (
	"fmt"
	"sort"
)

var (
	Boards []int = []int{8, 5, 8}
)

func main() {
	costs := 0
	for len(Boards) > 1 {
		sort.Ints(Boards)
		nb := Boards[0] + Boards[1]
		Boards = Boards[1:]
		Boards[0] = nb
		costs = costs + nb
	}
	fmt.Println(costs)
}
