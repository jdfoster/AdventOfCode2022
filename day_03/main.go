package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var fn = "./day_03/input.txt"

type Set map[rune]struct{}

func (s Set) AddRune(v rune) {
	s[v] = struct{}{}
}

func (s Set) AddString(vv string) {
	for _, v := range vv {
		s.AddRune(v)
	}
}

func (ss Set) Intersect(v Set) (res Set) {
	res = make(Set)

	for s := range ss {
		if _, ok := v[s]; ok {
			res.AddRune(s)
		}
	}

	return
}

func (ss Set) Runes() (res []rune) {
	for s := range ss {
		res = append(res, s)
	}

	return
}

func main() {
	lookup := make(map[rune]int, 52)

	{
		s := 1

		for i := 'a'; i <= 'z'; i++ {
			v := unicode.ToUpper(i)

			lookup[i] = s
			lookup[v] = s + 26

			s++
		}
	}

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	var (
		first  int
		second int
		count  int
		state  = make([]Set, 3)
	)

	for s.Scan() {
		l := s.Text()

		{
			var (
				one = make(Set)
				two = make(Set)
			)

			c := len(l)
			if c%2 != 0 {
				panic("count is odd number")
			}

			m := c / 2
			one.AddString(l[:m])
			two.AddString(l[m:])

			r := one.Intersect(two)
			if len(r) != 1 {
				panic("expect to match 1 rune")
			}

			h := r.Runes()[0]
			first += lookup[h]
		}

		{
			i := count % 3

			state[i] = make(Set)
			state[i].AddString(l)

			if i == 2 {
				a, b, c := state[0], state[1], state[2]
				r := a.Intersect(b).Intersect(c)

				if len(r) != 1 {
					panic("len not equl to 1")
				}

				h := r.Runes()[0]
				second += lookup[h]
			}
		}

		count++
	}

	fmt.Println("part one value: ", first)
	fmt.Println("part two value: ", second)
}
