package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func input(sc *bufio.Scanner) string {
	sc.Scan()
	N := sc.Text()
	return N
}

func s(n int) string {
	if n == 1 {
		return "1"
	}
	return s(n-1) + " " + strconv.Itoa(n) + " " + s(n-1)
}

func main() {
	file, _ := os.Open("test.txt")
	sc := bufio.NewScanner(file)
	N, _ := strconv.Atoi(input(sc))
	fmt.Println(s(N))
}
