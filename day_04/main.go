package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fn = "./day_04/input.txt"

type Points struct {
	Start int
	End   int
}

func (p Points) Includes(v Points) bool {
	if p.Start > v.Start {
		return false
	}

	if p.End < v.End {
		return false
	}

	return true
}

func (p Points) Overlap(v Points) bool {
	if p.Start > v.End {
		return false
	}

	if p.End < v.Start {
		return false
	}

	return true
}

func parseRange(v string) (res Points, ok bool) {
	e := strings.Split(v, "-")
	if len(e) != 2 {
		return
	}

	a, err := strconv.ParseInt(e[0], 10, 0)
	if err != nil {
		return
	}

	b, err := strconv.ParseInt(e[1], 10, 0)
	if err != nil {
		return
	}

	res.Start = int(a)
	res.End = int(b)
	ok = true

	return
}

type Collection struct {
	First  Points
	Second Points
}

func (c Collection) HasInclusion() bool {
	return c.First.Includes(c.Second) || c.Second.Includes(c.First)
}

func (c Collection) HasOverlap() bool {
	return c.First.Overlap(c.Second)
}

func parseLine(v string) (res Collection, ok bool) {
	e := strings.Split(v, ",")
	if len(e) != 2 {
		return
	}

	res.First, ok = parseRange(e[0])
	res.Second, ok = parseRange(e[1])
	return
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	var (
		first  int
		second int
	)

	for s.Scan() {
		l := s.Text()
		r, ok := parseLine(l)
		if !ok {
			panic("failed to parse line")
		}

		if r.HasInclusion() {
			first++
		}

		if r.HasOverlap() {
			second++
		}
	}

	fmt.Println("part one value: ", first)
	fmt.Println("part two value: ", second)
}
