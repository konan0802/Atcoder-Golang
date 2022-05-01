package main

import (
	"bufio"
	"fmt"
	"os"
)

func input(sc *bufio.Scanner) string {
	sc.Scan()
	N := sc.Text()
	return N
}

func main() {
	file, _ := os.Open("test.txt")
	sc := bufio.NewScanner(file)
	a := input(sc)
	fmt.Println(a)
}
