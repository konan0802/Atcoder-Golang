package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	solve()
}

// 標準入力
var sc = bufio.NewScanner(os.Stdin)

// ファイル入力
//var file, _ = os.Open("test.txt")
//var sc = bufio.NewScanner(file)

func next() string {
	sc.Scan()
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
	A := nextInts(N)
	B := nextInts(N)

	D := make([]bool, N)
	for i := range D {
		D[i] = false
	}
	E := make([]bool, N)
	for i := range E {
		E[i] = false
	}

	D[0] = true
	E[0] = true

	for i := 1; i < N; i++ {
		if D[i-1] == true {
			if int(math.Abs(float64(A[i-1]-A[i]))) <= K {
				D[i] = true
			}
			if int(math.Abs(float64(A[i-1]-B[i]))) <= K {
				E[i] = true
			}
		}
		if E[i-1] == true {
			if int(math.Abs(float64(B[i-1]-A[i]))) <= K {
				D[i] = true
			}
			if int(math.Abs(float64(B[i-1]-B[i]))) <= K {
				E[i] = true
			}
		}
	}

	if D[N-1] == true || E[N-1] == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
