package main

import (
	"fmt"
	"math/rand"
	"time"
)

var n = 100000
var l = make([]int, n)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		l[i] = int(rand.Intn(n))
	}
	start := time.Now().Nanosecond()
	// fmt.Printf("%+v \n", l)
	quickSortR(0, len(l)-1)
	// fmt.Printf("%+v \n", l)
	end := time.Now().Nanosecond()
	fmt.Printf("finished in: %d", end-start)
}

func quickSortR(p int, r int) {
	// fmt.Printf("p: %d, r: %d\n", p, r)
	if p >= r {
		return
	}
	k := partition(p, r)
	// fmt.Printf("k: %d\n", k)
	quickSortR(p, k-1)
	quickSortR(k+1, r)
}

func partition(p int, r int) (i int) {
	i = p
	for j := p; j < r; j++ {
		if l[j] > l[r] {
			// fmt.Printf("switch %d %d!!\n", i, j)
			l[i], l[j] = l[j], l[i]
			// fmt.Printf("%+v\n", l)
			i++
		}
	}

	l[i], l[r] = l[r], l[i]
	return i

}
