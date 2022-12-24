package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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

func (g Grid) String() string {
	b := strings.Builder{}

	for _, r := range g {
		for _, c := range r {
			b.WriteRune(toRune(c))
		}
		b.WriteRune('\n')
	}

	return b.String()
}

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

		if d := q - h; d <= 1 {
			result[idx] = n
			idx++
		}
	}

	return result[:idx]
}

type Path struct {
	Position Position
	Count    int
}

func Walk(grid Grid, start, end Position) (Path, bool) {
	queue := []Path{{Position: start, Count: 0}}
	seen := make(map[Position]struct{})

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if end.Equal(head.Position) {
			return head, true
		}

		if _, ok := seen[head.Position]; ok {
			continue
		}

		seen[head.Position] = struct{}{}

		for _, p := range grid.EqualOrAbove(head.Position) {
			queue = append(queue, Path{Position: p, Count: head.Count + 1})
		}

		grid[head.Position.Y][head.Position.X] = -1
	}

	return Path{}, false
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

func toRune(i int) rune {
	return rune(i + int('a'-1))
}

func main() {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	grid, start, end := Scan(f)
	first, ok := Walk(grid, start, end)
	if !ok {
		panic("failed to find any paths")
	}

	// fmt.Println(grid)
	fmt.Println("part 1 value: ", first.Count)
}
