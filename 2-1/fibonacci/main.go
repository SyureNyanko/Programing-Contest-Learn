package main

import "fmt"

// フィボナッチ数列のn項目を求める
func fibonati(n int) int {
	//0項目が0
	if n <= 0 {
		return 0
	}
	//1項目は1
	if n == 1 {
		return 1
	}
	//n項目は2つ前と1つ前の和
	return fibonati(n-2) + fibonati(n-1)
}

func main() {
	var s int
	s = fibonati(10)
	fmt.Println(s)
}
