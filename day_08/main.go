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

func (tt Trees) Scores(i int) (int, int) {
	l := len(tt) - 1

	if i < 1 || i >= l {
		return 0, 0
	}

	s := tt[i]
	left, right := 1, 1

	for left < i {
		j := i - left

		if tt[j].H >= s.H {
			break
		}

		left++
	}

	for right < l-i {
		j := i + right

		if tt[j].H >= s.H {
			break
		}

		right++
	}

	return left, right
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

func (g Grid) Scores(x, y int) (left int, right int, up int, down int) {
	left, right = g.rows[x].Scores(y)
	up, down = g.Column(y).Scores(x)
	return
}

func (g Grid) Scenic(x, y int) int {
	l, r, u, d := g.Scores(x, y)
	return l * r * u * d
}

func (g Grid) MaxSenic() int {
	var max = -1

	for x, row := range g.rows {
		for y := range row {
			s := g.Scenic(x, y)
			if s > max {
				max = s
			}
		}
	}

	return max
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

	defer f.Close()

	grid := Scan(f)
	fmt.Println("part one value: ", grid.Visible().Count())
	fmt.Println("part two value: ", grid.MaxSenic())
}
