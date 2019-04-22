package main

import (
	"fmt"
)

type Obj struct {
	w int
	v int
}

type key struct {
	/* i, weight */
	k1, k2 int
}

var (
	objs    = []Obj{{2, 3}, {1, 2}, {3, 4}, {2, 2}}
	W       = 5
	cache_m = make(map[key]int)
)

//retrun value
func maxvalue(i, w int) int {
	if val, ok := cache_m[key{i, w}]; ok {
		return val
	}
	if len(objs) == i {
		return 0
	}
	var ret int
	if w < objs[i].w {
		//next
		ret = maxvalue(i+1, w)
	} else {
		p := maxvalue(i+1, w)
		q := maxvalue(i+1, w-objs[i].w) + objs[i].v
		if p > q {
			cache_m[key{i, w}] = p
			ret = p
		} else {
			ret = q
		}
	}

	cache_m[key{i, w}] = ret
	return ret
}

func main() {
	fmt.Println(maxvalue(0, W))
}
