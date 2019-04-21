package main

import (
	"fmt"
)

var (
	S = "ACDBCB"
)

func main() {
	var T = ""
	for len(S) > 0 {
		fmt.Println(S)
		fmt.Printf("%d, %d \n", S[0], S[len(S)-1])
		if S[0] < S[len(S)-1] {
			T = T + string(S[0])
			S = S[1:]
		} else {
			T = T + string(S[len(S)-1])
			fmt.Println(S)
			S = S[0 : len(S)-1]
		}
	}
	fmt.Println(T)
}
