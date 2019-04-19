package main

import "fmt"

var matrix = [][]string{
	{"#", "S", "#", "#", "#", "#", "#", "#", ".", "#"},
	{".", ".", ".", ".", ".", ".", "#", ".", ".", "#"},
	{".", "#", ".", "#", "#", ".", "#", ".", ".", "#"},
	{".", "#", ".", ".", ".", ".", ".", ".", ".", "."},
	{"#", "#", ".", "#", "#", ".", "#", "#", "#", "#"},
	{".", ".", ".", ".", "#", ".", ".", ".", ".", "#"},
	{".", "#", "#", "#", "#", "#", "#", "#", ".", "#"},
	{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
	{".", "#", "#", "#", "#", ".", "#", "#", "#", "."},
	{".", ".", ".", ".", "#", ".", ".", ".", "G", "#"},
}

func Fill2D(slice [][]int, num int) {
	for i, v1 := range slice {
		for j := range v1 {
			slice[i][j] = num
		}
	}
}

type Point struct {
	x    int
	y    int
	turn int
}

func main() {
	INF := 1000000

	queue := make([]Point, 0)
	f := false
	for i, v1 := range matrix {
		for j, v2 := range v1 {
			if v2 == "S" {
				queue = append(queue, Point{i, j, 0})
				f = true
				break
			}
		}
		if f {
			break
		}
	}

	/* length */
	l := make([][]int, len(matrix))
	for i := range l {
		l[i] = make([]int, len(matrix[0]))
	}
	Fill2D(l, INF)

	/* vector */
	var v = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for len(queue) != 0 {
		for _, p := range v {
			x1 := queue[0].x + p[0]
			y1 := queue[0].y + p[1]

			if 0 < x1 && x1 < len(matrix) && 0 < y1 && y1 < len(matrix[0]) {
				if l[x1][y1] != INF {
					continue
				}
				if matrix[x1][y1] == "." {
					queue = append(queue, Point{x1, y1, queue[0].turn + 1})
					l[x1][y1] = queue[0].turn + 1
				}
				if matrix[x1][y1] == "G" {
					for _, v := range l {
						for _, v1 := range v {
							if v1 != INF {
								fmt.Printf("%02d", v1)
							} else {
								fmt.Printf("**")
							}

						}
						fmt.Println()
					}
				}
			}
		}
		queue = queue[1:]
	}
}
