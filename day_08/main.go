package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode/utf8"
)

var fn = "./day_08/input.txt"

type Tree struct {
	X int
	Y int
	H int
}

type TreeSet map[Tree]struct{}

func (ts TreeSet) Count() int {
	return len(ts)
}

func (ts TreeSet) AddTree(tt ...Tree) {
	for _, t := range tt {
		ts[t] = struct{}{}
	}
}

func (ts TreeSet) AddSet(tt TreeSet) {
	for t := range tt {
		ts.AddTree(t)
	}
}

type Trees []Tree

func (tt Trees) Left() TreeSet {
	set := make(TreeSet, len(tt))

	m := -1
	for _, t := range tt {
		if t.H > m {
			set.AddTree(t)
			m = t.H
		}
	}

	return set
}

func (tt Trees) Right() TreeSet {
	set := make(TreeSet, len(tt))

	m := -1
	for j := len(tt) - 1; j > 0; j-- {
		t := tt[j]
		if t.H > m {
			set.AddTree(t)
			m = t.H
		}
	}

	return set
}

func (tt Trees) Visible() TreeSet {
	set := make(TreeSet)

	for k := range tt.Left() {
		set[k] = struct{}{}
	}

	for k := range tt.Right() {
		set[k] = struct{}{}
	}

	return set
}

type Grid struct {
	rows []Trees
}

func (g Grid) Column(i int) Trees {
	res := make(Trees, len(g.rows))

	for j, r := range g.rows {
		res[j] = r[i]
	}

	return res
}

func (g Grid) Visible() TreeSet {
	set := make(TreeSet)

	for i, r := range g.rows {
		set.AddSet(r.Visible())
		set.AddSet(g.Column(i).Visible())
	}

	return set
}

func Scan(r io.Reader) Grid {
	s := bufio.NewScanner(r)
	var (
		count int
		grid  Grid
	)

	for s.Scan() {
		line := s.Text()
		l := utf8.RuneCountInString(line)

		tt := make(Trees, l)
		for i, r := range line {
			h, err := strconv.ParseInt(string(r), 10, 0)
			if err != nil {
				panic(err)
			}
			tt[i] = Tree{X: i, Y: count, H: int(h)}
		}

		grid.rows = append(grid.rows, tt)
		count++
	}

	return grid
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	grid := Scan(f)
	fmt.Println("part one value: ", grid.Visible().Count())
}
