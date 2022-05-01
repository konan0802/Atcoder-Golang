package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input(sc *bufio.Scanner) string {
	sc.Scan()
	N := sc.Text()
	return N
}

func main() {
	file, _ := os.Open("test.txt")
	sc := bufio.NewScanner(file)

	// 不要なため
	sc.Scan()

	tutuCount := []int{}
	tutuValue := []int{}
	for sc.Scan() {
		each := strings.Split(sc.Text(), " ")
		if each[0] == "1" {
			c, _ := strconv.Atoi(each[2])
			x, _ := strconv.Atoi(each[1])
			tutuCount = append(tutuCount, c)
			tutuValue = append(tutuValue, x)
		} else {
			c, _ := strconv.Atoi(each[1])
			totalCount := 0
			totalValue := 0
			for {
				if totalCount+tutuCount[0] >= c {
					totalValue += tutuValue[0] * (c - totalCount)
					tutuCount[0] = tutuCount[0] - (c - totalCount)
					break
				} else {
					totalValue += tutuValue[0] * tutuCount[0]
					totalCount += tutuCount[0]
					tutuCount = tutuCount[1:]
					tutuValue = tutuValue[1:]
				}
			}
			fmt.Println(totalValue)
		}
	}
}
