package main

import (
	"fmt"
)

func check(matrix [][]string, i, j int) {
	p := []int{i - 1, i, i + 1, i - 1, i + 1, i - 1, i, i + 1}
	q := []int{j - 1, j - 1, j - 1, j, j, j + 1, j + 1, j + 1}
	for r := range p {
		if p[r] >= 0 && q[r] >= 0 && p[r] < len(matrix) && q[r] < len(matrix[0]) {
			if matrix[p[r]][q[r]] == "W" {
				matrix[p[r]][q[r]] = "."
				check(matrix, p[r], q[r])
			}
		}
	}
}

func solve(matrix [][]string) {
	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			/* slice is ref */
			if matrix[i][j] == "W" {
				check(matrix, i, j)
				sum = sum + 1
				for k := 0; k < len(matrix); k++ {
					for l := 0; l < len(matrix[0]); l++ {
						fmt.Printf(matrix[k][l])
					}
					fmt.Printf("\n")
				}
				fmt.Println("#######################")
			}

		}
	}
	fmt.Println(sum)
}

func main() {
	matrix := [][]string{
		{"W", ".", ".", ".", ".", ".", ".", ".", ".", "W", "W", "."},
		{".", "W", "W", "W", ".", ".", ".", ".", ".", ".", "W", "."},
		{".", "W", "W", "W", ".", ".", ".", ".", ".", ".", "W", "."},
		{".", "W", "W", "W", ".", ".", ".", ".", ".", ".", "W", "."},
		{".", "W", ".", "W", ".", ".", ".", ".", ".", "W", "W", "."},
	}

	solve(matrix)
}
