package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

func buildAdd(a int) Operation {
	return func(b int) int {
		if a > math.MaxInt-b {
			panic(fmt.Sprintf("add %d, %d ; int overflow", a, b))
		}
		return a + b
	}
}

func buildMulti(a int) Operation {
	return func(b int) int {
		if a == 0 || b == 0 {
			return 0
		}

		result := a * b
		if a == 1 || b == 1 {
			return result
		}

		if a == math.MinInt || b == math.MinInt {
			panic(fmt.Sprintf("multiply %d, %d ; int overflow", a, b))
		}

		if result/b != a {
			panic(fmt.Sprintf("multiply %d, %d ; int overflow", a, b))
		}

		return result
	}
}

func buildSquare() Operation {
	return func(i int) int {
		return buildMulti(i)(i)
	}
}

type Barrel []*Monkey

func (b Barrel) Round(d int) {
	var supermod int = 1

	// Inspection calculation with the operation function overflows for part b
	// using signed integers (int), to avoid this we must reduce the value by
	// dividing by the product of all the denominators.

	for _, mky := range b {
		supermod *= mky.div
	}

	for _, mky := range b {
		for _, item := range mky.items {

			v := mky.op(item%supermod) / d
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

	defer f.Close()

	barrel := Scan(f)

	for i := 0; i < 20; i++ {
		barrel.Round(3)

	}

	fmt.Println("part 1 value: ", barrel.Business())

	f.Seek(0, io.SeekStart)
	barrel = Scan(f)

	for i := 0; i < 10000; i++ {
		barrel.Round(1)
	}

	fmt.Println("part 2 value: ", barrel.Business())
}
