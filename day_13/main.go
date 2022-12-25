package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var fn = "./day_13/input.txt"

type Signal struct {
	child *Signal
}

type Unknown struct {
	Value interface{}
}

func (u Unknown) Int() (int, bool) {
	v, ok := u.Value.(float64)
	if ok {
		return int(v), true
	}

	return 0, false
}

func (u Unknown) List() ([]Unknown, bool) {
	vv, ok := u.Value.([]interface{})
	if ok {
		uu := make([]Unknown, len(vv))
		for i, v := range vv {
			uu[i] = Unknown{v}
		}

		return uu, true
	}

	return []Unknown{}, false
}

func Compare(left, right Unknown) int {
	// l < r = +1; order is correct
	// r = r =  0; match check next input
	// l > r = -1; order is incorrect

	a, ao := left.Int()
	b, bo := right.Int()

	if ao && bo {
		if a == b {
			return 0
		}

		if a < b {
			return 1
		}

		return -1
	}

	aa, aao := left.List()
	bb, bbo := right.List()

	if aao && bbo {
		if len(aa) == 0 && len(bb) > 0 {
			return 1
		}

		for i, a := range aa {
			if i > len(bb)-1 {
				return -1
			}

			c := Compare(a, bb[i])
			if c < 0 {
				return -1
			}

			if c > 0 {
				return 1
			}
		}

		if len(aa) < len(bb) {
			return 1
		}

		return 0
	}

	if ao && bbo {
		return Compare(Unknown{[]interface{}{left.Value}}, right)
	}

	if aao && bo {
		return Compare(left, Unknown{[]interface{}{right.Value}})
	}

	fmt.Println("left  >>> ", left)
	fmt.Println("right >>> ", right)

	panic("failed to compare items")
}

type Pair struct {
	Left  *Unknown
	Right *Unknown
}

func (p Pair) Compare() int {
	return Compare(*p.Left, *p.Right)
}

func NewPair() *Pair {
	return &Pair{
		Left:  &Unknown{Value: nil},
		Right: &Unknown{Value: nil},
	}
}

type Pairs []*Pair

func (pp Pairs) Indices() []int {
	var result []int

	for i, p := range pp {
		switch p.Compare() {
		case -1, 0:
			continue
		case 1:
			// use +1 to offset for 1 based index
			result = append(result, i+1)
		}
	}

	return result
}

func (pp Pairs) Sum() int {
	var result int

	for _, v := range pp.Indices() {
		result += v
	}

	return result
}

func Scan(r io.Reader) Pairs {
	s := bufio.NewScanner(r)

	var (
		result []*Pair
		pair   *Pair = NewPair()
	)

	for s.Scan() {
		if s.Text() == "" {
			result = append(result, pair)
			pair = NewPair()
			continue
		}

		ptr := &pair.Left.Value
		if *ptr != nil {
			ptr = &pair.Right.Value
		}

		if err := json.Unmarshal(s.Bytes(), ptr); err != nil {
			panic(err)
		}
	}

	return result
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	pp := Scan(f)
	first := pp.Sum()
	fmt.Println("part 1 value: ", first)
}
