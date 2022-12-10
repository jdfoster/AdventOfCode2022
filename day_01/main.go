package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var fn = "./day_01/input.txt"

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	var sum int
	var sums []int

	for s.Scan() {
		l := s.Text()

		if l == "" {
			sums = append(sums, sum)
			sum = 0
			continue
		}

		v, err := strconv.ParseInt(l, 10, 0)
		if err != nil {
			panic(err)
		}

		sum += int(v)
	}

	var mx int = 0

	for i, v := range sums {
		if v > sums[mx] {
			mx = i
		}
	}

	fmt.Println("part one value: ", sums[mx])

	sort.Ints(sums)
	last := len(sums) - 1
	sum = 0

	for _, v := range sums[last-2:] {
		sum += v
	}

	fmt.Println("part two value: ", sum)
}
