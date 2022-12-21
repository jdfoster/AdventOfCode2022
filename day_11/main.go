package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var fn = "./day_11/input.txt"

const (
	MonkeyPrefix    = "Monkey"
	ItemsPrefix     = "Starting items"
	OperationPrefix = "Operation"
	OpSquareValue   = "new = old * old"
	OpAddPrefix     = "new = old + "
	OpMultiPrefix   = "new = old * "
	TestPrefix      = "Test"
	TruePrefix      = "If true"
	FalsePrefix     = "If false"
)

type Operation func(int) int

type Monkey struct {
	items []int
	op    Operation
	div   int
	pass  int
	fail  int
	count int
}

func buildAdd(v int) Operation {
	return func(i int) int {
		return v + i
	}
}

func buildMulti(v int) Operation {
	return func(i int) int {
		return v * i
	}
}

func buildSquare() Operation {
	return func(i int) int {
		return i * i
	}
}

type Barrel []*Monkey

func (b Barrel) Round() {
	for _, mky := range b {
		for _, item := range mky.items {
			v := mky.op(item) / 3
			if v%mky.div == 0 {
				b[mky.pass].items = append(b[mky.pass].items, v)
			} else {
				b[mky.fail].items = append(b[mky.fail].items, v)
			}

			mky.items = []int{}
			mky.count++
		}
	}
}

func (b Barrel) Business() int {
	l := len(b)
	counts := make([]int, l)

	for i, mky := range b {
		counts[i] = mky.count
	}

	sort.Ints(counts)
	i, j := counts[l-1], counts[l-2]

	return i * j
}

func Scan(r io.Reader) Barrel {
	s := bufio.NewScanner(r)

	var (
		index  int = -1
		barrel Barrel
	)

	for s.Scan() {
		line := strings.TrimSpace(s.Text())

		if line == "" {
			continue
		}

		ss := strings.Split(line, ":")
		if len(ss) != 2 {
			panic("unexpected length found")
		}

		head, tail := ss[0], ss[1]

		switch {
		case strings.HasPrefix(head, MonkeyPrefix):
			barrel = append(barrel, &Monkey{})
			index++
		case strings.HasPrefix(head, ItemsPrefix):
			mky := barrel[index]

			for _, v := range strings.Split(tail, ",") {
				j, err := strconv.ParseInt(strings.TrimSpace(v), 10, 0)
				if err != nil {
					panic(err)
				}
				mky.items = append(mky.items, int(j))
			}
		case strings.HasPrefix(head, OperationPrefix):
			mky := barrel[index]
			n := strings.TrimSpace(tail)

			if n == OpSquareValue {
				mky.op = buildSquare()
				continue
			}
			if strings.HasPrefix(n, OpAddPrefix) {
				v := strings.TrimPrefix(n, OpAddPrefix)
				j, err := strconv.ParseInt(v, 10, 0)
				if err != nil {
					panic(err)
				}
				mky.op = buildAdd(int(j))
				continue
			}
			if strings.HasPrefix(n, OpMultiPrefix) {
				v := strings.TrimPrefix(n, OpMultiPrefix)
				j, err := strconv.ParseInt(v, 10, 0)
				if err != nil {
					panic(err)
				}
				mky.op = buildMulti(int(j))
				continue
			}

			panic("failed to parse operation")
		case strings.HasPrefix(head, TestPrefix):
			mky := barrel[index]

			v := strings.TrimPrefix(tail, " divisible by ")
			j, err := strconv.ParseInt(v, 10, 0)
			if err != nil {
				panic(err)
			}

			mky.div = int(j)
		case strings.HasPrefix(head, TruePrefix):
			mky := barrel[index]

			v := strings.TrimPrefix(tail, " throw to monkey ")
			j, err := strconv.ParseInt(v, 10, 0)
			if err != nil {
				panic(err)
			}

			mky.pass = int(j)
		case strings.HasPrefix(head, FalsePrefix):
			mky := barrel[index]

			v := strings.TrimPrefix(tail, " throw to monkey ")
			j, err := strconv.ParseInt(v, 10, 0)
			if err != nil {
				panic(err)
			}

			mky.fail = int(j)
		default:
			panic("failed to parse line")
		}
	}

	return barrel
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	barrel := Scan(f)

	for i := 0; i < 20; i++ {
		barrel.Round()
	}

	fmt.Println("part 1 value: ", barrel.Business())
}
