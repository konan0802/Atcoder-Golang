package main

import (
	"bufio"
	"fmt"
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
//var sc = bufio.NewScanner(os.Stdin)

// ファイル入力
var file, _ = os.Open("test.txt")
var sc = bufio.NewScanner(file)

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
	fmt.Println(N)
}
