package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

var fn = "./day_12/input.txt"

type Position struct {
	X int
	Y int
}

func (p Position) Equal(o Position) bool {
	return p.X == o.X && p.Y == o.Y
}

type Grid [][]int

func (g Grid) Height(p Position) (int, bool) {
	if p.X < 0 || p.Y < 0 {
		return 0, false
	}

	if p.Y >= len(g) || p.X >= len(g[p.Y]) {
		return 0, false
	}

	return g[p.Y][p.X], true
}

func (g Grid) EqualOrAbove(p Position) []Position {
	result := make([]Position, 8)
	var idx int

	h, ok := g.Height(p)
	if !ok {
		return []Position{}
	}

	aa := []Position{
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: 0, Y: -1},
	}

	for _, a := range aa {
		n := Position{X: p.X + a.X, Y: p.Y + a.Y}

		q, ok := g.Height(n)
		if !ok {
			continue
		}

		if d := q - h; d >= 0 && d <= 1 {
			result[idx] = n
			idx++
		}
	}

	return result[:idx]
}

type Path struct {
	Position Position
	Visited  []Position
}

func (p Path) Seen(o Position) bool {
	for _, v := range p.Visited {
		if v.Equal(o) {
			return true
		}
	}

	return false
}

func Walk(grid Grid, start, end Position) []Path {
	var (
		result []Path
		mv     func(Path)
	)

	mv = func(a Path) {
		if end.Equal(a.Position) {
			result = append(result, a)
			return
		}

		for _, p := range grid.EqualOrAbove(a.Position) {
			if a.Seen(p) {
				continue
			}

			mv(Path{Position: p, Visited: append(a.Visited, a.Position)})
		}

	}

	for _, p := range grid.EqualOrAbove(start) {
		mv(Path{Position: p, Visited: []Position{start}})
	}

	return result
}

func Shortest(grid Grid, start, end Position) (int, bool) {
	pp := Walk(grid, start, end)
	if len(pp) == 0 {
		return 0, false
	}

	a := make([]int, len(pp))
	for i, p := range pp {
		a[i] = len(p.Visited)
	}

	sort.Ints(a)

	return a[0], true
}

func Scan(r io.Reader) (Grid, Position, Position) {
	s := bufio.NewScanner(r)
	var (
		count int = 0
		start Position
		end   Position
		grid  Grid
	)

	for s.Scan() {
		line := s.Text()
		grid = append(grid, make([]int, len(line)))

		for i, c := range line {
			grid[count][i] = toNum(c)

			if c == 'S' {
				grid[count][i] = toNum('a')
				start = Position{X: i, Y: count}
			}

			if c == 'E' {
				grid[count][i] = toNum('z')
				end = Position{X: i, Y: count}
			}
		}
		count++
	}

	return grid, start, end
}

func toNum(r rune) int {
	return int(r - 'a' + 1)
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	grid, start, end := Scan(f)
	first, ok := Shortest(grid, start, end)
	if !ok {
		panic("failed to find any paths")
	}

	fmt.Println("part 1 value: ", first)
}
