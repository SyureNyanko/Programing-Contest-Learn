package main

import "fmt"

var (
	A = 1002
)

func main() {
	coins := []int{500, 100, 50, 10, 5, 1}
	var sum, c int = A, 0
	for _, v := range coins {
		c = c + sum/v
		sum = sum % (v * c)
	}
	fmt.Println(c)
}
