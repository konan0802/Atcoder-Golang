package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	solve()
}

//var sc = bufio.NewScanner(os.Stdin)
var file, _ = os.Open("test.txt")
var sc = bufio.NewScanner(file)

func next() string {
	sc.Scan()
	fmt.Println(sc.Text())
	return sc.Text()
}

func nextInt() int {
	a, _ := strconv.Atoi(next())
	return a
}

func nextInts(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = nextInt()
	}
	return ret
}

func solve() {
	N := nextInt()
	K := nextInt()
	X := nextInt()
	A := nextInts(N)

	sort.Sort(sort.Reverse(sort.IntSlice(A)))

	for i := 0; i < N; i++ {
		coupon := A[i] / X
		if coupon == 0 {
			break
		}
		if coupon > K {
			coupon = K
		}
		A[i] -= X * coupon
		K -= coupon
		if K == 0 {
			break
		}
	}

	if K > 0 {
		sort.Sort(sort.Reverse(sort.IntSlice(A)))

		for i := 0; i < N; i++ {
			A[i] = 0
			K--

			if K == 0 {
				break
			}
		}
	}

	sum := 0
	for _, number := range A {
		sum += number
	}
	fmt.Println(sum)
}
